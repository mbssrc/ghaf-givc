use super::entry::*;
use crate::pb::{self, *};
use anyhow::{bail, Context};
use async_stream::try_stream;
use givc_common::query::Event;
use std::pin::Pin;
use std::sync::Arc;
use std::time::Duration;
use tonic::{Code, Response, Status};
use tracing::{debug, error, info};

pub use pb::admin_service_server::AdminServiceServer;

use crate::admin::registry::*;
use crate::systemd_api::client::SystemDClient;
use crate::types::*;
use crate::utils::naming::*;
use crate::utils::tonic::*;
use givc_client::endpoint::{EndpointConfig, TlsConfig};
use givc_common::query::*;

const VM_STARTUP_TIME: Duration = Duration::new(10, 0);

// FIXME: this is almost copy of sysfsm::Event.
#[derive(Copy, Clone, Debug, PartialEq)]
pub enum State {
    Init,
    InitComplete,
    HostRegistered,
    VmsRegistered,
}

#[derive(Debug)]
pub struct AdminServiceImpl {
    registry: Registry,
    state: State, // FIXME: use sysfsm statemachine
    tls_config: Option<TlsConfig>,
}

#[derive(Debug, Clone)]
pub struct AdminService {
    inner: Arc<AdminServiceImpl>,
}

impl AdminService {
    pub fn new(use_tls: Option<TlsConfig>) -> Self {
        let inner = Arc::new(AdminServiceImpl::new(use_tls));
        let clone = inner.clone();
        tokio::task::spawn(async move {
            clone.monitor().await;
        });
        Self { inner: inner }
    }
}

impl AdminServiceImpl {
    pub fn new(use_tls: Option<TlsConfig>) -> Self {
        Self {
            registry: Registry::new(),
            state: State::Init,
            tls_config: use_tls,
        }
    }

    fn host_endpoint(&self) -> anyhow::Result<EndpointConfig> {
        let host_mgr = self.registry.by_type(UnitType {
            vm: VmType::Host,
            service: ServiceType::Mgr,
        })?;
        let endpoint = host_mgr
            .agent()
            .with_context(|| "Resolving host agent".to_string())?;
        Ok(EndpointConfig {
            transport: endpoint.into(),
            tls: self.tls_config.clone(),
        })
    }

    pub fn agent_endpoint(&self, name: &String) -> anyhow::Result<EndpointConfig> {
        let endpoint = self.registry.by_name(&name)?.agent()?;
        Ok(EndpointConfig {
            transport: endpoint.into(),
            tls: self.tls_config.clone(),
        })
    }

    pub fn app_entries(&self, name: String) -> anyhow::Result<Vec<String>> {
        if name.contains("@") {
            let list = self.registry.find_names(&name)?;
            Ok(list)
        } else {
            Ok(vec![name])
        }
    }

    pub async fn get_remote_status(
        &self,
        entry: &RegistryEntry,
    ) -> anyhow::Result<crate::types::UnitStatus> {
        let transport = match &entry.placement {
            Placement::Managed(parent) => {
                let parent = self.registry.by_name(parent)?;
                parent.agent()? // Fail, if parent also `Managed`
            }
            Placement::Endpoint(endpoint) => endpoint.clone(), // FIXME: avoid clone!
        };
        let tls_name = transport.tls_name.clone();
        let endpoint = EndpointConfig {
            transport: transport.into(),
            tls: self.tls_config.clone().map(|mut tls| {
                tls.tls_name = Some(tls_name);
                tls
            }),
        };

        let client = SystemDClient::new(endpoint);
        client.get_remote_status(entry.name.clone()).await
    }

    pub async fn send_system_command(&self, name: String) -> anyhow::Result<()> {
        let endpoint = self.host_endpoint()?;
        let client = SystemDClient::new(endpoint);
        client.start_remote(name).await?;
        Ok(())
    }

    pub async fn start_vm(&self, name: &str) -> anyhow::Result<()> {
        let endpoint = self.host_endpoint()?;
        let client = SystemDClient::new(endpoint);

        let status = client
            .get_remote_status(name.to_string())
            .await
            .with_context(|| format!("cannot retrieve vm status for {name}"))?;

        if status.load_state != "loaded" {
            bail!("vm {name} not loaded")
        };

        if status.active_state != "active" {
            client
                .start_remote(name.to_string())
                .await
                .with_context(|| format!("spawn remote VM service {name}"))?;

            tokio::time::sleep(VM_STARTUP_TIME).await;

            let new_status = client
                .get_remote_status(name.to_string())
                .await
                .with_context(|| format!("cannot retrieve vm status for {name}"))?;

            if new_status.active_state != "active" {
                bail!("Unable to launch VM {name}")
            }
        }
        Ok(())
    }

    pub async fn handle_error(&self, entry: RegistryEntry) -> anyhow::Result<()> {
        match (entry.r#type.vm, entry.r#type.service) {
            (VmType::AppVM, ServiceType::App) => {
                if entry.status.is_exitted() {
                    self.registry.deregister(&entry.name)?;
                }
                Ok(())
            }
            (VmType::AppVM, ServiceType::Mgr) | (VmType::SysVM, ServiceType::Mgr) => {
                let name = parse_service_name(&entry.name)?;
                self.start_vm(name)
                    .await
                    .with_context(|| format!("handing error, by restart VM {}", &entry.name))?;
                Ok(()) // FIXME: should use `?` from line above, why it didn't work?
            }
            (x, y) => bail!(
                "Don't known how to handle_error for VM type: {:?}:{:?}",
                x,
                y
            ),
        }
    }

    async fn monitor_routine(&self) -> anyhow::Result<()> {
        let watch_list = self.registry.watch_list();
        for entry in watch_list {
            debug!("Monitoring {}...", &entry.name);
            match self.get_remote_status(&entry).await {
                Err(err) => {
                    error!("could not get status of unit {}: {}", &entry.name, err);
                    self.handle_error(entry)
                        .await
                        .with_context(|| "during handle error")?
                }
                Ok(status) => {
                    let inactive = status.active_state != "active";
                    // Difference from "go" algorithm -- save new status before recovering attempt
                    if inactive {
                        error!(
                            "Status of {} is {}, instead of active. Recovering.",
                            &entry.name, status.active_state
                        )
                    };

                    debug!("Status of {} is {:#?} (updated)", &entry.name, status);
                    // We have immutable copy of entry here, but need update _in registry_ copy
                    self.registry.update_state(&entry.name, status)?;

                    if inactive {
                        self.handle_error(entry)
                            .await
                            .with_context(|| "during handle error")?
                    }
                }
            }
        }
        Ok(())
    }

    pub async fn monitor(&self) {
        let mut watch = tokio::time::interval(Duration::from_secs(5));
        watch.tick().await; // First tick fires instantly
        loop {
            watch.tick().await;
            if let Err(err) = self.monitor_routine().await {
                error!("Error during watch: {}", err);
            }
        }
    }

    // Refactoring kludge
    pub fn register(&self, entry: RegistryEntry) {
        self.registry.register(entry)
    }

    pub async fn start_app(&self, req: ApplicationRequest) -> anyhow::Result<()> {
        if self.state != State::VmsRegistered {
            info!("not all required system-vms are registered")
        }
        let name = req.app_name;
        let vm = req.vm_name.as_deref();
        let vm_name = format_vm_name(&name, vm);
        let systemd_agent = format_service_name(&name, vm);

        info!("Starting app {name} on {vm_name}");
        info!("Agent: {systemd_agent}");

        // Entry unused in "go" code
        match self.registry.by_name(&systemd_agent) {
            std::result::Result::Ok(e) => e,
            Err(_) => {
                self.start_vm(&vm_name)
                    .await
                    .context(format!("Starting vm for {}", &name))?;
                self.registry
                    .by_name(&systemd_agent)
                    .context("after starting VM")?
            }
        };
        let endpoint = self.agent_endpoint(&systemd_agent)?;
        let client = SystemDClient::new(endpoint.clone());
        let app_name = self.registry.create_unique_entry_name(&name.to_string());
        client.start_application(app_name.clone()).await?;
        let status = client.get_remote_status(app_name.clone()).await?;
        if status.active_state != "active" {
            bail!("cannot start unit: {app_name}")
        };

        let app_entry = RegistryEntry {
            name: app_name,
            status: status,
            watch: true,
            r#type: UnitType {
                vm: VmType::AppVM,
                service: ServiceType::App,
            },
            placement: Placement::Managed(systemd_agent),
        };
        self.registry.register(app_entry);
        Ok(())
    }
}

fn app_success() -> anyhow::Result<ApplicationResponse> {
    // FIXME: what should be response
    let res = ApplicationResponse {
        cmd_status: String::from("Command successful."),
        app_status: String::from("Command successful."),
    };
    Ok(res)
}

type Stream<T> =
    Pin<Box<dyn tokio_stream::Stream<Item = std::result::Result<T, Status>> + Send + 'static>>;

#[tonic::async_trait]
impl pb::admin_service_server::AdminService for AdminService {
    async fn register_service(
        &self,
        request: tonic::Request<RegistryRequest>,
    ) -> std::result::Result<tonic::Response<pb::RegistryResponse>, tonic::Status> {
        let req = request.into_inner();

        let entry = RegistryEntry::try_from(req)
            .map_err(|e| Status::new(Code::InvalidArgument, format!("{e}")))?;
        self.inner.register(entry);

        let res = RegistryResponse {
            cmd_status: String::from("Registration successful"),
        };
        Ok(Response::new(res))
    }
    async fn start_application(
        &self,
        request: tonic::Request<ApplicationRequest>,
    ) -> std::result::Result<tonic::Response<ApplicationResponse>, tonic::Status> {
        escalate(request, |req| async {
            self.inner.start_app(req).await?;
            app_success()
        })
        .await
    }
    async fn pause_application(
        &self,
        request: tonic::Request<ApplicationRequest>,
    ) -> std::result::Result<tonic::Response<ApplicationResponse>, tonic::Status> {
        escalate(request, |req| async {
            let agent = self.inner.agent_endpoint(&req.app_name)?;
            let client = SystemDClient::new(agent);
            for each in self.inner.app_entries(req.app_name)? {
                _ = client.pause_remote(each).await?
            }
            app_success()
        })
        .await
    }
    async fn resume_application(
        &self,
        request: tonic::Request<ApplicationRequest>,
    ) -> std::result::Result<tonic::Response<ApplicationResponse>, tonic::Status> {
        escalate(request, |req| async {
            let agent = self.inner.agent_endpoint(&req.app_name)?;
            let client = SystemDClient::new(agent);
            for each in self.inner.app_entries(req.app_name)? {
                _ = client.resume_remote(each).await?
            }
            app_success()
        })
        .await
    }
    async fn stop_application(
        &self,
        request: tonic::Request<ApplicationRequest>,
    ) -> std::result::Result<tonic::Response<ApplicationResponse>, tonic::Status> {
        escalate(request, |req| async {
            let agent = self.inner.agent_endpoint(&req.app_name)?;
            let client = SystemDClient::new(agent);
            for each in self.inner.app_entries(req.app_name)? {
                _ = client.stop_remote(each).await?
            }
            app_success()
        })
        .await
    }
    async fn poweroff(
        &self,
        request: tonic::Request<Empty>,
    ) -> std::result::Result<tonic::Response<Empty>, tonic::Status> {
        escalate(request, |_| async {
            self.inner
                .send_system_command(String::from("poweroff.target"))
                .await?;
            Ok(Empty {})
        })
        .await
    }
    async fn reboot(
        &self,
        request: tonic::Request<Empty>,
    ) -> std::result::Result<tonic::Response<Empty>, tonic::Status> {
        escalate(request, |_| async {
            self.inner
                .send_system_command(String::from("reboot.target"))
                .await?;
            Ok(Empty {})
        })
        .await
    }

    async fn query_list(
        &self,
        request: tonic::Request<Empty>,
    ) -> Result<tonic::Response<QueryListResponse>, tonic::Status> {
        escalate(request, |_| async {
            // Kludge
            let list: Vec<QueryResult> = self
                .inner
                .registry
                .contents()
                .into_iter()
                .map(|item| item.into())
                .collect();
            Ok(QueryListResponse {
                list: list.into_iter().map(|item| item.into()).collect(), // Kludge
            })
        })
        .await
    }

    type WatchStream = Stream<WatchItem>;
    async fn watch(
        &self,
        request: tonic::Request<Empty>,
    ) -> Result<tonic::Response<Self::WatchStream>, tonic::Status> {
        escalate(request, |_| async {
            let (initial_list, mut chan) = self.inner.registry.subscribe();

            let stream = try_stream! {
                yield Event::into_initial(initial_list);

                loop {
                    match chan.recv().await {
                        Ok(event) => {
                            yield event.into()
                        },
                        Err(e) => {
                            error!("Failed to receive subscription item from registry: {e}");
                            break;
                        },
                     }
                 }
            };
            Ok(Box::pin(stream) as Self::WatchStream)
        })
        .await
    }
}


<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.admin\.enable



Whether to enable the GIVC admin module, which is responsible for managing the system\.
The admin module is responsible for registration, monitoring, and proxying commands across a virtualized system
of host, system VMs, and application VMs\.



*Type:*
boolean



*Default:*
` false `



## givc\.admin\.addresses

List of addresses for the admin service to listen on\. Requires a list of type ` transportSubmodule `\.



*Type:*
list of (submodule)



*Default:*

```
addresses = [
  {
    name = "localhost";
    addr = "127.0.0.1";
    protocol = "tcp";
    port = "9000";
  }
];
```



*Example:*

```
addresses = [
  {
    name = "admin-vm";
    addr = "192.168.100.3";
    protocol = "tcp";
    port = "9001";
  }
  {
    name = "admin-vm";
    addr = "unix:///run/givc-admin.sock";
    protocol = "unix";
    # port is ignored
  }
];
```



## givc\.admin\.addresses\.\*\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.admin\.addresses\.\*\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.admin\.addresses\.\*\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.admin\.addresses\.\*\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.admin\.debug



Whether to enable givc-admin debug logging\. This increases the verbosity of the logs\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.admin\.name



Network name of the host running the admin service\.

> **Caution**
> This is used to validate the TLS host name and must match the names used in the transport configurations (addresses)\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.admin\.services



List of microvm services of the system-vms for the admin module to administrate, excluding any dynamic VMs such as app-vm\. Expects a space separated list\.
Must be a of type ‘service’, e\.g\., ‘microvm@net-vm\.service’\.



*Type:*
list of string



*Default:*
` [ ] `



*Example:*
` services = ["microvm@net-vm.service"]; `



## givc\.admin\.tls



TLS options for gRPC connections\. It is enabled by default to discourage unprotected connections,
and requires paths to certificates and key being set\. To disable it use ` tls.enable = false; `\.

> **Caution**
> It is recommended to use a global TLS flag to avoid inconsistent configurations that will result in connection errors\.



*Type:*
submodule



*Default:*

```
tls = {
  enable = true;
  caCertPath = "/etc/givc/ca-cert.pem";
  certPath = /etc/givc/cert.pem";
  keyPath = "/etc/givc/key.pem";
};
```



*Example:*

```
tls = {
  enable = true;
  caCertPath = "/etc/ssl/certs/ca-certificates.crt";
  certPath = "/etc/ssl/certs/server.crt";
  keyPath = "/etc/ssl/private/server.key";
};
```



## givc\.admin\.tls\.enable



Enable the TLS module\. Defaults to ‘true’ and should only be disabled for debugging\.



*Type:*
boolean



*Default:*
` true `



## givc\.admin\.tls\.caCertPath



Path to the CA certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/ca-cert.pem" `



## givc\.admin\.tls\.certPath



Path to the service certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/cert.pem" `



## givc\.admin\.tls\.keyPath



Path to the service key file\.



*Type:*
string



*Default:*
` "/etc/givc/key.pem" `



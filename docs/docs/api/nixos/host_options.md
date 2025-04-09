
<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.host\.enable



Whether to enable givc host agent module, which is responsible for managing system VMs and app VMs…



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.host\.admin

Admin server transport configuration\. This configuration tells the agent how to reach the admin server\.



*Type:*
submodule



*Default:*

```
{
  name = "localhost";
  addr = "127.0.0.1";
  protocol = "tcp";
  port = "9000";
};
```



*Example:*

```
transport =
  {
    name = "admin-vm";
    addr = "192.168.100.3";
    protocol = "tcp";
    port = "9001";
  };
```



## givc\.host\.admin\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.host\.admin\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.host\.admin\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.host\.admin\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.host\.appVms



List of app VM services for the host to administrate\. Expects a space separated list\.
Should be a unit file of type ‘service’ or ‘target’\.



*Type:*
list of string



*Default:*
` [ ] `



*Example:*

```
services = [
  "microvm@app1-vm.service"
  "microvm@app2-vm.service"
];
```



## givc\.host\.debug



Whether to enable enable appvm GIVC agent debug logging\. This increases the verbosity of the logs\.

> **Caution**
> Enabling debug logging may expose sensitive information in the logs, especially if the appvm uses the DBUS submodule\.
> \.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.host\.services



List of systemd units for the manager to administrate\. Expects a space separated list\.
Should be a unit file of type ‘service’ or ‘target’\.



*Type:*
list of string



*Default:*

```
[
  "reboot.target"
  "poweroff.target"
  "sleep.target"
  "suspend.target"
]
```



*Example:*

```
services = [
  "poweroff.target"
  "reboot.target"
];
```



## givc\.host\.systemVms



List of system VM services for the host to administrate, which is joined with the generic “services” option\.
Expects a space separated list\. Should be a unit file of type ‘service’\.



*Type:*
list of string



*Default:*
` [ ] `



*Example:*

```
services = [
  "microvm@net-vm.service"
  "microvm@gui-vm.service"
];
```



## givc\.host\.tls



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



## givc\.host\.tls\.enable



Enable the TLS module\. Defaults to ‘true’ and should only be disabled for debugging\.



*Type:*
boolean



*Default:*
` true `



## givc\.host\.tls\.caCertPath



Path to the CA certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/ca-cert.pem" `



## givc\.host\.tls\.certPath



Path to the service certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/cert.pem" `



## givc\.host\.tls\.keyPath



Path to the service key file\.



*Type:*
string



*Default:*
` "/etc/givc/key.pem" `



## givc\.host\.transport



Transport configuration of the GIVC agent of type ` transportSubmodule `\.

> **Caution**
> This parameter is used to generate and validate the TLS host name\.



*Type:*
submodule



*Default:*
` { } `



*Example:*

```
transport =
  {
    name = "host";
    addr = "192.168.100.2";
    protocol = "tcp";
    port = "9000";
  };
```



## givc\.host\.transport\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.host\.transport\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.host\.transport\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.host\.transport\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



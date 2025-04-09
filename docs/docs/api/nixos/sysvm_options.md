
<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.sysvm\.enable



Whether to enable givc sysvm agent module, which is responsible for managing a system VM and respective services\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.sysvm\.enableUserTlsAccess



Whether to enable user access to TLS keys for the client to run\. This will copy the keys to ` /run/givc ` and makes it accessible to the group
` users ` (default for regular users in NixOS)\.
\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.sysvm\.admin

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
admin = {
  {
    name = "admin-vm";
    addr = "192.168.100.3";
    protocol = "tcp";
    port = "9001";
  };
```



## givc\.sysvm\.admin\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.sysvm\.admin\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.sysvm\.admin\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.sysvm\.admin\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.sysvm\.debug



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



## givc\.sysvm\.hwidIface



Hardware identifier to be used with ` hwidService `\.



*Type:*
string



*Default:*
` "" `



## givc\.sysvm\.hwidService



Hardware identifier service that fetches the MAC address from a network interface\.

> **Note**
> This module is can be used to generate a (somewhat) reproducible hardware id\. It is
> currently unused in the Ghaf project for privacy reasons\.



*Type:*
boolean



*Default:*
` false `



## givc\.sysvm\.services



List of systemd services for the manager to administrate\. Expects a space separated list\.
Should be a unit file of type ‘service’ or ‘target’\.



*Type:*
list of string



*Default:*

```
[
  "reboot.target"
  "poweroff.target"
]
```



## givc\.sysvm\.socketProxy



Optional socket proxy module\. The socket proxy provides a VM-to-VM streaming mechanism with socket enpoints, and can be used
to remote DBUS functionality across VMs\. Hereby, the side running the dbusproxy (e\.g\., a network VM running NetworkManager) is
considered the ‘server’, and the receiving end (e\.g\., the GUI VM) is considered the ‘client’\.

The socket proxy module must be configured on both ends with explicit transport information, and must run on a dedicated TCP port\.
The detailed socket proxy options are described in the respective ` .socketProxy.* ` options\.

> **Note**
> The socket proxy module is a possible transport mechanism for the DBUS proxy module, and must be appropriately configured on both
> ends if used\. In this use case, the ` server ` option is configured automatically and does not need to be set\.



*Type:*
null or (list of (submodule))



*Default:*
` null `



*Example:*

```
givc.appvm.socketProxy = [
  {
    # Configure the remote endpoint
    transport = {
      name = "gui-vm";
      addr = "192.168.100.5;
      port = "9013";
      protocol = "tcp";
    };
    # Socket path
    socket = "/tmp/.dbusproxy_app.sock";
  }
];

```



## givc\.sysvm\.socketProxy\.\*\.server



Whether the module runs as server or client\.

The client/server logic follows the socket providing the service\. The server connects to a local socket
(e\.g\., local system dbus or xdg-dbus-module) and upon successful connection allows connection of a remote socket
client(s)\. The socket proxy client provides a local socket to any service to connect to (e\.g\., dbus client application)\.

> **Note**
> This setting defaults to ` config.givc.dbusproxy.enable ` and can be ignored if dbusproxy is used\.



*Type:*
boolean



*Default:*

```
if hasAttrByPath [ "givc" "dbusproxy" ] config
then
  config.givc.dbusproxy.enable
else false;

```



## givc\.sysvm\.socketProxy\.\*\.socket



Path to the system socket\. Defaults to ` /tmp/.dbusproxy.sock `\.



*Type:*
string



*Default:*
` "/tmp/.dbusproxy.sock" `



## givc\.sysvm\.socketProxy\.\*\.transport



Transport configuration of the socket proxy module of type ` transportSubmodule `\.



*Type:*
submodule



*Default:*
` { } `



*Example:*

```
transport =
  {
    name = "app-vm";
    addr = "192.168.100.123";
    protocol = "tcp";
    port = "9012";
  };
```



## givc\.sysvm\.socketProxy\.\*\.transport\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.sysvm\.socketProxy\.\*\.transport\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.sysvm\.socketProxy\.\*\.transport\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.sysvm\.socketProxy\.\*\.transport\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.sysvm\.tls



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



## givc\.sysvm\.tls\.enable



Enable the TLS module\. Defaults to ‘true’ and should only be disabled for debugging\.



*Type:*
boolean



*Default:*
` true `



## givc\.sysvm\.tls\.caCertPath



Path to the CA certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/ca-cert.pem" `



## givc\.sysvm\.tls\.certPath



Path to the service certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/cert.pem" `



## givc\.sysvm\.tls\.keyPath



Path to the service key file\.



*Type:*
string



*Default:*
` "/etc/givc/key.pem" `



## givc\.sysvm\.transport



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
    name = "net-vm";
    addr = "192.168.100.4";
    protocol = "tcp";
    port = "9000";
  };
```



## givc\.sysvm\.transport\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.sysvm\.transport\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.sysvm\.transport\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.sysvm\.transport\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.sysvm\.wifiManager



Wifi manager to handle wifi related queries with a defined interface\. Deprecated in favor of DBUS proxy\.



*Type:*
boolean



*Default:*
` false `



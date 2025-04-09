
<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.appvm\.enable



Whether to enable GIVC appvm agent module\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.appvm\.admin

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
admin =
  {
    name = "admin-vm";
    addr = "192.168.100.3";
    protocol = "tcp";
    port = "9001";
  };
```



## givc\.appvm\.admin\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.appvm\.admin\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.appvm\.admin\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.appvm\.admin\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.appvm\.applications



List of applications to be supported by the ` appvm ` module\. Interface and options are detailed under ` givc.appvm.applications.*.<option> `\.



*Type:*
list of (submodule)



*Default:*

```
[
  { }
]
```



*Example:*

```
applications = [
  {
    name = "app";
    command = "/run/current-system/sw/bin/app";
    args = [
      "url"
      "file"
    ];
    directories = [ "/tmp" ];
  }
];
```



## givc\.appvm\.applications\.\*\.args



List of allowed argument types for the application\. Currently implemented argument types:

 - ‘url’: URL provided to the application as string
 - ‘flag’: Flag (boolean) provided to the application as string
 - ‘file’: File path provided to the application as string
   If the file argument is used, a list of allowed directories must be provided\.



*Type:*
list of string



*Default:*
` [ ] `



## givc\.appvm\.applications\.\*\.command



Command to run the application\.



*Type:*
string



*Default:*
` "/run/current-system/sw/bin/app" `



## givc\.appvm\.applications\.\*\.directories



List of directories (absolute path) to be whitelisted and used with file arguments\.



*Type:*
list of string



*Default:*
` [ ] `



## givc\.appvm\.applications\.\*\.name



Name of the application\.



*Type:*
string



*Default:*
` "app" `



## givc\.appvm\.debug



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



## givc\.appvm\.socketProxy



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



## givc\.appvm\.socketProxy\.\*\.server



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



## givc\.appvm\.socketProxy\.\*\.socket



Path to the system socket\. Defaults to ` /tmp/.dbusproxy.sock `\.



*Type:*
string



*Default:*
` "/tmp/.dbusproxy.sock" `



## givc\.appvm\.socketProxy\.\*\.transport



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



## givc\.appvm\.socketProxy\.\*\.transport\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.appvm\.socketProxy\.\*\.transport\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.appvm\.socketProxy\.\*\.transport\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.appvm\.socketProxy\.\*\.transport\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.appvm\.tls



TLS options for gRPC connections\. It is enabled by default to discourage unprotected connections,
and requires paths to certificates and key being set\. To disable it use ` tls.enable = false; `\. The
TLS modules default paths’ are overwritten for the ` appvm ` module to allow access for the appvm user (see UID)\.

> **Caution**
> It is recommended to use a global TLS flag to avoid inconsistent configurations that will result in connection errors\.



*Type:*
submodule



*Default:*

```
tls = {
  enable = true;
  caCertPath = "/run/givc/ca-cert.pem";
  certPath = "/run/givc/cert.pem";
  keyPath = "/run/givc/key.pem";
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



## givc\.appvm\.tls\.enable



Enable the TLS module\. Defaults to ‘true’ and should only be disabled for debugging\.



*Type:*
boolean



*Default:*
` true `



## givc\.appvm\.tls\.caCertPath



Path to the CA certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/ca-cert.pem" `



## givc\.appvm\.tls\.certPath



Path to the service certificate file\.



*Type:*
string



*Default:*
` "/etc/givc/cert.pem" `



## givc\.appvm\.tls\.keyPath



Path to the service key file\.



*Type:*
string



*Default:*
` "/etc/givc/key.pem" `



## givc\.appvm\.transport



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
    name = "app-vm";
    addr = "192.168.100.123";
    protocol = "tcp";
    port = "9000";
  };
```



## givc\.appvm\.transport\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.appvm\.transport\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.appvm\.transport\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.appvm\.transport\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.appvm\.uid



UID of the user session to run the ` appvm ` module in\. This prevents to run agent instances for other users (e\.g\., admin) on login\.

> **Note**
> If the application VM is expected to run upon start, the user corresponding to the given UID is expected to
> [linger](https://search\.nixos\.org/options?channel=unstable\&show=users\.users\.%3Cname%3E\.linger\&from=0\&size=50\&sort=relevance\&type=packages\&query=linger)
> to keep the user session alive in the application VM without specific login\.



*Type:*
signed integer



*Default:*
` 1000 `



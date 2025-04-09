
<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.tls\.enable



Whether to enable the givc TLS module\. This module generates GIVC keys and certificates using a transient CA that is removed after generation\.

The hosts lock file ‘/etc/givc/tls\.lock’ is used to lock generation at boot\. If removed, a new CA
is generated and all keys and certificates are re-generated\.

The lock file is automatically removed if a new VM is detected, thus trigerring re-generation\. Removal of a VM will
currently not result in removal of the VMs key and certificates\.

> **Caution**
> This module is not intended to be used in production\. Use in development and testing environments\.



*Type:*
boolean



*Default:*
` false `



## givc\.tls\.agents

List of agents to generate TLS certificates for\. Requires a list of ‘transportSubmodule’\.

> **Note**
> This module generates an ext4 image file for each agent (except the host)\. The image file is created in the storage path
> and named after the agent name\. The image can be mounted read-only into a VM using virtiofs\.



*Type:*
list of (submodule)



*Default:*
` [ ] `



*Example:*

```
agents = [
  {
    {
      name = "app1-vm";
      addr = "192.168.100.123";
    }
    {
      name = "app2-vm";
      addr = "192.168.100.124";
    }
  }
];
```



## givc\.tls\.agents\.\*\.addr



Address identifier\. Can be one of IPv4 address, vsock address, or unix socket path\.



*Type:*
string



*Default:*
` "127.0.0.1" `



## givc\.tls\.agents\.\*\.name



Identifier for network, host, and/or TLS name\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.tls\.agents\.\*\.port



Port identifier for TCP or vsock addresses\. Ignored for unix socket addresses\.



*Type:*
string



*Default:*
` "9000" `



## givc\.tls\.agents\.\*\.protocol



Protocol identifier\. Can be one of ‘tcp’, ‘unix’, or ‘vsock’\.



*Type:*
one of “tcp”, “unix”, “vsock”



*Default:*
` "tcp" `



## givc\.tls\.generatorHostName



Host name of the certificate generator\. This is necessary to prevent generating an image file for the host\.



*Type:*
string



*Default:*
` "localhost" `



## givc\.tls\.storagePath



Storage path for generated keys and certificates\. Will use subdirectories for each agent by name\.



*Type:*
string



*Default:*
` "/etc/givc" `



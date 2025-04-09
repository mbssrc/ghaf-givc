
<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

## givc\.dbusproxy\.enable

Whether to enable the givc-dbusproxy module\. This module is a wrapper of the ` xdg-dbus-proxy `, and can be used to filter
specific dbus namespaces in the system or session bus, and expose them on a local socket\. It can be used with the socket
proxy to allow remote access to the configured system functionality\.

> **Caution**
> The ` dbusproxy ` module exposes the system or session bus to a remote endpoint when used in combination with the socket
> proxy module\. Make sure to configure and audit the policies to prevent unwanted access\. If policies are incorrect or
> too broadly defined, this can result in severe security issues such as remote code execution\.

When the dbusproxy module is enabled, it exposes the respecive bus as a Unix Domain Socket file\. If the socket proxy is
enabled as well, it is automatically configured to run as server\.

> **Note**
> \*
> 
>  - If enabled, either the system or session bus option must be set
>  - At least one policy value (see/talk/own/call/broadcast) must be set
>  - To run the session bus proxy, a non-system user with a configured UID is required

Filtering is enabled by default, and the config requires at least one policy value (see/talk/own) to be set\. For more
details, please refer to the [xdg-dbus-proxy manual](https://www\.systutorials\.com/docs/linux/man/1-xdg-dbus-proxy/)\.

Policy values are a list of strings, where each string is translated into the respective argument\. Multiple instances
of the same value are allowed\. In order to create your policies, consider using ` busctl ` to list the available services
and their properties\.



*Type:*
boolean



*Default:*
` false `



## givc\.dbusproxy\.session



Configuration of givc-dbusproxy for user session bus\.



*Type:*
submodule



*Default:*

```
session = {
  user = "root";
  socket = "/tmp/.dbusproxy.sock";
  policy = { };
  debug = false;
};
```



*Example:*

```
givc.dbusproxy = {
  enable = true;
  session = {
    enable = true;
    user = "ghaf";
    socket = "/tmp/.dbusproxy_app.sock";
    policy.talk = [
      "org.mpris.MediaPlayer2.playerctld.*"
    ];
  };
};
```



## givc\.dbusproxy\.session\.enable



Whether to enable givc dbus component\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.dbusproxy\.session\.debug



Whether to enable monitoring of the underlying xdg-dbus-proxy\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.dbusproxy\.session\.policy



Policy submodule for the dbus proxy\.

Filtering is applied only to outgoing signals and method calls and incoming broadcast signals\. All replies (errors or method returns) are allowed once for an outstanding method call, and never otherwise\.
If a client ever receives a message from another peer on the bus, the senders unique name is made visible, so the client can track caller lifetimes via NameOwnerChanged signals\. If a client calls a method on
or receives a broadcast signal from a name (even if filtered to some subset of paths or interfaces), that names basic policy is considered to be (at least) TALK, from then on\.



*Type:*
submodule



*Default:*
` { } `



## givc\.dbusproxy\.session\.policy\.broadcast



> BROADCAST policy:
> You can receive broadcast signals from the name/ID

From xdg-dbus-proxy manual:

The RULE in these options determines what interfaces, methods and object paths are allowed\. It must be of the form \[METHOD]\[@PATH], where METHOD
can be either ‘*’ or a D-Bus interface, possible with a '\.*’ suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a ‘/\*’ suffix\.



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.session\.policy\.call



CALL policy:

 - You can call the specific methods

From xdg-dbus-proxy manual:

The RULE in these options determines what interfaces, methods and object paths are allowed\. It must be of the form \[METHOD]\[@PATH], where METHOD
can be either ‘*’ or a D-Bus interface, possible with a '\.*’ suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a ‘/\*’ suffix\.



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.session\.policy\.own



OWN policy:

 - You are allowed to call RequestName/ReleaseName/ListQueuedOwners on the name



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.session\.policy\.see



SEE policy:

 - The name/ID is visible in the ListNames reply
 - The name/ID is visible in the ListActivatableNames reply
 - You can call GetNameOwner on the name
 - You can call NameHasOwner on the name
 - You see NameOwnerChanged signals on the name
 - You see NameOwnerChanged signals on the ID when the client disconnects
 - You can call the GetXXX methods on the name/ID to get e\.g\. the peer pid
 - You get AccessDenied rather than NameHasNoOwner when sending messages to the name/ID



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.session\.policy\.talk



TALK policy:

 - You can send any method calls and signals to the name/ID
 - You will receive broadcast signals from the name/ID (if you have a match rule for them)
 - You can call StartServiceByName on the name



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.session\.socket



Socket path used to connect to the bus\. Defaults to ` /tmp/.dbusproxy.sock `\.



*Type:*
string



*Default:*
` "/tmp/.dbusproxy.sock" `



## givc\.dbusproxy\.session\.user



User to run the xdg-dbus-proxy service as\. This option must be set to allow a remote user to connect to the bus\.
Defaults to ` root `\.



*Type:*
string



*Default:*
` "root" `



## givc\.dbusproxy\.system



Configuration of givc-dbusproxy for system bus\.



*Type:*
submodule



*Default:*

```
system = {
  user = "root";
  socket = "/tmp/.dbusproxy.sock";
  policy = { };
  debug = false;
};
```



*Example:*

```
givc.dbusproxy = {
  enable = true;
  system = {
    enable = true;
    user = "ghaf";
    socket = "/tmp/.dbusproxy_net.sock";
    policy = {
      talk = [
        "org.freedesktop.NetworkManager.*"
        "org.freedesktop.Avahi.*"
      ];
      call = [
        "org.freedesktop.UPower=org.freedesktop.UPower.EnumerateDevices"
      ];
    };
  };
```



## givc\.dbusproxy\.system\.enable



Whether to enable givc dbus component\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.dbusproxy\.system\.debug



Whether to enable monitoring of the underlying xdg-dbus-proxy\.



*Type:*
boolean



*Default:*
` false `



*Example:*
` true `



## givc\.dbusproxy\.system\.policy



Policy submodule for the dbus proxy\.

Filtering is applied only to outgoing signals and method calls and incoming broadcast signals\. All replies (errors or method returns) are allowed once for an outstanding method call, and never otherwise\.
If a client ever receives a message from another peer on the bus, the senders unique name is made visible, so the client can track caller lifetimes via NameOwnerChanged signals\. If a client calls a method on
or receives a broadcast signal from a name (even if filtered to some subset of paths or interfaces), that names basic policy is considered to be (at least) TALK, from then on\.



*Type:*
submodule



*Default:*
` { } `



## givc\.dbusproxy\.system\.policy\.broadcast



> BROADCAST policy:
> You can receive broadcast signals from the name/ID

From xdg-dbus-proxy manual:

The RULE in these options determines what interfaces, methods and object paths are allowed\. It must be of the form \[METHOD]\[@PATH], where METHOD
can be either ‘*’ or a D-Bus interface, possible with a '\.*’ suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a ‘/\*’ suffix\.



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.system\.policy\.call



CALL policy:

 - You can call the specific methods

From xdg-dbus-proxy manual:

The RULE in these options determines what interfaces, methods and object paths are allowed\. It must be of the form \[METHOD]\[@PATH], where METHOD
can be either ‘*’ or a D-Bus interface, possible with a '\.*’ suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a ‘/\*’ suffix\.



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.system\.policy\.own



OWN policy:

 - You are allowed to call RequestName/ReleaseName/ListQueuedOwners on the name



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.system\.policy\.see



SEE policy:

 - The name/ID is visible in the ListNames reply
 - The name/ID is visible in the ListActivatableNames reply
 - You can call GetNameOwner on the name
 - You can call NameHasOwner on the name
 - You see NameOwnerChanged signals on the name
 - You see NameOwnerChanged signals on the ID when the client disconnects
 - You can call the GetXXX methods on the name/ID to get e\.g\. the peer pid
 - You get AccessDenied rather than NameHasNoOwner when sending messages to the name/ID



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.system\.policy\.talk



TALK policy:

 - You can send any method calls and signals to the name/ID
 - You will receive broadcast signals from the name/ID (if you have a match rule for them)
 - You can call StartServiceByName on the name



*Type:*
null or (list of string)



*Default:*
` null `



## givc\.dbusproxy\.system\.socket



Socket path used to connect to the bus\. Defaults to ` /tmp/.dbusproxy.sock `\.



*Type:*
string



*Default:*
` "/tmp/.dbusproxy.sock" `



## givc\.dbusproxy\.system\.user



User to run the xdg-dbus-proxy service as\. This option must be set to allow a remote user to connect to the bus\.
Defaults to ` root `\.



*Type:*
string



*Default:*
` "root" `




<!--
  Copyright 2025 TII (SSRC) and the Ghaf contributors
  SPDX-License-Identifier: CC-BY-SA-4.0
-->

# GRPC API Documentation

## Table of Contents


- [Admin Service API](#admin-service-api)


- [HWID Service API](#hwid-service-api)


- [Locale Service API](#locale-service-api)


- [Socket Proxy Service API](#socket-proxy-service-api)


- [Stats Service API](#stats-service-api)


- [Systemd Service API](#systemd-service-api)


- [WiFi Service API](#wifi-service-api)

- [Scalar Value Types](#scalar-value-types)




## <a id="admin-service-api">Admin Service API (admin/admin.proto)</a>

### Services

<a name="admin.AdminService"/>

#### AdminService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterService | [RegistryRequest](#admin.RegistryRequest) | [RegistryResponse](#admin.RegistryRequest) | Register a remote agent or service |
| StartApplication | [ApplicationRequest](#admin.ApplicationRequest) | [StartResponse](#admin.ApplicationRequest) | Start a remote application |
| StartVM | [StartVMRequest](#admin.StartVMRequest) | [StartResponse](#admin.StartVMRequest) | Start a VM |
| StartService | [StartServiceRequest](#admin.StartServiceRequest) | [StartResponse](#admin.StartServiceRequest) | Start a remote service |
| PauseApplication | [ApplicationRequest](#admin.ApplicationRequest) | [ApplicationResponse](#admin.ApplicationRequest) | Pause (freeze) a remote application |
| ResumeApplication | [ApplicationRequest](#admin.ApplicationRequest) | [ApplicationResponse](#admin.ApplicationRequest) | Resume (un-freeze) a remote application |
| StopApplication | [ApplicationRequest](#admin.ApplicationRequest) | [ApplicationResponse](#admin.ApplicationRequest) | Stop a remote application |
| SetLocale | [LocaleRequest](#admin.LocaleRequest) | [Empty](#admin.LocaleRequest) | Set locale (broadcasted across system) |
| SetTimezone | [TimezoneRequest](#admin.TimezoneRequest) | [Empty](#admin.TimezoneRequest) | Set timezone (broadcasted across system) |
| Poweroff | [Empty](#admin.Empty) | [Empty](#admin.Empty) | System poweroff command |
| Reboot | [Empty](#admin.Empty) | [Empty](#admin.Empty) | System reboot command |
| Suspend | [Empty](#admin.Empty) | [Empty](#admin.Empty) | System suspend command |
| Wakeup | [Empty](#admin.Empty) | [Empty](#admin.Empty) | System wakeup command |
| GetUnitStatus | [UnitStatusRequest](#admin.UnitStatusRequest) | [.systemd.UnitStatus](#admin.UnitStatusRequest) | Get systemd unit status |
| GetStats | [StatsRequest](#admin.StatsRequest) | [.stats.StatsResponse](#admin.StatsRequest) | Get stats information |
| QueryList | [Empty](#admin.Empty) | [QueryListResponse](#admin.Empty) | Get list of monitored units |
| Watch | [Empty](#admin.Empty) | [WatchItem](#admin.Empty) | Get stream of monitored units |

 <!-- end services -->

### Messages

<a name="admin.ApplicationRequest"/>

#### ApplicationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AppName | [string](#string) |  | Application name |
| VmName | [string](#string) | optional | Name of the VM hosting the application |
| Args | [string](#string) | repeated | Application arguments |






<a name="admin.ApplicationResponse"/>

#### ApplicationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CmdStatus | [string](#string) |  | Status of the RPC command |
| AppStatus | [string](#string) |  | Status of the application |






<a name="admin.Empty"/>

#### Empty
Empty message






<a name="admin.LocaleRequest"/>

#### LocaleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Locale | [string](#string) |  | String with locale value (TODO: format) |






<a name="admin.QueryListItem"/>

#### QueryListItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  | Name of the unit to query |
| Description | [string](#string) |  | Description of the unit |
| VmStatus | [string](#string) |  | Status of the VM hosting the unit |
| TrustLevel | [string](#string) |  | Trust Level (future use) |
| VmType | [string](#string) |  | Type of the VM (future use) |
| ServiceType | [string](#string) |  | Type of the service (future use) |
| VmName | [string](#string) | optional | Name of the VM to query; None for host running services |
| AgentName | [string](#string) | optional | NAme of the agent to query; None for agents |






<a name="admin.QueryListResponse"/>

#### QueryListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| List | [QueryListItem](#admin.QueryListItem) | repeated | List of query responses |






<a name="admin.RegistryRequest"/>

#### RegistryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  | Component name for registry entry |
| Parent | [string](#string) |  | Parent component identifier (registry name) |
| Type | [uint32](#uint32) |  | Component type |
| Transport | [TransportConfig](#admin.TransportConfig) |  | TransportConfig |
| State | [systemd.UnitStatus](#systemd.UnitStatus) |  | Unit status of the component (systemd) |






<a name="admin.RegistryResponse"/>

#### RegistryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Error | [string](#string) | optional | Error message |






<a name="admin.StartResponse"/>

#### StartResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| registryId | [string](#string) |  | Registry ID for newly started app, service or vm |






<a name="admin.StartServiceRequest"/>

#### StartServiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ServiceName | [string](#string) |  | Name of the service to start |
| VmName | [string](#string) |  | Name of the VM hosting service |






<a name="admin.StartVMRequest"/>

#### StartVMRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| VmName | [string](#string) |  | Name of the VM to start |






<a name="admin.StatsRequest"/>

#### StatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| VmName | [string](#string) |  | VM name to query status information |






<a name="admin.TimezoneRequest"/>

#### TimezoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Timezone | [string](#string) |  | String with timezone value (TODO: format) |






<a name="admin.TransportConfig"/>

#### TransportConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Protocol | [string](#string) |  | Protocol identifier, one of tcp, vsock, unix |
| Address | [string](#string) |  | IPv4 address, vsock CID, or unix socket path |
| Port | [string](#string) |  | Port number |
| Name | [string](#string) |  | Host name |






<a name="admin.UnitStatusRequest"/>

#### UnitStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| VmName | [string](#string) |  | Name of the VM hosting the unit |
| UnitName | [string](#string) |  | Name of the unit |






<a name="admin.WatchItem"/>

#### WatchItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Initial | [QueryListResponse](#admin.QueryListResponse) |  |  |
| Added | [QueryListItem](#admin.QueryListItem) |  |  |
| Updated | [QueryListItem](#admin.QueryListItem) |  |  |
| Removed | [QueryListItem](#admin.QueryListItem) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="hwid-service-api">HWID Service API (hwid/hwid.proto)</a>

### Services

<a name="hwid.HwidService"/>

#### HwidService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetHwId | [HwIdRequest](#hwid.HwIdRequest) | [HwIdResponse](#hwid.HwIdRequest) | Request the hardware identifier |

 <!-- end services -->

### Messages

<a name="hwid.HwIdRequest"/>

#### HwIdRequest







<a name="hwid.HwIdResponse"/>

#### HwIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Identifier | [string](#string) |  | Hardware identifier |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="locale-service-api">Locale Service API (locale/locale.proto)</a>

### Services

<a name="locale.LocaleClient"/>

#### LocaleClient


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LocaleSet | [LocaleMessage](#locale.LocaleMessage) | [Empty](#locale.LocaleMessage) |  |
| TimezoneSet | [TimezoneMessage](#locale.TimezoneMessage) | [Empty](#locale.TimezoneMessage) |  |

 <!-- end services -->

### Messages

<a name="locale.Empty"/>

#### Empty
Empty message






<a name="locale.LocaleMessage"/>

#### LocaleMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Locale | [string](#string) |  | Locale string |






<a name="locale.TimezoneMessage"/>

#### TimezoneMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Timezone | [string](#string) |  | Timezone |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="socket-proxy-service-api">Socket Proxy Service API (socket/socket.proto)</a>

### Services

<a name="socketproxy.SocketStream"/>

#### SocketStream


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TransferData | [BytePacket](#socketproxy.BytePacket) | [BytePacket](#socketproxy.BytePacket) | Initiate bi-directional socket stream |

 <!-- end services -->

### Messages

<a name="socketproxy.BytePacket"/>

#### BytePacket
Data package in byte format


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Data | [bytes](#bytes) |  | Data bytes |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="stats-service-api">Stats Service API (stats/stats.proto)</a>

### Services

<a name="stats.StatsService"/>

#### StatsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetStats | [StatsRequest](#stats.StatsRequest) | [StatsResponse](#stats.StatsRequest) | Get process statistics |

 <!-- end services -->

### Messages

<a name="stats.LoadStats"/>

#### LoadStats
Load stats over different time frames


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Load1Min | [float](#float) |  |  |
| Load5Min | [float](#float) |  |  |
| Load15Min | [float](#float) |  |  |






<a name="stats.MemoryStats"/>

#### MemoryStats
Memory stats info


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Total | [uint64](#uint64) |  |  |
| Free | [uint64](#uint64) |  |  |
| Available | [uint64](#uint64) |  |  |
| Cached | [uint64](#uint64) |  |  |






<a name="stats.ProcessStat"/>

#### ProcessStat
Statistics for a singular process


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| User | [float](#float) |  |  |
| Sys | [float](#float) |  |  |
| ResSetSize | [uint64](#uint64) |  |  |






<a name="stats.ProcessStats"/>

#### ProcessStats
Cummulative process stats


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| CpuProcesses | [ProcessStat](#stats.ProcessStat) | repeated |  |
| MemProcesses | [ProcessStat](#stats.ProcessStat) | repeated |  |
| Total | [uint64](#uint64) |  |  |
| Running | [uint64](#uint64) |  |  |






<a name="stats.StatsRequest"/>

#### StatsRequest
Empty stats request message






<a name="stats.StatsResponse"/>

#### StatsResponse
Combined stats response message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Memory | [MemoryStats](#stats.MemoryStats) |  |  |
| Load | [LoadStats](#stats.LoadStats) |  |  |
| Process | [ProcessStats](#stats.ProcessStats) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="systemd-service-api">Systemd Service API (systemd/systemd.proto)</a>

### Services

<a name="systemd.UnitControlService"/>

#### UnitControlService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartApplication | [AppUnitRequest](#systemd.AppUnitRequest) | [UnitResponse](#systemd.AppUnitRequest) | Start remote application (session service) |
| StartUnit | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Start systemd unit (any, if whitelisted) |
| StopUnit | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Stop systemd unit (any, if whitelisted) |
| KillUnit | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Kill systemd unit (any, if whitelisted) |
| FreezeUnit | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Freeze/pause systemd unit (any, if whitelisted) |
| UnfreezeUnit | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Unfreeze/resume systemd unit (any, if whitelisted) |
| GetUnitStatus | [UnitRequest](#systemd.UnitRequest) | [UnitResponse](#systemd.UnitRequest) | Get systemd unit status (any, if whitelisted) |
| MonitorUnit | [UnitResourceRequest](#systemd.UnitResourceRequest) | [UnitResourceResponse](#systemd.UnitResourceRequest) | Obsolete monitoring function |

 <!-- end services -->

### Messages

<a name="systemd.AppUnitRequest"/>

#### AppUnitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UnitName | [string](#string) |  | Systemd unit name of the application |
| Args | [string](#string) | repeated | Application arguments |






<a name="systemd.UnitRequest"/>

#### UnitRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UnitName | [string](#string) |  | Full systemd unit name |






<a name="systemd.UnitResourceRequest"/>

#### UnitResourceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UnitName | [string](#string) |  |  |






<a name="systemd.UnitResourceResponse"/>

#### UnitResourceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu_usage | [double](#double) |  |  |
| memory_usage | [float](#float) |  |  |






<a name="systemd.UnitResponse"/>

#### UnitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| UnitStatus | [UnitStatus](#systemd.UnitStatus) |  |  |






<a name="systemd.UnitStatus"/>

#### UnitStatus
Systemd Unit Status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  | Full systemd unit name |
| Description | [string](#string) |  | A short human readable title of the unit |
| LoadState | [string](#string) |  | LoadState contains a state value that reflects whether the configuration file of this unit has been loaded |
| ActiveState | [string](#string) |  | ActiveState contains a state value that reflects whether the unit is currently active or not |
| SubState | [string](#string) |  | SubState encodes more fine-grained states that are unit-type-specific |
| Path | [string](#string) |  | Bus object path of the unit |
| FreezerState | [string](#string) |  | Freezer sub-state, indicates whether unit is frozen |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->




## <a id="wifi-service-api">WiFi Service API (wifi/wifi.proto)</a>

### Services

<a name="wifimanager.WifiService"/>

#### WifiService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListNetwork | [WifiNetworkRequest](#wifimanager.WifiNetworkRequest) | [WifiNetworkResponse](#wifimanager.WifiNetworkRequest) | List wifi networks |
| GetActiveConnection | [EmptyRequest](#wifimanager.EmptyRequest) | [AccessPoint](#wifimanager.EmptyRequest) | Retrieve all active connections |
| ConnectNetwork | [WifiConnectionRequest](#wifimanager.WifiConnectionRequest) | [WifiConnectionResponse](#wifimanager.WifiConnectionRequest) | Connect to a wifi network |
| DisconnectNetwork | [EmptyRequest](#wifimanager.EmptyRequest) | [WifiConnectionResponse](#wifimanager.EmptyRequest) | Disconnect from wifi network |
| TurnOn | [EmptyRequest](#wifimanager.EmptyRequest) | [WifiConnectionResponse](#wifimanager.EmptyRequest) | Turn wifi on |
| TurnOff | [EmptyRequest](#wifimanager.EmptyRequest) | [WifiConnectionResponse](#wifimanager.EmptyRequest) | Turn wifi off |

 <!-- end services -->

### Messages

<a name="wifimanager.AccessPoint"/>

#### AccessPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Connection | [bool](#bool) |  |  |
| SSID | [string](#string) |  |  |
| Signal | [uint32](#uint32) |  |  |
| Security | [string](#string) |  |  |






<a name="wifimanager.EmptyRequest"/>

#### EmptyRequest







<a name="wifimanager.WifiConnectionRequest"/>

#### WifiConnectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SSID | [string](#string) |  |  |
| Password | [string](#string) |  |  |
| Settings | [string](#string) |  |  |






<a name="wifimanager.WifiConnectionResponse"/>

#### WifiConnectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Response | [string](#string) |  |  |






<a name="wifimanager.WifiNetworkRequest"/>

#### WifiNetworkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| NetworkName | [string](#string) |  |  |






<a name="wifimanager.WifiNetworkResponse"/>

#### WifiNetworkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| networks | [AccessPoint](#wifimanager.AccessPoint) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->



### Scalar Value Types

| .proto Type | Notes | Go Type |
| ----------- | ----- | -------- |
| <a name="double" /> double |  | float64 |
| <a name="float" /> float |  | float32 |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 |
| <a name="bool" /> bool |  | bool |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | []byte |

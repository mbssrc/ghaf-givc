<!--
    Copyright 2025 TII (SSRC) and the Ghaf contributors
    SPDX-License-Identifier: CC-BY-SA-4.0
-->
# GIVC Security

## Scope

<!--

### Code/Data at rest

In general, GIVC code and configuration data should be integrity protected by the VMs nix store, which must be read-only. This only prevents non-root attacks, as root is able to remount the store with rw permissions. Additionally no non-root user should have access to the nix toolchain (trusted-user configuration), unless specifically assessed and subsequent risks are accepted.

An attacker should not be able to create persistance (even with root privileges) if the VMs nix store is read-only protected from host side. Additional integrity protections on the host should be implemented to protect the hosts store, and should be enhanced with secure boot integrity as well as security mechanisms to detect potential modifications.

Due to the nature of the nix store, no confidential information should be stored in the nix store, unless specifically confidentiality protected with an additional encryption mechanism.

User data is not stored or administrated at rest within GIVC. The components however may process user data in the form of file paths, or directly through the dbus proxy.

Code runtime protection is generally out of scope for this threat model.

An exception is code modification or execution through the GIVC components themselves, i.e., an attacker mis-using GIVC functionality to alter its execution logic or execute malicious code. We capture this as a generic asset. Further, services provided with GIVC code, may it be helper services or the components themselves, should be protected with the principle of least privileges and sandboxing mechanisms to prevent misuse.

Misuse of legitimate functionality...

 -->

## Trust Zones (Software Stack)

Trust zones in this project are the different areas within the software stack. They define how trustworthy an area is, based on how accessible it is: the more accessible, the less trustworthy.

|  ID |  Name | Description  | Scope  |  Risk  |
| -------- | -------- | ------------------------------------------------------------------------ |----|----|
|  TZ-1 | Network Layer | TCP/IP, UDS, VSOCK  | This TZ is considered partially out of scope - the generic VM network stack is not considered, whereas the GIVC-specific network stack is part of the scope. For this TM, the generic stack is considered potentially compromised. | 100 |
|  TZ-2 | GRPC Layer  | protobuf + implementation | The protocol buffer stack implementation is out of scope, but the GIVC specific usage and configuration is considered. This layer is responsible for the transport authentication and traffic authorization, which is exposed to the network. | 90  |
|  TZ-3 | Transport Layer  | givc transport layer | The GIVC transport stack is the translation layer from GRPC to the application, and is in scope. It is guarded by the GRPC authentication and authorization layer and is therefore less exposed to the network. | 50  |
|  TZ-4 | Business Layer  | givc bl layer | The application/business layer implements the GIVC components logic and is exposed to any legitimate party. Its exposure value is higher than the Transport Layer due to the assumption that an attacker has sucessfully (by)passed the authentication/authorization, e.g., by using legitimate interfaces such as the CLI application. | 80  |
|  TZ-5 | System Layer  | givc service layer | The system layer is the GIVC components systemd service, and potential subsequent access to the VM. It is not directly exposed, but security controls may become important if a vulnerability in another layer is exploited. | 30  |

## Assumptions

| ID       | Name | Description                                                                                               |
| -------- | ----- | --------------------------------------------------------------------------------------------------------- |
| AS-1     | Host Security | It is assumed that the host is not compromised. A compromised host leads to disclosure of all secrets and full system access. This includes the management of key/certificate files that are managed by the host. Thus, host to VM attacks are not part of this threat model. |
| AS-2     | Software Dependencies | All software inputs and dependencies such as flake inputs or Rust/Go libraries are out of scope for this project, as well as any system the components are embedded in. While the utilized software inputs are required to be updated to latest releases, no guarantees for their security can be given. This does NOT include security relevant configuration parameters of these components. |
| AS-3    | Hardware Dependencies | No hardware threats, including software-based hardware attacks such as micro-architectural attacks, side channel or fault attacks are considered as part of the scope.  |
| AS-4    | Application/Service Threats | No guarantees for any application or service that is administered within GIVC is part of the scope. Application or services must be reviewed individually to assess whether the security controls for administration, input parameters, and filters are sufficient to protect them and the system. |
| AS-5    | Usage and Configuration  | Several features can impact the security of your system, such as disabling security features or enabling certain debug features. A security analysis is advised to determine the impact of the usage of GIVC for any system. |
| AS-6    | VM Security |  |
| AS-7    | Supply Chain Security | ... |

## Attacker Model

For the purpose of this threat model, the different attackers are reflected in different categories based on their access level as defined in the [Trust Zones](#trust-zones).

| ID       | Name  | Description                                                                                               |
| -------- | ----- | --------------------------------------------------------------------------------------------------------- |
| AT-1     | TZ-1 Attacker | Attacker that has access to the internal TCP/IP network only - no access to other layers is assumed. |
| AT-2     | TZ-2 Attacker | Attacker that has access to network (TZ-1) and GRPC layer (TZ-2). This attacker can perform GRPC queries, but cannot authenticate to any of the components. |
| AT-3     | TZ-3 Attacker | Attacker that has access to the Transport Layer (TZ-3) and all previous layers. This attacker has passed the GRPC authentication and authorization (either legitimately or not), and can perform RPCs. We do not distinguish between levels of access to particular funtionality at this stage. |
| AT-4     | TZ-4 Attacker (non-root) | Attacker that has sucessfully compromised a GIVC agent with non-elevated privileges, i.e., an application VM agent. |
| AT-5     | Service Attacker (root) | An internal attacker that has sucessfully compromised a GIVC agent with elevated privileges, i.e., a system VM. This excludes the host, as per assumption AS-1 (Host Security). |
| AT-6     | Admin Service Attacker | An attacker that has compromised the admin service, and is in control of all GIVC functionality except host access. |
| AT-7     | External Attacker (for future use) | An attacker that has remote access to a machine running GIVC funcionality, i.e., a (remote) attacker that has unauthenticated/unauthorized access to the GIVC remote admin interface. |
| AT-8     | Admin Attacker (for future use) | An attacker that has remote admin access to GIVC funcionality, i.e., a (remote) system administrator that has access to the GIVC remote admin interface. A local system administrator is not considered, as that depends on system configuration (they likely have root access to the system). |


## Assets


Asset objectives are denoted as
C - Confidentiality
I - Integrity
A - Availability
A/A - Authentication/Authorization


|  ID |  Name | Description  | Objective  |  Risk  |
| -------- | -------- | -------------------------------------------------------- |----|----|
|  AS-GEN-1 | GIVC Source Code | GIVC source code integrity shall be covered by the respective mechanisms provided by the languages and Nix flake lock, and supply chain rules guarding the integrity of the code. | I | 40 |
|  AS-GEN-2 | GIVC Configuration Data | The integrity of GIVC configuration data shall be enforced. This includes all configuration data that stems from nix configurations and external files that have an impact on the security of the GIVC components. Ultimately, it relies on the security and protection of the system that uses GIVC. | I | 40 |
|  AS-GEN-3 | GIVC Code | The integrity of GIVC code shall be enforced. This includes code integrity at rest as well as during loading and runtime. Ultimately, it relies on the security and protection of the system that uses GIVC. | I | 40 |
|  AS-GEN-4 | GIVC Component Availability | GIVC components shall be highly available for use. While there are of course huge implications of the system using GIVC, the implementation must implement measures to decrease attack surface. | A | 40 |
|  AS-GEN-5 | GIVC Component Data | Data communicated between via GIVC components shall be integrity protected and confidential. | CI | 100 |
|  AS-TLS-1 | TLS CA key | The TLS Certificate Authority (CA) keys are used as to sign component certificates for TLS  protection, and must be protected from usage, disclosure, and/or modification. | CI | 100 |
|  AS-TLS-2 | TLS Keys | The TLS keys are used to authenticate the GIVC hosts/components, and must be protected from disclosure or modification. | CI | 40 |
|  AS-TLS-3 | TLS Certificates | The TLS certificates must be protected from modification. | I | 80 |
|  AS-TLS-4 | TLS Server interface | GRPC server connections must be authenticated before usage. | (A/-) | 40 |
|  AS-RPC-1 | RPC CA Key  | The RPC Certificate Authority (CA) keys are used to sign certificates for RPC token generation and verification, and must be protected from usage, disclosure, and/or modification. | CI | 100 |
|  AS-RPC-2 | RPC keys  | The RPC keys are used to authenticate RPC access, and must be protected from disclosure or modification. | CI | 40 |
|  AS-RPC-3 | RPC certificates  | The RPC certiciates are used to verify RPC access tokens, and must be protected from modification. | I | 40 |
|  AS-RPC-4 | RPC interface | The RPC connections must be authorized before usage. | (-/A) | 40 |

<!-- |  AS-USR-1 | Arguments: File paths | File paths to user data must be valid and integrity protected. Paths are not considered confidential. | I | 40 |
|  AS-USR-2 | Arguments: Flag | Flag arguments must be valid and integrity protected. Flags are not considered confidential. | I | 40 |
|  ASSET-12 | Arguments: URL | URL arguments must be valid and integrity protected. URLs are not considered confidential. | I | 40 |
|  ASSET-13 | DBUS User Secrets | User data processed via the DBUS proxy functionality shall be confidentiality and integrity protected. | CI | 40 |
|  ASSET-x | System | ???  | CI | 40 | -->


## Threats

### Supply Chain

The repository security relies on github and the organization, and the software supply-chain security measures of the embedding project. This project provides the interfaces for integrity protection through Rust (Cargo.lock), Go (go.sum, vendor hash), and Nix's flake source protection (flake.lock).

Note that code dependencies are not reviewed or evaluated at this point.


### Network Layer

The network layer implementation is part of the embedding system.

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-NL-1 | Network Layer - Eavesdropping | An attacker may act as a MITM and eavesdrop traffic to extract confidential information. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-2 | Network Layer - Injection | An attacker may act as a MITM and modify traffic. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-3 | Network Layer - Replay | An attacker may replay recorded traffic. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-4 | Network Layer - Relay | An attacker may relay traffic, e.g., across different interface attacks. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-5 | Network Layer - Interface Attacks | An attacker may perform logical cross-interface attacks, misusing multiple connections via different network interfaces. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-6 | Network Layer -  | An attacker may perform DOS attacks on the network. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-NL-7 | Network Layer - Spoofing | An attacker may perform impersonation attacks on the network level, e.g., via address spoofing. | tbd | AS-GEN-5 | 19 | 29 | tbc |

### GRPC Layer

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-GL-1 | GRPC - TLS Bypass | An attacker may bypass TLS authentication through various attacks, e.g., downgrade attacks. | tbd | AS-TLS-2,AS-TLS-3,AS-TLS-4,AS-GEN-5  | 19 | 29 | tbc |
| T-GL-2 |  | An attacker may use the GRPC reflection API to investingate exposed services. | tbd | - | 19 | 29 | tbc |
| T-GL-3 | GRPC - Authentication Bypass | An attacker may bypass the TLS authentication. | tbd | AS-TLS-4 | 19 | 29 | tbc |
| T-GL-4 | GRPC - Authorization Bypass | An attacker may bypass the GRPC API authorization. | tbd | AS-RPC-4 | 19 | 29 | tbc |
| T-GL-5 | GRPC - Unauthenticated DOS | An attacker may use the GRPC API to perform DOS attacks without authentication/authorization. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-GL-6 | GRPC - Authenticated DOS | An attacker may use available GRPC API calls to perform a DOS attack. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-x | GRPC - Spoofing | An attacker may use stolen TLS/RPC credentials to perform impersonation attacks or expand functionality access. | tbd | tbd | 19 | AS-TLS-4,AS-RPC-4 | tbc |

### Transport Layer

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-TL-1 | Transport Layer - RCE | An attacker may misuse the Transport API to achieve remote code execution. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-TL-2 | Transport Layer - DOS | An attacker may misuse the Transport API to perform DOS attacks. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-TL-3 | Transport Layer - Corruption | An attacker may misuse the Transport API to perform memory corruption attacks by misusing legitimate data fields. | tbd | AS-GEN-3 | 19 | 29 | tbc |

### Business Layer

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-BL-GEN-1 | Business Layer - Memory Corruption | An attacker may use specifically crafted messages to perform memory corruption. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-GEN-2 | Business Layer - RCE | An attacker may use specifically crafted messages to perform remote code execution. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-GEN-3 | Business Layer - DOS | An attacker may use the APIs to perform DOS attacks on an application. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-BL-GEN-4 | Business Layer - Debug | An attacker may use the debug features of the APIs for data exfiltration. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-BL-GEN-5 | Business Layer - Logging | An attacker may cover their attacks by supressing logging functionality. | tbd | - | 19 | 29 | tbc |

#### Agent Module

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-BL-AGT-APP-1 | URL args - File Disclosure | An attacker may use the URL flag of the applicaton API to reveal files on a remote machine. | tbd | AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-2 | URL args - Malicious Links | An attacker may use the URL flag of the applicaton API to open malicious links on a remote machine. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-3 | URL args - Shellcode | An attacker may use the URL flag of the applicaton API to inject shell code on a remote machine. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-4 | Flag args - Shellcode | An attacker may use flags of the applicaton API to inject shell code on a remote machine. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-5 | Flag args - Vulnerable Config | An attacker may use flags of the applicaton API to start an application with vulnerable configurations or otherwise weaken its security posture. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-7 | File args - Shellcode | An attacker may use the file argument of the applicaton API to inject shell code on a remote machine. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-8 | File args - Vulnerable File | An attacker may use the file argument of the applicaton API to start an application with vulnerable files. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-APP-9 | File args - Illegitimate Path | An attacker may use the file argument of the applicaton API to start an application with illegitimate file paths, to reveal information on the remote machine. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-BUS-1 | DBUS - Service Access | An attacker may use the DBUS proxy to access system services that are not intended to be accessed. | tbd | - | 19 | 29 | tbc |
| T-BL-AGT-BUS-2 | DBUS - Functionality Access | An attacker may use the DBUS proxy to access functionality of an allowed system service that is not intended to be accessed. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-BUS-3 | DBUS - RCE | An attacker may use the DBUS proxy to achieve remote code execution in exposed services. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-BUS-4 | DBUS - Exfiltration | An attacker may use the DBUS proxy to extract confidential information from the exposed services. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-BUS-5 | DBUS - Service Impersonation | An attacker may use the DBUS proxy to impersonate a legitimate service at the remote end. This may lead to information disclosure or service disruption. | tbd | AS-GEN-3,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-AGT-LTZ-1 | TZ/Locale - Corruption | An attacker may use malformed timezone or locale strings to corrupt the function of applications and services. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-AGT-LTZ-2 | TZ/Locale - Time Attacks | An attacker may use the timezone API to change the timezone for malicious purposes. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-AGT-ST-1 | Stats - Disclosure | An attacker may use the stats API to monitor processes on a remote machine, e.g., to debug remote exploits and attacks. | tbd | tbd | 19 | 29 | tbc |
| T-BL-AGT-SYS-1 | Systemd - Arbitrary Commands | An attacker may use the systemd API to administrate arbitrary systemd units including services, targets, or other units. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-AGT-SYS-2 | Systemd - Disruption | An attacker may use legitimately exposed units by the systemd API to disrupt the function of the service and/or system, including DOS. | tbd | AS-GEN-3,AS-GEN-4 | 19 | 29 | tbc |

#### Admin Module

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-BL-ADM-REG-1 | Registry - Rouge Entity | An attacker may register a rouge entity (VM, agent, system, or application service) with the admin service, which may lead to disclosure of information and/or disruptions. | tbd | AS-GEN-4,AS-GEN-5 | 19 | 29 | tbc |
| T-BL-ADM-REG-2 | Registry - Poisoning | An attacker may use the registration function to poison the admin registry. | tbd | AS-GEN-3 | 19 | 29 | tbc |
| T-BL-ADM-PRX-1 | Proxy - Resource Exhaustion | An attacker may use admin proxy commands to exhaust system resources, e.g., by repeatably rebooting VMs. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-BL-ADM-PRX-2 | Proxy - System Control | An attacker may use system commands to reboot, suspend, or shutdown the system. | tbd | AS-GEN-3,AS-TLS-4 | 19 | 29 | tbc |
| T-BL-ADM-MON-1 | Monitoring - Disruption | An attacker may disrupt the connection of a legitimate entity (network or else) to force de-registration from the admin registry to disrupt system functionality. | tbd | AS-GEN-4 | 19 | 29 | tbc |
| T-BL-ADM-MON-2 | Monitoring - Discovery | An attacker may perform systemd discovery by using the system monitoring commands. | tbd | - | 19 | 29 | tbc |

### System Layer

System Layer threats only incude potential threats that are directly applicable to the GIVC configurations. General system
attacks are out of scope, as possible configurations make the assumption space unmanagable.

| ID | Name | Description | Category | Assets | Likelihood | Impact | Comments |
| -- | ---- | ----------- | -------- | --- | ---------- | ------ | -------- |
| T-SL-1 | System Layer - Service Modification | An attacker may alter the GIVC systemd service configuration, leading to misconfigurations of services. | tbd | AS-GEN-3,AS-GEN-4,AS-GEN-5 | 19 | 29 | tbc |
| T-SL-2 | System Layer - DNS Poisoning | An attacker may poison the local DNS to re-route traffic. | tbd | AS-TLS-4 | 19 | 29 | tbc |
| T-SL-3 | System Layer - Lateral Movement | An attacker may attack the local system after achieving code execution in the context of a GIVC service. | tbd | - | 19 | 29 | tbc |
| T-SL-4 | System Layer - Extended Privileges | An attacker may use GIVC functionality after taking over a GIVC service, e.g., to escape a VM or extend compromise reach. | tbd | - | 19 | 29 | tbc |
| T-SL-5 | System Layer - Privilege Escalation | An attacker may use GIVC functionality to elevate local privileges, or perform single actions with elevated privileges (e.g., copy data). | tbd | - | 19 | 29 | tbc |
| T-SL-6 | System Layer - Credential Access | An attacker may access GIVC credentials (keys, certificates) stored in the embedding system which are highly relevant to the security. | tbd | AS-TLS-1,AS-TLS-2,AS-TLS-3,AS-RPC-1,AS-RPC-2,AS-RPC-3 | 19 | 29 | tbc |

## Data Flows

| ID |  Name | Description  | bidirectional  |  Source  | Destination | Assets | Threats |
| -------- | -------- | ---------------- | ---------------- | ---------------- | ---------------- | | ---------------- |---------------- |
| ID |  Name | Description  | bidirectional  |  Source  | Destination | Assets | Threats |


## Security Controls

### Mutual TLS

Mutual TLS is enforced between all components.

Authentication
TLS keys and certificates are used for both client and server authentication.


**Summary**

| ID | Name | Description | Category | Comments |
| -- | ---- | ----------- | -------- | -------- |
| SC-1 | Mutual TLS | Enforcement of mutual TLS connection between GIVC components. | Transport Security |  |


key storage
sandboxing
ACL generation
token security
input validation



## Security Requirements


### Internal Requirements

| ID | Name | Description | |
| -- | ---- | ----------- | |
| T-SS-1 | Code Sign-off | All code contributions must be signed-off by the contributor to allow attribution. |
| T-SS-2 | Code Review | All code contributions must be reviewed by at least one contributor with write access. Security implications should be highlighted and potentially included in the treat model. |
| T-SS-3 | Automated Tests | Code must be automatically tested, if possible. Currently, no standards for testing are applied.  |
| T-SS-4 | Automated Dependency Updates | Code dependencies must be kept reasonably up-to-date. |
| T-SS-5 | Code Tracability | All code must be traceable to their origin utilizing the language specific features. |

### External Requirements
...
<!--
Due to GIVCs flexibility, several components maybe running in different virtual machines. For the purpopse of this TM, we assign the following trust zones:

|  ID |  Name | Description  | GIVC Component  |  Risk  |
| -------- | -------- | -------- |----|----|
|  TZ-1 | NetVM  | Networking enabled VM, manages hardware network interfaces and external connections. | SysVM Agent | 100  |
|  TZ-2 | AppVM  | Isolated application VM. Manages one or more internet facing applications. | AppVM Agent | 60   |
|  TZ-3 | GuiVM  | Desktop VM that manages the UI, user interface devices, and non-isolated applications. | SysVM Agent / CLI | 70  |
|  TZ-4 | PeripheralVM | Peripheral VM that manages the audio and bluetooth stacks. Not internet facing but connected to most VMs. | SysVM Agent | 40  |
|  TZ-5 | AdminVM  | Isolated admin VM that manages the GIVC admin and other admin services. Mainanence and logging services are constantly available and exposed to the internet. | Admin Service | 90  |
|  TZ-6 | Host  | Desktop VM that manages the UI, user interface devices, and non-isolated applications. Not directly internet facing, and can only be reached via AdminVM. | Host Agent | 30  |

This example setup highlights that the exposure of the components is directly related to the location within a platform. Accordingly, each use case requires it's individual analysis. For example, the SysVM component may be used in an offline VM that is not exposed to the internet or any other external channel. -->
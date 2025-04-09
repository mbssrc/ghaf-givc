<!--
    Copyright 2025 TII (SSRC) and the Ghaf contributors
    SPDX-License-Identifier: CC-BY-SA-4.0
-->
<style>
img[src*='#center'] {
    display: block;
    float: none;
    margin-left: auto;
    margin-right: auto;
    width: 50%;
}
</style>

# TII SSRC Secure Technologies: Ghaf gRPC inter-vm communication framework

[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-darkgreen.svg)](https://github.com/tiiuae/ghaf/tree/main/LICENSES/Apache-2.0.txt)  [![License: CC-BY-SA 4.0](https://img.shields.io/badge/License-CC--BY--SA--4.0-orange.svg)](https://github.com/tiiuae/ghaf/tree/main/LICENSES/CC-BY-SA-4.0.txt)

![Ghaf/GRPC Inter-Vm Communication](docs/docs/images/givc.png#center)

This is the repository for the gRPC-based control channel of the [Ghaf Framework](https://github.com/tiiuae/ghaf).
The GIVC (Ghaf/gRPC Inter-Vm Communication) framework is a collection of services to to administrate and control virtual machines, their services, and applications.

This project was started to add a control channel for a virtualized system where hardware related system functionality (such as network and graphics) and applications are isolated into separate virtual machines. The objective is to provide inter-vm communication using a unified framework with gRPC over multiple channels. NixOS modules and packages are provided through the flake.

## Documentation

The GIVC documentation site is located at https://tiiuae.github.io/givc/. To build the automated API documentation locally, use:
```nix
nix build .#docs
```

The Ghaf Framework documentation site is located at https://tiiuae.github.io/ghaf/. It is under cooperative development.

## Licensing

The Ghaf team uses several licenses to distribute software and documentation:

| License Full Name | SPDX Short Identifier | Description |
| -------- | ----------- | ----------- |
| Apache License 2.0 | [Apache-2.0](https://spdx.org/licenses/Apache-2.0.html) | Ghaf source code. |
| Creative Commons Attribution Share Alike 4.0 International | [CC-BY-SA-4.0](https://spdx.org/licenses/CC-BY-SA-4.0.html) | Ghaf documentation. |

See [LICENSE.Apache-2.0](https://github.com/tiiuae/ghaf/tree/main/LICENSES/Apache-2.0.txt) and [LICENSE.CC-BY-SA-4.0](https://github.com/tiiuae/ghaf/tree/main/LICENSES/CC-BY-SA-4.0.txt) for the full license text.
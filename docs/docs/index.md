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

# The GIVC Framework

The GIVC (Ghaf/gRPC Inter-Vm Communication) framework is a collection of services to to administrate and control virtual machines, their services, and applications across the [Ghaf Framework](https://github.com/tiiuae/ghaf).

![Ghaf/GRPC Inter-Vm Communication](images/givc.png#center)

Albeit being efficient, GIVC as a control channel does not do heavy lifting - graphics, file transfers, and other high-bandwidth or low-latency tasks are implemented with more efficient technologies and out of scope for this framework.

## Technologies

The framework is written in Rust and Go, modern languages providing excellent features for microservices. It exports its
components as Nix modules and packages, described in the Nix language. As part of the [Ghaf Framework](https://github.com/tiiuae/ghaf), this project inherits many technology dependencies from Ghaf itself.

* [gRPC](https://grpc.io/) <br>
  The high performance and versatile RPC framework offers performant binary serialization, scalability, bi-directional streaming, and works across multiple platforms and languages.
* [NixOS](https://github.com/NixOS) <br>
  The purely-functional Linux distribution with over [120k packages](https://search.nixos.org/packages). GIVC exports NixOS modules and packages to allow integration with the Ghaf Framework.
* [microvm.nix](https://github.com/astro/microvm.nix) <br>
  The microvm.nix project is used to build and run NixOS VMs on several type-2 hypervisors on NixOS hosts. It provides wrappers for virtio devices, storage, network, and is the foundation to build and run virtual machines on Ghaf.
* NixOS Ecosystem <br>
  blah
## FAQ

**Why?** <br>
A custom framework offers the required flexibility to organize our virtual machines and their services!

**Q: How can I inspect the outputs?**
**A: To see all available outputs, run the `nix flake show` command.**

**Q: How can I ?**

**A: !**

**Q: ?**

**A: !**

## License

[![License: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-darkgreen.svg)](./LICENSE) <br>
Copyright 2025 TII (SSRC) and the Ghaf contributors
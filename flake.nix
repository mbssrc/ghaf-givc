# Copyright 2024 TII (SSRC) and the Ghaf contributors
# SPDX-License-Identifier: Apache-2.0
{
  description = "Go modules for inter-vm communication with gRPC.";

  # Inputs
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };
    flake-root = {
      url = "github:srid/flake-root";
    };
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    devshell = {
      url = "github:numtide/devshell";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    crane = {
      url = "github:ipetkov/crane";
    };
    pre-commit-hooks-nix = {
      url = "github:cachix/pre-commit-hooks.nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.nixpkgs-stable.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{
      self,
      flake-parts,
      crane,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
      ];
      imports = [
        ./nixos/checks/treefmt.nix
        ./nixos/checks/vmTests.nix
        ./devshell.nix
        ./nixos/tests
      ];

      perSystem =
        { pkgs, lib, ... }:
        {
          # Packages
          packages =
            let
              src = lib.fileset.toSource {
                root = ./.;
                fileset = lib.fileset.unions [
                  ./go.mod
                  ./go.sum
                  ./api
                  ./internal
                ];
              };
              givc-admin-rs = pkgs.callPackage ./nixos/packages/givc-admin-rs.nix {
                inherit crane;
                src = ./.;
              };
            in
            {
              inherit givc-admin-rs;
              givc-agent = pkgs.callPackage ./nixos/packages/givc-agent.nix { inherit src; };
              givc-cli = givc-admin-rs.cli;
            };
        };
      flake = {
        # NixOS Modules
        nixosModules = {
          admin = import ./nixos/modules/admin.nix { inherit self; };
          host = import ./nixos/modules/host.nix { inherit self; };
          sysvm = import ./nixos/modules/sysvm.nix { inherit self; };
          appvm = import ./nixos/modules/appvm.nix { inherit self; };
          dbus = import ./nixos/modules/dbus.nix { inherit self; };
        };

        # Overlays
        overlays.default = _final: prev: {
          givc-cli = self.packages.${prev.stdenv.hostPlatform.system}.givc-admin-rs.cli;
        };
      };
    };
}

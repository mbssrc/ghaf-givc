# Copyright 2024 TII (SSRC) and the Ghaf contributors
# SPDX-License-Identifier: Apache-2.0
{ self }:
{
  config,
  pkgs,
  lib,
  ...
}:
let
  cfg = config.givc.sysvm;
  inherit (self.packages.${pkgs.stdenv.hostPlatform.system}) givc-agent;
  inherit (lib)
    mkIf
    mkOption
    mkEnableOption
    types
    trivial
    strings
    lists
    concatStringsSep
    optionalString
    optionals
    ;
  inherit (builtins) toJSON;
  inherit (import ./definitions.nix { inherit config lib; })
    transportSubmodule
    proxySubmodule
    tlsSubmodule
    ;
in
{
  options.givc.sysvm = {
    enable = mkEnableOption "Enable givc-sysvm module.";

    agent = mkOption {
      description = ''
        Host configuration for the system VM.
      '';
      type = transportSubmodule;
    };

    services = mkOption {
      description = ''
        List of systemd services for the manager to administrate. Expects a space separated list.
        Should be a unit file of type 'service' or 'target'.
      '';
      type = types.listOf types.str;
      default = [
        "reboot.target"
        "poweroff.target"
      ];
    };

    debug = mkEnableOption "Enable verbose logs for debugging.";

    admin = mkOption {
      description = "Admin server configuration.";
      type = transportSubmodule;
    };

    wifiManager = mkOption {
      description = ''
        Wifi manager to handle wifi related queries.
      '';
      type = types.bool;
      default = false;
    };

    hwidService = mkOption {
      description = ''
        Hardware identifier service.
      '';
      type = types.bool;
      default = false;
    };

    hwidIface = mkOption {
      description = ''
        Interface for hardware identifier.
      '';
      type = types.str;
      default = "";
    };

    localeListener = mkOption {
      description = ''
        Locale handler.
      '';
      type = types.bool;
      default = false;
    };

    socketProxy = mkOption {
      description = ''
        Socket proxy module. If not provided, the module will not use a socket proxy.
      '';
      type = types.nullOr (types.listOf proxySubmodule);
      default = null;
    };

    tls = mkOption {
      description = "TLS configuration.";
      type = tlsSubmodule;
    };
  };

  imports = [ self.nixosModules.dbus ];

  config = mkIf cfg.enable {

    assertions = [
      {
        assertion = cfg.services != [ ];
        message = "A list of services (or targets) is required for this module to run.";
      }
      {
        assertion =
          !(cfg.tls.enable && (cfg.tls.caCertPath == "" || cfg.tls.certPath == "" || cfg.tls.keyPath == ""));
        message = ''
          The TLS configuration requires paths' to CA certificate, service certificate, and service key.
          To disable TLS, set 'tls.enable = false;'.
        '';
      }
      {
        assertion =
          cfg.socketProxy == null
          || lists.allUnique (map (p: (strings.toInt p.transport.port)) cfg.socketProxy);
        message = "SocketProxy: Each socket proxy instance requires a unique port number.";
      }
      {
        assertion = cfg.socketProxy == null || lists.allUnique (map (p: p.socket) cfg.socketProxy);
        message = "SocketProxy: Each socket proxy instance requires a unique socket.";
      }
    ];

    systemd.targets.givc-setup = {
      enable = true;
      description = "Ghaf givc target";
      bindsTo = [ "network-online.target" ];
      after = [ "network-online.target" ];
      wantedBy = [ "network-online.target" ];
    };

    systemd.services."givc-${cfg.agent.name}" = {
      description = "GIVC remote service manager for system VMs";
      enable = true;
      after = [ "givc-setup.target" ];
      partOf = [ "givc-setup.target" ];
      wantedBy = [ "givc-setup.target" ];
      serviceConfig = {
        Type = "exec";
        ExecStart = "${givc-agent}/bin/givc-agent";
        Restart = "always";
        RestartSec = 1;
      };
      path = [ pkgs.dbus ];
      environment = {
        "AGENT" = "${toJSON cfg.agent}";
        "DEBUG" = "${trivial.boolToString cfg.debug}";
        "TYPE" = "8";
        "SUBTYPE" = "9";
        "WIFI" = "${trivial.boolToString cfg.wifiManager}";
        "HWID" = "${trivial.boolToString cfg.hwidService}";
        "HWID_IFACE" = "${cfg.hwidIface}";
        "LOCALE_LISTENER" = "${trivial.boolToString cfg.localeListener}";
        "SOCKET_PROXY" = "${optionalString (cfg.socketProxy != null) (toJSON cfg.socketProxy)}";
        "PARENT" = "microvm@${cfg.agent.name}.service";
        "SERVICES" = "${concatStringsSep " " cfg.services}";
        "ADMIN_SERVER" = "${toJSON cfg.admin}";
        "TLS_CONFIG" = "${toJSON cfg.tls}";
      };
    };
    networking.firewall.allowedTCPPorts =
      let
        agentPort = strings.toInt cfg.agent.port;
        proxyPorts = optionals (cfg.socketProxy != null) (
          map (p: (strings.toInt p.transport.port)) cfg.socketProxy
        );
      in
      [ agentPort ] ++ proxyPorts;
  };
}

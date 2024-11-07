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
  cfg = config.givc.dbusproxy;
  inherit (lib)
    mkOption
    mkEnableOption
    mkIf
    types
    concatStringsSep
    concatMapStringsSep
    optionalString
    optionalAttrs
    ;

  # Dbus policy submodule
  policySubmodule = types.submodule {
    options = {
      see = mkOption {
        description = ''
          SEE policy:
            The name/ID is visible in the ListNames reply
            The name/ID is visible in the ListActivatableNames reply
            You can call GetNameOwner on the name
            You can call NameHasOwner on the name
            You see NameOwnerChanged signals on the name
            You see NameOwnerChanged signals on the ID when the client disconnects
            You can call the GetXXX methods on the name/ID to get e.g. the peer pid
            You get AccessDenied rather than NameHasNoOwner when sending messages to the name/ID
        '';
        type = types.nullOr (types.listOf types.str);
        default = null;
      };
      talk = mkOption {
        description = ''
          TALK policy:
            You can send any method calls and signals to the name/ID
            You will receive broadcast signals from the name/ID (if you have a match rule for them)
            You can call StartServiceByName on the name
        '';
        type = types.nullOr (types.listOf types.str);
        default = null;
      };
      own = mkOption {
        description = ''
          OWN policy:
            You are allowed to call RequestName/ReleaseName/ListQueuedOwners on the name
        '';
        type = types.nullOr (types.listOf types.str);
        default = null;
      };

      call = mkOption {
        description = ''
          CALL policy:
            You can call the specific methods

            From xdg-dbus-proxy manual:

            The RULE in these options determines what interfaces, methods and object paths are allowed. It must be of the form [METHOD][@PATH], where METHOD
            can be either '*' or a D-Bus interface, possible with a '.*' suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a '/*' suffix.
        '';
        type = types.nullOr (types.listOf types.str);
        default = null;
      };
      broadcast = mkOption {
        description = ''
          BROADCAST policy:
            You can receive broadcast signals from the name/ID

            From xdg-dbus-proxy manual:

            The RULE in these options determines what interfaces, methods and object paths are allowed. It must be of the form [METHOD][@PATH], where METHOD
            can be either '*' or a D-Bus interface, possible with a '.*' suffix, or a fully-qualified method name, and PATH is a D-Bus object path, possible with a '/*' suffix.
        '';
        type = types.nullOr (types.listOf types.str);
        default = null;
      };
    };
  };

  # Dbus component submodule
  dbusSubmodule = types.submodule {
    options = {
      enable = mkEnableOption "Enable the dbus component";
      user = mkOption {
        description = ''
          User to run the xdg-dbus-proxy service as. This option must be set to allow a remote user to connect to the bus.
          Defaults to `root`.
        '';
        type = types.str;
        default = "root";
      };
      socket = mkOption {
        description = "Socket path used to connect to the bus. Defaults to `/tmp/.dbusproxy.sock`.";
        type = types.str;
        default = "/tmp/.dbusproxy.sock";
      };
      policy = mkOption {
        description = ''
          Policy submodule for the dbus proxy.

          Filtering is applied only to outgoing signals and method calls and incoming broadcast signals. All replies (errors or method returns) are allowed once for an outstanding method call, and never otherwise.
          If a client ever receives a message from another peer on the bus, the senders unique name is made visible, so the client can track caller lifetimes via NameOwnerChanged signals. If a client calls a method on
          or receives a broadcast signal from a name (even if filtered to some subset of paths or interfaces), that names basic policy is considered to be (at least) TALK, from then on.
        '';
        type = policySubmodule;
      };
    };
  };

in
{
  options.givc.dbusproxy = {
    enable = mkEnableOption ''
      Enables givc-dbusproxy module. This module is a wrapper for the `xdg-dbus-proxy`, and configures systemd services for
      the system and/or session bus. The respective service is enabled if a policy for the bus is set.

      Filtering is enabled by default, and the config requires at least one policy value (see/talk/own) to be set. For more
      details, please refer to the xdg-dbus-proxy manual (e.g., https://www.systutorials.com/docs/linux/man/1-xdg-dbus-proxy/).

      Policy values are a list of strings, where each string is translated into the respective argument. Multiple instances
      of the same value are allowed.

      Example:
      ```
      config.givc.dbusproxy.system.policy = {
        see = [ "org.freedesktop.NetworkManager.*" "org.freedesktop.Avahi.*" ];
        talk = [ "org.freedesktop.NetworkManager.*" ];
      };
      ```
      In order to create your policies, consider using `busctl` to list the available services and their properties.
    '';
    system = mkOption {
      description = "Configuration of givc-dbusproxy for system bus.";
      type = dbusSubmodule;
      default = { };
    };
    session = mkOption {
      description = "Configuration of givc-dbusproxy for user session bus.";
      type = dbusSubmodule;
      default = { };
    };
  };

  config = mkIf cfg.enable {
    assertions = [
      {
        assertion = cfg.enable -> (cfg.system.enable || cfg.session.enable);
        message = ''
          The DBUS proxy module requires at least one of the system or session bus to be enabled.
        '';
      }
      {
        assertion =
          cfg.system.enable
          -> (
            cfg.system.policy.see != null
            || cfg.system.policy.talk != null
            || cfg.system.policy.own != null
            || cfg.system.policy.call != null
            || cfg.system.policy.broadcast != null
          );
        message = ''
          At least one policy value (see/talk/own/call/broadcast) for the system bus must be set. For more information, please
          refer to the xdg-dbus-proxy manual (e.g., https://www.systutorials.com/docs/linux/man/1-xdg-dbus-proxy/).
        '';
      }
      {
        assertion =
          cfg.session.enable
          -> (
            cfg.session.policy.see != null
            || cfg.session.policy.talk != null
            || cfg.session.policy.own != null
            || cfg.session.policy.call != null
            || cfg.session.policy.broadcast != null
          );
        message = ''
          At least one policy value (see/talk/own/call/broadcast) for the session bus must be set. For more information, please
          refer to the xdg-dbus-proxy manual (e.g., https://www.systutorials.com/docs/linux/man/1-xdg-dbus-proxy/).
        '';
      }
      {
        assertion =
          cfg.session.enable
          -> (
            config.users.users.${cfg.session.user}.isNormalUser
            && config.users.users.${cfg.session.user}.uid != null
          );
        message = ''
          You need to specify a non-system user with UID set to run the session bus proxy.
        '';
      }
    ];

    environment.systemPackages = [
      pkgs.xdg-dbus-proxy
    ];

    systemd =

      optionalAttrs cfg.system.enable {
        services."givc-dbusproxy-system" =
          let
            args = concatStringsSep " " [
              "${optionalString (cfg.system.policy.see != null) (
                concatMapStringsSep " " (x: "--see=${x}") cfg.system.policy.see
              )}"
              "${optionalString (cfg.system.policy.talk != null) (
                concatMapStringsSep " " (x: "--talk=${x}") cfg.system.policy.talk
              )}"
              "${optionalString (cfg.system.policy.own != null) (
                concatMapStringsSep " " (x: "--own=${x}") cfg.system.policy.own
              )}"
              "${optionalString (cfg.system.policy.call != null) (
                concatMapStringsSep " " (x: "--call=${x}") cfg.system.policy.call
              )}"
              "${optionalString (cfg.system.policy.broadcast != null) (
                concatMapStringsSep " " (x: "--broadcast=${x}") cfg.system.policy.broadcast
              )}"
            ];
          in
          {
            description = "GIVC local xdg-dbus-proxy system service";
            enable = true;
            before = [ "givc-setup.target" ];
            wantedBy = [ "givc-setup.target" ];
            serviceConfig = {
              Type = "exec";
              ExecStart = "${pkgs.xdg-dbus-proxy}/bin/xdg-dbus-proxy unix:path=/run/dbus/system_bus_socket ${cfg.system.socket} --filter ${args}";
              Restart = "always";
              RestartSec = 1;
              User = cfg.system.user;
            };
          };
      }
      // optionalAttrs cfg.session.enable {
        user.services."givc-dbusproxy-session" =
          let
            args = concatStringsSep " " [
              "${optionalString (cfg.session.policy.see != null) (
                concatMapStringsSep " " (x: "--see=${x}") cfg.session.policy.see
              )}"
              "${optionalString (cfg.session.policy.talk != null) (
                concatMapStringsSep " " (x: "--talk=${x}") cfg.session.policy.talk
              )}"
              "${optionalString (cfg.session.policy.own != null) (
                concatMapStringsSep " " (x: "--own=${x}") cfg.session.policy.own
              )}"
              "${optionalString (cfg.session.policy.call != null) (
                concatMapStringsSep " " (x: "--call=${x}") cfg.session.policy.call
              )}"
              "${optionalString (cfg.session.policy.broadcast != null) (
                concatMapStringsSep " " (x: "--broadcast=${x}") cfg.session.policy.broadcast
              )}"
            ];
            uid = toString config.users.users.${cfg.session.user}.uid;
          in
          {
            description = "GIVC local xdg-dbus-proxy session service";
            enable = true;
            after = [ "sockets.target" ];
            wants = [ "sockets.target" ];
            wantedBy = [ "default.target" ];
            unitConfig.ConditionUser = cfg.session.user;
            serviceConfig = {
              Type = "exec";
              ExecStart = "${pkgs.xdg-dbus-proxy}/bin/xdg-dbus-proxy unix:path=/run/user/${uid}/bus ${cfg.session.socket} --filter ${args}";
              Restart = "always";
              RestartSec = 1;
            };
          };
      };
  };
}

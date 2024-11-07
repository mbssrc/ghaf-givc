{
  self,
  lib,
  inputs,
  ...
}:
let
  tls = true;
  snakeoil = ./snakeoil;
  addrs = {
    netvm = "192.168.101.1";
    audiovm = "192.168.101.2";
    guivm = "192.168.101.3";
    appvm = "192.168.101.100";
    adminvm = "192.168.101.10";
  };
  admin = {
    name = "admin-vm";
    addr = addrs.adminvm;
    port = "9001";
    protocol = "tcp";
  };
  mkTls = name: {
    enable = tls;
    caCertPath = "${snakeoil}/${name}/ca-cert.pem";
    certPath = "${snakeoil}/${name}/${name}-cert.pem";
    keyPath = "${snakeoil}/${name}/${name}-key.pem";
  };
in
{
  perSystem = _: {
    vmTests.tests.dbus = {
      module = {
        nodes = {
          adminvm =
            { pkgs, ... }:
            {
              imports = [ self.nixosModules.admin ];

              networking.interfaces.eth1.ipv4.addresses = lib.mkOverride 0 [
                {
                  address = addrs.adminvm;
                  prefixLength = 24;
                }
              ];
              environment.systemPackages = [ pkgs.grpcurl ];
              givc.admin = {
                enable = true;
                inherit (admin) name;
                inherit (admin) addr;
                inherit (admin) port;
                inherit (admin) protocol;
                tls = mkTls "admin-vm";
                debug = false;
              };
            };

          guivm =
            { pkgs, ... }:
            let
              inherit (import "${inputs.nixpkgs.outPath}/nixos/tests/ssh-keys.nix" pkgs)
                snakeOilPublicKey
                ;
            in
            {
              imports = [
                self.nixosModules.sysvm
              ];

              environment.systemPackages = [
                pkgs.networkmanager
              ];

              # Network
              networking.interfaces.eth1.ipv4.addresses = lib.mkOverride 0 [
                {
                  address = addrs.guivm;
                  prefixLength = 24;
                }
              ];

              # Setup users and keys
              users.groups.users = { };
              users.users = {
                ghaf = {
                  isNormalUser = true;
                  group = "users";
                  uid = 1000;
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                };
              };
              services.getty.autologinUser = "ghaf";

              # Test users to check access controls work correctly

              # Parameters:
              # - name: evil1
              # - isNormalUser: User is a normal user
              # - uid: User ID >= 1000
              # - group: 'users', 'networkmanager'
              users.users = {
                evil1 = {
                  isNormalUser = true;
                  uid = 4269;
                  group = "users";
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                };
              };

              # Parameters:
              # - name: evil2
              # - isSystemUser: User is a system user
              # - uid: User ID < 1000
              # - group: 'root', 'networkmanager'
              users.users = {
                evil2 = {
                  isSystemUser = true;
                  uid = 42;
                  group = "users";
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                };
              };

              givc.sysvm = {
                enable = true;
                inherit admin;
                agent = {
                  addr = addrs.guivm;
                  name = "gui-vm";
                };
                tls = mkTls "gui-vm";
                socketProxy = [
                  {
                    transport = {
                      name = "net-vm";
                      addr = addrs.netvm;
                      port = "9010";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_net.sock";
                  }
                  {
                    transport = {
                      name = "audio-vm";
                      addr = addrs.audiovm;
                      port = "9011";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_snd.sock";
                  }
                  {
                    transport = {
                      name = "chromium-vm";
                      addr = addrs.appvm;
                      port = "9012";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_app.sock";
                  }
                ];
                debug = true;
              };
            };

          netvm =
            { pkgs, ... }:
            let
              inherit (import "${inputs.nixpkgs.outPath}/nixos/tests/ssh-keys.nix" pkgs)
                snakeOilPublicKey
                ;
            in
            {
              imports = [
                self.nixosModules.sysvm
              ];

              # Network
              networking.interfaces.eth1.ipv4.addresses = lib.mkOverride 0 [
                {
                  address = addrs.netvm;
                  prefixLength = 24;
                }
              ];

              # Services
              networking.networkmanager.enable = true;
              services.avahi.enable = true;
              services.upower.enable = true;

              # Setup users and keys
              users.groups.users = { };
              users.users = {
                ghaf = {
                  isNormalUser = true;
                  group = "users";
                  uid = 1000;
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                };
              };

              givc.sysvm = {
                enable = true;
                inherit admin;
                agent = {
                  addr = addrs.netvm;
                  name = "net-vm";
                };
                tls = mkTls "net-vm";
                socketProxy = [
                  {
                    transport = {
                      name = "gui-vm";
                      addr = addrs.guivm;
                      port = "9010";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_net.sock";
                  }
                ];
                debug = true;
              };

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
              };
            };

          audiovm =
            { pkgs, ... }:
            let
              inherit (import "${inputs.nixpkgs.outPath}/nixos/tests/ssh-keys.nix" pkgs)
                snakeOilPublicKey
                ;
            in
            {
              imports = [
                self.nixosModules.sysvm
              ];

              # Network
              networking.interfaces.eth1.ipv4.addresses = lib.mkOverride 0 [
                {
                  address = addrs.audiovm;
                  prefixLength = 24;
                }
              ];

              # Service
              services.upower.enable = true;

              # Setup users and keys
              users.groups.users = { };
              users.users = {
                ghaf = {
                  isNormalUser = true;
                  group = "users";
                  uid = 1000;
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                };
              };

              givc.sysvm = {
                enable = true;
                inherit admin;
                agent = {
                  addr = addrs.audiovm;
                  name = "audio-vm";
                };
                tls = mkTls "audio-vm";
                socketProxy = [
                  {
                    transport = {
                      name = "gui-vm";
                      addr = addrs.guivm;
                      port = "9011";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_snd.sock";
                  }
                ];
                debug = true;
              };

              givc.dbusproxy = {
                enable = true;
                system = {
                  enable = true;
                  user = "ghaf";
                  socket = "/tmp/.dbusproxy_snd.sock";
                  policy.talk = [
                    "org.freedesktop.UPower.*"
                  ];
                };
              };
            };

          appvm =
            { pkgs, ... }:
            let
              inherit (import "${inputs.nixpkgs.outPath}/nixos/tests/ssh-keys.nix" pkgs)
                snakeOilPublicKey
                ;
            in
            {
              imports = [
                self.nixosModules.appvm
              ];

              # Network
              networking.interfaces.eth1.ipv4.addresses = lib.mkOverride 0 [
                {
                  address = addrs.appvm;
                  prefixLength = 24;
                }
              ];

              # Service
              services.playerctld.enable = true;

              # Setup users and keys
              users.groups.users = { };
              users.users = {
                ghaf = {
                  isNormalUser = true;
                  group = "users";
                  uid = 1000;
                  openssh.authorizedKeys.keys = [ snakeOilPublicKey ];
                  linger = true;
                };
              };
              services.getty.autologinUser = "ghaf";

              givc.appvm = {
                enable = true;
                inherit admin;
                agent = {
                  addr = addrs.appvm;
                  name = "chromium-vm";
                };
                tls = mkTls "chromium-vm";
                applications = [
                  {
                    name = "test";
                    command = "/bin/bash";
                    args = [ ];
                  }
                ];
                socketProxy = [
                  {
                    transport = {
                      name = "gui-vm";
                      addr = addrs.guivm;
                      port = "9012";
                      protocol = "tcp";
                    };
                    socket = "/tmp/.dbusproxy_app.sock";
                  }
                ];
                debug = true;
              };

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
            };

        };

        testScript = _: ''
          # import time

          with subtest("boot_completed"):
              adminvm.wait_for_unit("multi-user.target")
              audiovm.wait_for_unit("multi-user.target")
              netvm.wait_for_unit("multi-user.target")
              appvm.wait_for_unit("multi-user.target")
              guivm.wait_for_unit("multi-user.target")

          with subtest("success_tests_systembus"):

              # SUCCESS: remote access to netvms NetworkManager service; dbus-send
              print(guivm.succeed("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.NetworkManager /org/freedesktop/NetworkManager org.freedesktop.DBus.Properties.Get string:'org.freedesktop.NetworkManager' string:'ActiveConnections'"))

              # SUCCESS: remote access to netvms NetworkManager service; nmcli
              print(guivm.succeed("sudo -u ghaf -- bash -c 'export DBUS_SYSTEM_BUS_ADDRESS=unix:path=/tmp/.dbusproxy_net.sock; nmcli d'"))

              # SUCCESS: access to additional specified netvm service
              print(guivm.succeed("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.Avahi /org/freedesktop/Avahi org.freedesktop.DBus.Introspectable.Introspect"))

              # SUCCESS: 'call' method access to specified netvm service
              print(guivm.succeed("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.UPower /org/freedesktop/UPower org.freedesktop.UPower.EnumerateDevices"))

              # SUCCESS: connection to secondary system vm (audio)
              print(guivm.succeed("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_snd.sock --print-reply --dest=org.freedesktop.UPower /org/freedesktop/UPower org.freedesktop.DBus.Introspectable.Introspect"))

          with subtest("failure_tests_systembus"):

              # FAIL: 'call' access to non-specified netvm service
              print(guivm.fail("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.UPower /org/freedesktop/UPower org.freedesktop.UPower.GetCriticalAction"))

              # FAIL: root user access to netvm service
              print(guivm.fail("dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.UPower /org/freedesktop/UPower org.freedesktop.DBus.Introspectable.Introspect"))

              # FAIL: evil1 user access to netvm service
              print(guivm.fail("sudo -u evil1 dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.UPower /org/freedesktop/UPower org.freedesktop.UPower.EnumerateDevices"))

              # FAIL: evil2 user access to netvm service
              print(guivm.fail("sudo -u evil2 dbus-send --bus=unix:path=/tmp/.dbusproxy_net.sock --print-reply --dest=org.freedesktop.Avahi /org/freedesktop/Avahi org.freedesktop.DBus.Introspectable.Introspect"))

          with subtest("remote_user_to_sesssionbus_access"):

            # SUCCESS: ghaf user access to audiovm session bus
            print(guivm.succeed("sudo -u ghaf dbus-send --bus=unix:path=/tmp/.dbusproxy_app.sock --print-reply --dest=org.mpris.MediaPlayer2.playerctld /org/mpris/MediaPlayer2 org.freedesktop.DBus.Introspectable.Introspect"))

            # FAIL: root user access to audiovm session bus
            print(guivm.fail("dbus-send --bus=unix:path=/tmp/.dbusproxy_app.sock --print-reply --dest=org.mpris.MediaPlayer2.playerctld /org/freedesktop/MediaPlayer2 org.freedesktop.DBus.Introspectable.Introspect"))

        '';
      };
    };
  };
}

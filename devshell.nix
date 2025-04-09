{ inputs, ... }:
{
  imports = [
    inputs.devshell.flakeModule
    inputs.pre-commit-hooks-nix.flakeModule
  ];

  perSystem =
    {
      self',
      pkgs,
      config,
      ...
    }:
    {
      devshells.default = {
        devshell = {
          name = "GIVC";
          motd = ''
            {14}{bold}❄️ Welcome to the givc devshell ❄️{reset}
            $(type -p menu &>/dev/null && menu)
            $(type -p update-pre-commit-hooks &>/dev/null && update-pre-commit-hooks)
          '';
        };
        packages = [
          config.treefmt.build.wrapper
          pkgs.reuse
          pkgs.gopls
          pkgs.gosec
          pkgs.gotests
          pkgs.go-tools
          pkgs.golangci-lint
          pkgs.rustfmt
          pkgs.clippy
          pkgs.stdenv.cc # Need for build rust components
          pkgs.protobuf
          pkgs.protoc-gen-go
          pkgs.protoc-gen-go-grpc
          pkgs.grpcurl
          # Documentation
          (pkgs.python312.withPackages (
            ps: with ps; [
              pkgs.python312Packages.mkdocs
              pkgs.python312Packages.mkdocs-material
              pkgs.python312Packages.pygments
              pkgs.python312Packages.pymdown-extensions
            ]
          ))
          pkgs.protoc-gen-doc
        ];
        packagesFrom = builtins.attrValues self'.packages;
        commands = [
          {
            name = "update-pre-commit-hooks";
            command = config.pre-commit.installationScript;
            category = "tools";
            help = "update git pre-commit hooks";
          }
          {
            help = "Generate go files from protobuffers. Examples: '$ protogen systemd'";
            name = "go-protogen";
            command = "./modules/api/protoc.sh $@";
          }
          {
            help = "Check golang vulnerabilities";
            name = "go-checksec";
            command = "gosec -exclude=G302,G204 -no-fail ./modules/...";
          }
          {
            help = "Run go tests";
            name = "go-tests";
            command = "go test -v ./modules/...";
          }
          {
            help = "Update go dependencies";
            name = "go-update";
            command = "go get -u ./... && go mod tidy && echo Done - do not forget to update the vendorHash in the packages.";
          }
          {
            help = "golang linter";
            package = "golangci-lint";
            category = "linters";
          }
          {
            help = "Run local docs server";
            name = "docs-server";
            command = "cd docs && mkdocs serve";
          }
        ];
      };
      pre-commit.settings = {
        hooks.treefmt.enable = true;
        hooks.treefmt.package = config.treefmt.build.wrapper;
      };
    };
}

# Copyright 2025 TII (SSRC) and the Ghaf contributors
# SPDX-License-Identifier: Apache-2.0
{
  pkgs,
  lib,
  self,
  src,
  ...
}:
let
  inherit (pkgs)
    stdenv
    nixosOptionsDoc
    ;
  inherit (lib)
    concatMapStringsSep
    evalModules
    filterAttrsRecursive
    ;

  mkOptionsDoc =
    name:
    nixosOptionsDoc {
      inherit pkgs lib;
      options = filterAttrsRecursive (n: _v: n != "_module") (evalModules {
        modules = [
          { _module.check = false; }
          (import (./. + "/../modules/${name}.nix") { inherit self; })
        ];
      });
    };

  opt_docs =
    map
      (doc: {
        name = doc;
        options = (mkOptionsDoc doc).optionsCommonMark;
      })
      [
        "admin"
        "appvm"
        "dbus"
        "host"
        "tls"
        "sysvm"
        # "update-server"
      ];

  mkHeader = file: ''
    echo "
    <!--
      Copyright 2025 TII (SSRC) and the Ghaf contributors
      SPDX-License-Identifier: CC-BY-SA-4.0
    -->
    " > ${file}
  '';

  mkGoGrpcDocs =
    map
      (file: ''
        ${mkHeader "$out/go/grpc_${file}.md"}
        gomarkdoc --output tmp.md $src/modules/api/${file}
        cat tmp.md >> $out/go/grpc_${file}.md
        rm tmp.md
      '')
      [
        "systemd"
        "admin"
        "locale"
        "socket"
        "stats"
      ];

  mkGoPkgsDocs =
    map
      (file: ''
        ${mkHeader "$out/go/pkgs_${file}.md"}
        gomarkdoc --output tmp.md $src/modules/pkgs/${file}
        cat tmp.md >> $out/go/pkgs_${file}.md
        rm tmp.md
      '')
      [
        "grpc"
        "servicemanager"
        "serviceclient"
        "applications"
        "localelistener"
        "statsmanager"
        "types"
        "utility"
      ];

  mkGoCmdDocs = ''
    ${mkHeader "$out/go/givc_agent.md"}
    gomarkdoc --output tmp.md $src/modules/cmd/...
    cat tmp.md >> $out/go/givc_agent.md
    rm tmp.md
  '';
in
stdenv.mkDerivation {
  inherit src;
  name = "docs";

  nativeBuildInputs = [
    pkgs.protobuf
    pkgs.protoc-gen-doc
    pkgs.gomarkdoc
  ];

  dontConfigure = true;
  doCheck = false;

  buildPhase = ''
    mkdir -p $out
    mkdir -p $out/go
    mkdir -p $out/grpc
    mkdir -p $out/nixos

    # Make nixosModules options docs
    ${concatMapStringsSep "\n" (opt_doc: ''
      ${mkHeader "$out/nixos/${opt_doc.name}_options.md"}
      cat ${opt_doc.options} >> $out/nixos/${opt_doc.name}_options.md
    '') opt_docs}

    # Generate protobuf documentation
    cd api
    ${mkHeader "$out/grpc/api.md"}
    protoc --doc_out=$out/grpc --doc_opt=$src/docs/templates/grpc2.tmpl,tmp.md */*.proto
    cat $out/grpc/tmp.md >> $out/grpc/api.md
    rm $out/grpc/tmp.md
    cd ..

    # Generate go documentation
    ${mkGoCmdDocs}
    ${concatMapStringsSep "\n" (go_doc: ''${go_doc}'') mkGoGrpcDocs}
    ${concatMapStringsSep "\n" (go_doc: ''${go_doc}'') mkGoPkgsDocs}
  '';

  installPhase = ''
    for file in $out/go/*; do chmod 0666 $file; done
    for file in $out/grpc/*; do chmod 0666 $file; done
    for file in $out/nixos/*; do chmod 0666 $file; done
  '';
}

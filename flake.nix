{
  description = "go env flake";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    golines = {
      url = "github:segmentio/golines";
      flake = false;
    };
    gotest-tools = {
      url = "github:gotestyourself/gotest.tools";
      flake = false;
    };
    godlv = {
      url = "github:go-delve/delve";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, golines, godlv,gotest-tools }@inputs:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
      golines = pkgs.buildGoModule rec {
        pname = "golines";
        src = inputs.golines;
        vendorSha256 = "sha256-It7lD8Ix9oX8xXILCnWUfit9ZlPJ4zjMElNa14mCkGI=";
        name = "golines";
        proxyVendor = true;
      };
      godlv = pkgs.buildGoModule rec {
        pname = "dlv";
        src = inputs.godlv;
        vendorSha256 = null;
        name = "dlv";
        proxyVendor = true;
        doCheck = false;
      };

      gotest= pkgs.buildGoModule rec {
        pname = "gotest.tools";
        src = inputs.godlv;
        vendorSha256 = null;
        name = "gotest-tools";
        proxyVendor = true;
        doCheck = false;
      };
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        nativeBuildInputs = [
          pkgs.bashInteractive
          pkgs.go
          pkgs.golangci-lint
          pkgs.gofumpt
          pkgs.iferr
          pkgs.impl
          pkgs.gotestsum
          pkgs.gomodifytags
          pkgs.mockgen
          pkgs.richgo
          pkgs.gotools
          pkgs.gotests
          pkgs.reftools
          golines
          godlv
          gotest
        ];

        buildInputs = [ ];
        PATH = "$PATH:$(go env GOPATH)/bin";
      };
    };
}

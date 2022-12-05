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
    gods = {
      url = "github:emirpasic/gods";
      flake = false;
    };
    hashset = {
      url = "github:emirpasic/gods?dir=sets/hashset";
      flake = false;
    };

    stack = {
      url = "github:emirpasic/gods?dir=stacks/arraystacks";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, golines, godlv,gotest-tools,gods,hashset,stack }@inputs:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
      golines = pkgs.buildGoModule  {
        pname = "golines";
        src = inputs.golines;
        vendorSha256 = "sha256-It7lD8Ix9oX8xXILCnWUfit9ZlPJ4zjMElNa14mCkGI=";
        name = "golines";
        proxyVendor = true;
      };
      godlv = pkgs.buildGoModule  {
        pname = "dlv";
        src = inputs.godlv;
        vendorSha256 = null;
        name = "dlv";
        proxyVendor = true;
        doCheck = false;
      };

      stack = pkgs.buildGoModule  {
        pname = "stack";
        src = inputs.stack;
        vendorSha256 = null;
        name = "stack";
        proxyVendor = true;
        doCheck = false;
      };

      hashset = pkgs.buildGoModule  {
        pname = "hashset";
        src = inputs.hashset;
        vendorSha256 = null;
        name = "hashset";
        proxyVendor = true;
        doCheck = false;
      };
      gods = pkgs.buildGoModule  {
        pname = "gods";
        src = inputs.gods;
        vendorSha256 = null;
        name = "gods";
        proxyVendor = true;
        doCheck = false;
      };

      gotest= pkgs.buildGoModule {
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
          godlv
          gods
          golines
          gotest
          hashset
          stack
          pkgs.bashInteractive
          pkgs.ginkgo
          pkgs.go
          pkgs.gofumpt
          pkgs.golangci-lint
          pkgs.gomodifytags
          pkgs.gotests
          pkgs.gotestsum
          pkgs.gotools
          pkgs.iferr
          pkgs.impl
          pkgs.mockgen
          pkgs.reftools
          pkgs.richgo
        ];

        buildInputs = [ ];
        PATH = "$PATH:$(go env GOPATH)/bin";
      };
    };
}

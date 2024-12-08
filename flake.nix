{
  description = "nix-flake-sync is a go tool intended to synchronize inputs accross multiple flakes";
  inputs = {
    flake-utils = {
      url = "github:numtide/flake-utils";
    };
    nixpkgs = {
      url = "nixpkgs/nixos-24.11";
    };
    nixpkgs-unstable = {
      url = "nixpkgs/nixos-unstable";
    };
    gomod2nix = {
      # url = "github:nix-community/gomod2nix";
      url = "github:ghthor/gomod2nix?ref=fix/go_mod_vendor_go_1_23"; # support for go1.23
      inputs.flake-utils.follows = "flake-utils";
      inputs.nixpkgs.follows = "nixpkgs-unstable";
    };
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
      nixpkgs-unstable,
      gomod2nix,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        pkgs-unstable = import nixpkgs-unstable { inherit system; };

        go_1_23 = pkgs.go;

        inherit (gomod2nix.legacyPackages.${system}) buildGoApplication;
      in
      {
        formatter = pkgs.nixfmt-rfc-style;
        apps =
          let
            selfPkgs = self.packages.${system};
          in
          rec {
            default = nix-flake-sync;
            gomod2nix = {
              type = "app";
              program = "${selfPkgs.gomod2nix}/bin/gomod2nix";
            };
            golangci-lint = {
              type = "app";
              program = "${pkgs.golangci-lint}/bin/golangci-lint";
            };
            nix-flake-sync = {
              type = "app";
              program = "${selfPkgs.nix-flake-sync}/bin/nix-flake-sync";
            };
          };
        packages.default = self.packages.${system}.nix-flake-sync;
        packages = {
          gomod2nix = gomod2nix.packages.${system}.default;
          nix-flake-sync = pkgs.callPackage buildGoApplication {
            pname = "nix-flake-sync";
            version = "0.0.1";
            pkg = ./.;
            src = ./.;
            go = go_1_23;
            modules = ./gomod2nix.toml;
            # subPackages = [ "." ];

            ldflags = [ ];
            enableParallelBuilding = true;

            checkPhase = ''
              runHook preCheck
              # for pkg in $(getGoDirs test); do
              #   buildGoDir test $checkFlags "$pkg"
              # done
              runHook postCheck
            '';

            nativeBuildInputs = [ pkgs.installShellFiles ];
            postInstall = ''
              installShellCompletion --cmd nix-flake-sync \
                --bash <($out/bin/nix-flake-sync completion bash) \
                --fish <($out/bin/nix-flake-sync completion fish) \
                --zsh  <($out/bin/nix-flake-sync completion zsh)
            '';
          };
        };
        devShells.default = pkgs.mkShell {
          nativeBuildInputs = [
            go_1_23
            gomod2nix.packages.${system}.default
            pkgs.golangci-lint
            self.packages.${system}.nix-flake-sync
          ];
        };
      }
    );
}

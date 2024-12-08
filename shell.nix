{
  pkgs,
  pkgs-unstable,
  go,
  gomod2nix,
  nix-flake-sync,
  ...
}:
pkgs.mkShell {
  nativeBuildInputs = [
    go
    gomod2nix
    pkgs.golangci-lint
    pkgs.cobra-cli
    nix-flake-sync
  ];

  shellHook = ''
     echo
    command -v go
    command -v golangci-lint
    command -v cobra-cli
    command -v gomod2nix
    command -v nix-flake-sync
    echo
  '';
}

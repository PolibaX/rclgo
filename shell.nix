# Nix shells are used to create a confined development environment, with
# everything needed to develop or build the software.
#
# To create the shell, run `nix-shell ./shell.nix`.
#
# This command will place you in a development shell with all the
# following packages at your disposal.
#
# More info:
#  - What is NixOS: https://nixos.org/
#  - What is nix-shell: https://nixos.org/manual/nix/stable/command-ref/nix-shell.html
#

{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs.buildPackages; [
    gnumake
    go
    gopls
  ];

  shellHook = ''
    export GO111MODULE=on
    export GOPRIVATE=github.com/PolibaX/rclgo/
  '';
}

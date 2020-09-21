{ pkgs ? import <nixpkgs> {} }:
with pkgs;
let
  extra_ignores = ''
    /.*
    /pystarport
    /go.mod
    /go.sum
    /integration-tests
  '';
in
buildGoPackage rec {
  pname = "chain-main";
  version = "0.0.1";
  goPackagePath = "github.com/crypto-com/chain-main";
  src = nix-gitignore.gitignoreSource extra_ignores ./.;
  # deps.nix is generated by vgo2nix
  # (FIXME use the patched version here: https://github.com/yihuang/vgo2nix/tree/fix-go-module)
  goDeps = ./deps.nix;
}

{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    avrogoSrc = {
      url = "github:heetch/avro/v0.4.4";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-utils, avrogoSrc }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        avrogo = pkgs.buildGoModule {
          pname = "avrogo";
          version = "0.4.4";
          src = avrogoSrc;
          vendorSha256 = "sha256-UCYBR3QzznJTefgSM14lCYtaEyT4ykSUiW1COjfJtkA=";
          subPackages = [ "cmd/avrogo" ];

          # tests do not work without external variables set so don't bother
          checkPhase = "";
        };
      in rec {
        devShells.default =
          pkgs.mkShell { packages = with pkgs; [ go gotools gopls avrogo ]; };
      });
}

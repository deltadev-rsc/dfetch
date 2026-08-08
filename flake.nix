{
    description = "Nix flake for dfetch";
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-26.05";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils }:
        flake-utils.lib.eachDefaultSystem (system: 
            let 
                pkgs = import nixpkgs { inherit system; };
            in 
            {
                packages.default = pkgs.buildGoModule {
                    pname = "dfetch";
                    version = "0.1";
                    src = ./.;
                    vendorSha256 = pkgs.lib.flakeSha256;
                    vendor = false;
                };

                devShells.default = pkgs.mkShell {
                    packages = with pkgs; [
                        go
                        gopls
                        golangci-lint
                        git
                        gnumake
                    ];

                    shellHook = ''
                        echo "Golang dev shell ready!"
                        echo "You can run 'make run' to build dfetch"
                    '';
                };
            }
        );
}

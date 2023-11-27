{
  description = "Game Queue Flake";
  
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-23.05";
    unstable.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    templ.url = "github:a-h/templ";

    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "flake-utils";
    };
  };

  outputs = { self, nixpkgs, flake-utils, gomod2nix, unstable, templ }: 
  flake-utils.lib.eachDefaultSystem 
  (system:
    let 
      pkgs = import nixpkgs {
        inherit system;
        overlays = [ 
          gomod2nix.overlays.default
        ];
      };
      unstablePkgs = import unstable {
        inherit system;
      };
      dev = pkgs.writeShellScriptBin "dev" ''
        npx tailwindcss -i ./input.css -o ./static/output.css
        templ generate 
        go run ./cmd/main.go
      '';
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          unstablePkgs.go_1_21
          unstablePkgs.gopls
          gotools
          go-tools
          templ.packages.${system}.default
          gomod2nix.packages.${system}.default
          nodejs
          nodePackages.npm
          dev
        ];
      };
    }
  );
}

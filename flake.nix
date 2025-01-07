{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = inputs@{ self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [ go golangci-lint fish envsubst bashInteractive ];
        shellHook = ''
          echo "Go devShell"

          pushd $(git rev-parse --show-toplevel) >/dev/null
          complete -W "$(ls -1 ./environments/*.env | xargs -n1 basename | xargs -d '\n' -I {} printf ' {}')" runWithEnv || true
          popd >/dev/null
        '';
      };

    };
}

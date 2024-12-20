# This flake was initially generated by fh, the CLI for FlakeHub (version 0.1.10)
{
  # A helpful description of your flake
  description = "Smyl DevOps";

  # Flake inputs
  inputs = {
    flake-schemas.url = "https://flakehub.com/f/DeterminateSystems/flake-schemas/*.tar.gz";

    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/0.1.*.tar.gz";
  };

  # Flake outputs that other flakes can use
  outputs =
    {
      self,
      flake-schemas,
      nixpkgs,
    }:
    let
      # Helpers for producing system-specific outputs
      supportedSystems = [
        "x86_64-linux"
        "aarch64-darwin"
        "x86_64-darwin"
        "aarch64-linux"
      ];
      forEachSupportedSystem =
        f: nixpkgs.lib.genAttrs supportedSystems (system: f { pkgs = import nixpkgs { inherit system; }; });
    in
    {
      # Schemas tell Nix about the structure of your flake's outputs
      schemas = flake-schemas.schemas;

      # Development environments
      devShells = forEachSupportedSystem (
        { pkgs }:
        {
          default = pkgs.mkShell {
            # Pinned packages available in the environment
            packages = with pkgs; [
              awscli2
              go
              gopls
              pulumi-bin
              ssm-session-manager-plugin
            ];

            # A hook that runs every time you enter the environment
            shellHook = ''
              echo -e "\033[0;31m### Welcome to Pulumi Zephyr. Remember to commit early & often! ###\033[0m"
              echo -e "\033[0;33mThis shell provides the following development tools:\033[0m"
              echo "* $(aws --version)"
              echo "* AWS session-manager-plugin version: $(session-manager-plugin --version)"
              echo "* $(go version)"
              echo "* $(gopls version)"
              echo "* pulumi version: $(pulumi version)"
              echo -e "\033[0;32m------End Tools------\033[0m"
            '';
          };
        }
      );
    };
}

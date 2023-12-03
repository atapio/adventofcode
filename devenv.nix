{ pkgs, ... }:

{
  env.GREET = "Advent of Code devenv";

  packages = [
  ];

  languages.rust = {
    enable = true;
  };

  dotenv.enable = true;

  enterShell = ''
    echo $GREET
  '';
}



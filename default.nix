{pkgs ? import <nixpkgs> {}}:
pkgs.buildGoModule rec {
  name = "send2kindle";
  version = "0.0.1";
  vendorSha256 = "sha256-cptWnGcOPcL4aFsYL/zdf9b5tlySLtlqjA4SVqKL5D4=";
  src = ./.;
  buildInputs = with pkgs; [
    calibre-py2
  ];
  meta = with pkgs.lib; {
    description = "send2kindle: basically send files over email to your kindle with the option to convert to mobi before using calibre";
    homepage = "https://github.com/lucasew/send2kindle";
    platforms = platforms.linux;
  };
}

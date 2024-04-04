{ buildGoModule
, calibre
, lib
}:

buildGoModule {
  name = "send2kindle";
  version = "0.0.1";
  vendorHash = "sha256-cptWnGcOPcL4aFsYL/zdf9b5tlySLtlqjA4SVqKL5D4=";
  src = ./.;
  buildInputs = [
    calibre
  ];
  meta = with lib; {
    description = "send2kindle: basically send files over email to your kindle with the option to convert to mobi before using calibre";
    homepage = "https://github.com/lucasew/send2kindle";
    platforms = platforms.linux;
  };
}

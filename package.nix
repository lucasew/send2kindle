{ buildGoModule
, calibre
, lib
}:

buildGoModule {
  name = "send2kindle";
  version = "0.0.1";
  vendorHash = "sha256-TKzqGLYjkiq2EYNP4bMXWx5JpagMXCJkSVEgE0BmLDw=";
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

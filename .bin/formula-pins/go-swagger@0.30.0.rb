class GoSwaggerAT0300 < Formula
  desc "Toolkit to work with swagger for golang"
  homepage "https://github.com/go-swagger/go-swagger"
  version "0.30.0"
  @@filename = nil
  if OS.mac?
    if Hardware::CPU.arm?
      @@filename = "swagger_darwin_arm64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "757ce974b8da0eb024c948d521191d9f80db5007ec47c9a95a89f3d7f39a922b"
    else
      @@filename = "swagger_darwin_amd64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "bed597232c8a82d2fc7e341774890f0b000e27d9c657eca16da098da823c9ab0"
    end
  elsif OS.linux?
    case RbConfig::CONFIG["host_cpu"]
    when "aarch64"
      @@filename = "swagger_linux_arm64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "12dd702a75ed4ca47cabf8c0dc03a5b58dc28f89091afdc1f82ea90b58fda134"
    when "x86_64"
      @@filename = "swagger_linux_amd64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "1ed5bf204c45e9f8614c7d65b6bee5cf10087db267d5f50a6185302cb8484bd6"
    else
      ohdie "Architecture not supported by this formula"
    end
  end

  option "with-goswagger", "Names the binary goswagger instead of swagger"

  def install
    nm = "swagger"
    if build.with? "goswagger"
      nm = "goswagger"
    end
    system "mv", @@filename, nm
    bin.install nm
  end

  test do
    if build.with? "goswagger"
      system "#{bin}/goswagger", "version"
    else
      system "#{bin}/swagger", "version"
    end
  end
end

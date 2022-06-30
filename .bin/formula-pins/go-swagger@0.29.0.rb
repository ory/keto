class GoSwaggerAT0290 < Formula
  desc "Toolkit to work with swagger for golang"
  homepage "https://github.com/go-swagger/go-swagger"
  version "0.29.0"
  @@filename = nil
  if OS.mac?
    if Hardware::CPU.arm?
      @@filename = "swagger_darwin_arm64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "babf863e2a54c6e173cf7104db63269d5c01105f6fd00b04c2ace0409278a7de"
    else
      @@filename = "swagger_darwin_amd64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "90438d5fc13cc0586d64187da1879ae5e01f0de23718225c2fc1fbee1a1be59f"
    end
  elsif OS.linux?
    case RbConfig::CONFIG["host_cpu"]
    when "aarch64"
      @@filename = "swagger_linux_arm64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "eeb8f3444c4f3622207e8dca31188bb889e7dda10f0c5ad76779f13f24e945e7"
    when "x86_64"
      @@filename = "swagger_linux_amd64"
      url "https://github.com/go-swagger/go-swagger/releases/download/v#{version}/#{@@filename}"
      sha256 "0666361b45e11862e3d6487693da9f498710e395660ed0fcbea835a9e8c7272d"
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

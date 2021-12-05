PB_REL="https://github.com/protocolbuffers/protobuf/releases"
PB_VERION="3.13.0"
curl -LO $PB_REL/download/v${PB_VERION}/protoc-${PB_VERION}-linux-x86_64.zip
unzip protoc-${PB_VERION}-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"

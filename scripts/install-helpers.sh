#!/bin/env bash
set -euo pipefail

is_command() {
	command -v "$1" >/dev/null
}
uname_os() {
	os=$(uname -s | tr '[:upper:]' '[:lower:]')
	case "$os" in
	cygwin_nt*) os="windows" ;;
	mingw*) os="windows" ;;
	msys_nt*) os="windows" ;;
  darwin) os="osx" ;;
	esac
	echo "$os"
}
uname_arch() {
	arch=$(uname -m)
	case $arch in
	x86_64) arch="amd64" ;;
	x86) arch="386" ;;
	i686) arch="386" ;;
	i386) arch="386" ;;
	aarch64) arch="arm64" ;;
	armv5*) arch="armv5" ;;
	armv6*) arch="armv6" ;;
	armv7*) arch="armv7" ;;
	esac
	echo ${arch}
}
untar() {
	tarball=$1
	case "${tarball}" in
	*.tar.gz | *.tgz) tar --no-same-owner -xzf "${tarball}" ;;
	*.tar) tar --no-same-owner -xf "${tarball}" ;;
	*.zip) unzip "${tarball}" ;;
	*)
		echo "untar unknown archive format for ${tarball}"
		return 1
		;;
	esac
}
http_download_curl() {
	local_file=$1
	source_url=$2
	header=$3
	if [ -z "$header" ]; then
		code=$(curl -w '%{http_code}' -sL -o "$local_file" "$source_url")
	else
		code=$(curl -w '%{http_code}' -sL -H "$header" -o "$local_file" "$source_url")
	fi
	if [ "$code" != "200" ]; then
		echo "http_download_curl received HTTP status $code"
		return 1
	fi
	return 0
}
http_download_wget() {
	local_file=$1
	source_url=$2
	header=$3
	if [ -z "$header" ]; then
		wget -q -O "$local_file" "$source_url"
	else
		wget -q --header "$header" -O "$local_file" "$source_url"
	fi
}
http_download() {
	echo "downloading $2"
	if is_command curl; then
		http_download_curl "$@"
		return
	elif is_command wget; then
		http_download_wget "$@"
		return
	fi
	echo "http_download unable to find wget or curl"
	return 1
}
update_binary_version() {
  bindir=$1
  binname=$2
  version=$3
  echo "updating binary version to ${version}"
  echo "${version}" > "${bindir}/.${binname}.version"
}
check_binary_version() {
  bindir=$1
  binname=$2
  version=$3
  echo "checking binary version"
  [ "$(cat "${bindir}/.${binname}.version")" = "${version}" ]
}

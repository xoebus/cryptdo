#!/bin/bash

set -e
set -u

check_installed() {
  if ! command -v "$1" > /dev/null; then
    echo "$1 must be installed and in your PATH!" >&2
    exit 1
  fi
}

downcase() {
  tr '[:upper:]' '[:lower:]'
}

sha256() {
  shasum -a 256 "$@" | awk '{ print "SHA256", "(" $2 ")", "=", $1 }'
}

verify_dependencies() {
  check_installed git
  check_installed bazel
  check_installed bats
  check_installed signify
}

distro() {
  uname | downcase
}

run_tests() {
  ROOT="$(git rev-parse --show-toplevel)"
  "$ROOT"/tools/test.sh
}

build_release() {
  bazel build //...

  BAZEL_BIN=$(bazel info bazel-bin)
  RELEASE_DIR=$(mktemp -d)

  cp "$BAZEL_BIN/cryptdo.tgz" "$RELEASE_DIR/cryptdo-$(distro)-$VERSION.tgz"

  pushd "$RELEASE_DIR" >/dev/null
    sha256 "cryptdo-$(distro)-$VERSION.tgz" > SHA256
    signify -S -e -s "$SIGNIFY_SECKEY" -m SHA256 -x SHA256.sig
    rm SHA256
    signify -C -p "$SIGNIFY_PUBKEY" -x SHA256.sig
  popd >/dev/null

  git tag -a "$VERSION" -m "version $VERSION"
}

show_instructions() {
  echo "If everything looks good you need to:"
  echo
  echo "1. Push the tag ($VERSION)."
  echo "2. Upload the files in $RELEASE_DIR to"
  echo "   the associated release on GitHub."
  echo "3. Compile this on a different platform."
}

main() {
  verify_dependencies
  run_tests
  build_release
  show_instructions
}

VERSION=$1
SIGNIFY_PUBKEY=$2
SIGNIFY_SECKEY=$3

main


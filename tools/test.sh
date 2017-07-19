#!/bin/bash

set -e
set -u

check_installed() {
  if ! command -v "$1" > /dev/null; then
    echo "$1 must be installed and in your PATH!" >&2
    exit 1
  fi
}

verify_dependencies() {
  check_installed git
  check_installed bazel
  check_installed bats
}

run_tests() {
  bazel test //...

  bazel build //cmd/...
  BAZEL_BIN=$(bazel info bazel-bin)
  CRYPTDO_PATH="$BAZEL_BIN/cmd/cryptdo"
  CRYPTDO_PATH="$BAZEL_BIN/cmd/cryptdo-rekey:$CRYPTDO_PATH"
  CRYPTDO_PATH="$BAZEL_BIN/cmd/cryptdo-bootstrap:$CRYPTDO_PATH"

  ROOT="$(git rev-parse --show-toplevel)"
  env PATH="$CRYPTDO_PATH:$PATH" bats "$ROOT"/test/*.bats
}

verify_dependencies
run_tests

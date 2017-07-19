#!/bin/bash

set -e
set -u

bazel build //cmd/...
BAZEL_BIN=$(bazel info bazel-bin)

for bin in cryptdo cryptdo-rekey cryptdo-bootstrap; do
  cp "$BAZEL_BIN/cmd/$bin/$bin" "$HOME/bin/$bin"
done


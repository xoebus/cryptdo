#!/bin/bash

set -e
set -u

dep ensure

# regenerate pruned BUILD files
bazel run //:gazelle

# make sure everything still works
bazel test //...

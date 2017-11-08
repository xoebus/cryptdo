BATS_REPOSITORY_BUILD_FILE = """
package(default_visibility = [ "//visibility:public" ])

exports_files([
  "libexec/bats",
])

filegroup(
  name = "bats",
  srcs = ["libexec/bats"],
)

# filegroup(
#   name = "bats_tools",
#   srcs = [
#     "libexec/bats-exec-suite",
#     "libexec/bats-exec-test",
#     "libexec/bats-format-tap-stream",
#     "libexec/bats-preprocess",
#   ],
# )
"""

def bats_repositories(version="v0.4.0"):
    native.new_git_repository(
      name = "bats",
      remote = "https://github.com/sstephenson/bats",
      tag = version,
      build_file_content = BATS_REPOSITORY_BUILD_FILE
    )

BASH_TEMPLATE = """
#!/usr/bin/env bash
set -e

BATS_BINS_PATH="{bats_bins_path}"

for dir in ${{BATS_BINS_PATH//:/ }}; do
  export PATH="$(dirname $PWD/$dir):$PATH"
done

export TMPDIR="$TEST_TMPDIR"

"{bats}" "{test_paths}"
"""

def _bats_test_impl(ctx):
  bats = ctx.file._bats

  test_files = []
  test_paths = []
  for src in ctx.attr.srcs:
    for file in src.files:
      test_files.append(file)
      test_paths.append(file.short_path)

  bin_files = []
  bin_dirs = []
  for bin in ctx.attr.bins:
    for bin_file in bin.files:
      bin_files.append(bin_file)
      bin_dirs.append(bin_file.short_path)

  ctx.file_action(
      output = ctx.outputs.executable,
      executable = True,
      content = BASH_TEMPLATE.format(
          bats = bats.path,
          test_paths = " ".join(test_paths),
          bats_bins_path = ":".join(bin_dirs),
      ),
  )

  runfiles = [bats] + test_files + bin_files

  return struct(
      runfiles = ctx.runfiles(
          files = runfiles,
          collect_data = True,
      ),
  )

bats_test = rule(
    attrs = {
        "srcs": attr.label_list(
            allow_files = True,
        ),
        "bins": attr.label_list(),
        "_bats": attr.label(
            default = Label("@bats//:bats"),
            single_file = True,
            allow_files = True,
            executable = True,
            cfg = "host",
        ),
    },
    test = True,
    implementation = _bats_test_impl,
)

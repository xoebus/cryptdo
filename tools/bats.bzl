BATS_REPOSITORY_BUILD_FILE = """
package(default_visibility = [ "//visibility:public" ])

sh_binary(
  name = "bats",
  srcs = ["libexec/bats"],
  data = [
    "libexec/bats-exec-suite",
    "libexec/bats-exec-test",
    "libexec/bats-format-tap-stream",
    "libexec/bats-preprocess",
  ],
)
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

export TMPDIR="$TEST_TMPDIR"
export PATH="{bats_bins_path}":$PATH

"{bats}" "{test_paths}"
"""

def _dirname(path):
  prefix, _, _ = path.rpartition("/")
  return prefix.rstrip("/")

def _bats_test_impl(ctx):
  runfiles = ctx.runfiles(
      files = ctx.files.srcs,
      collect_data = True,
  )

  tests = [f.short_path for f in ctx.files.srcs]
  path = ["$PWD/" + _dirname(b.short_path) for b in ctx.files.deps]

  sep = ctx.configuration.host_path_separator

  ctx.file_action(
      output = ctx.outputs.executable,
      executable = True,
      content = BASH_TEMPLATE.format(
          bats = ctx.executable._bats.short_path,
          test_paths = " ".join(tests),
          bats_bins_path = sep.join(path),
      ),
  )

  runfiles = runfiles.merge(ctx.attr._bats.default_runfiles)

  return DefaultInfo(
      runfiles = runfiles,
  )

bats_test = rule(
    attrs = {
        "srcs": attr.label_list(
            allow_files = True,
        ),
        "deps": attr.label_list(),
        "_bats": attr.label(
            default = Label("@bats//:bats"),
            executable = True,
            cfg = "host",
        ),
    },
    test = True,
    implementation = _bats_test_impl,
)

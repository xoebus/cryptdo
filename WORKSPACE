workspace(name = "cryptdo")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.1/rules_go-0.10.1.tar.gz",
    sha256 = "4b14d8dd31c6dbaf3ff871adcd03f28c3274e42abc855cb8fb4d01233c0154dc",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

load("//tools:bats.bzl", "bats_repositories")
bats_repositories()

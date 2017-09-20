workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "0fb90c43c5fab2a0b2d7a8684f26f6995d9aa212",
)

load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_rules_dependencies", "go_register_toolchains", "go_repository")
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
proto_register_toolchains()

load("//:bats.bzl", "bats_repositories")
bats_repositories()

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "6cf8f02b4ae8ba723ddc64dcfd403e530c06d927",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "7d9177d70076375b9a59c8fde23d52d9c4a7ecd5",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "062cd7e4e68206d8bab9b18396626e855c992658",
    importpath = "golang.org/x/sys",
)

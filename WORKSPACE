workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.6.0",
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
    commit = "f88afde2fa19a30cf50ba4b05b3d13bc6bae3079",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "9419663f5a44be8b34ca85f08abc5fe1be11f8a3",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "686000749eaec0b8855b8eef5336cf63899fe51d",
    importpath = "golang.org/x/sys",
)

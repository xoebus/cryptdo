workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")
go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")
go_proto_repositories()

go_repository(
    name = "com_github_golang_protobuf",
    commit = "fec3b39b059c0f88fa6b20f5ed012b1aa203a8b4",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "460c7bb0abd6e927f2767cadc91aa6ef776a98b4",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "0fe963104e9d1877082f8fb38f816fcd97eb1d10",
    importpath = "golang.org/x/crypto",
)


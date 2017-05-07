workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")
go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")
go_proto_repositories()

new_go_repository(
    name = "com_github_golang_protobuf",
    commit = "18c9bb3261723cd5401db4d0c9fbc5c3b6c70fe8",
    importpath = "github.com/golang/protobuf",
)

new_go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "460c7bb0abd6e927f2767cadc91aa6ef776a98b4",
    importpath = "github.com/jessevdk/go-flags",
)

new_go_repository(
    name = "org_golang_x_crypto",
    commit = "5a033cc77e57eca05bdb50522851d29e03569cbe",
    importpath = "golang.org/x/crypto",
)


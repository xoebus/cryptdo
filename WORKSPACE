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
    commit = "2bba0603135d7d7f5cb73b2125beeda19c09f4ef",
    importpath = "github.com/golang/protobuf",
)

new_go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "48cf8722c3375517aba351d1f7577c40663a4407",
    importpath = "github.com/jessevdk/go-flags",
)

new_go_repository(
    name = "org_golang_x_crypto",
    commit = "96846453c37f0876340a66a47f3f75b1f3a6cd2d",
    importpath = "golang.org/x/crypto",
)


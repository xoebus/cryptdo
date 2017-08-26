workspace(name = "cryptdo")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.3",
)

load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_repositories", "go_repository")
go_repositories()

load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_repositories")
go_proto_repositories()

go_repository(
    name = "com_github_golang_protobuf",
    commit = "0a4f71a498b7c4812f64969510bcb4eca251e33a",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "5695738f733662da3e9afc2283bba6f3c879002d",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "dd85ac7e6a88fc6ca420478e934de5f1a42dd3c6",
    importpath = "golang.org/x/crypto",
)


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
    commit = "e325f446bebc2998605911c0a2650d9920361d4a",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    commit = "5695738f733662da3e9afc2283bba6f3c879002d",
    importpath = "github.com/jessevdk/go-flags",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "850760c427c516be930bc91280636328f1a62286",
    importpath = "golang.org/x/crypto",
)


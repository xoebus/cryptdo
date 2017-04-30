load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix", "go_test")

go_prefix("github.com/xoebus/cryptdo")

go_library(
    name = "go_default_library",
    srcs = [
        "crypto.go",
        "fuzz.go",
        "passphrase.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//proto:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_x_crypto//pbkdf2:go_default_library",
        "@org_golang_x_crypto//ssh/terminal:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["crypto_test.go"],
    deps = ["//:go_default_library"],
)

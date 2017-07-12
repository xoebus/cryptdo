load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix", "go_test")

go_prefix("code.xoeb.us/cryptdo")

go_library(
    name = "go_default_library",
    srcs = [
        "crypto.go",
        "passphrase.go",
        "versions.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//cryptdopb:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_x_crypto//pbkdf2:go_default_library",
        "@org_golang_x_crypto//ssh/terminal:go_default_library",
    ],
)

go_test(
    name = "crypto",
    srcs = [
        "crypto_test.go",
        "versions_test.go",
    ],
    deps = [
        "//:go_default_library",
        "//cryptdopb:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
    data = glob(["testdata/**"]),
    size = "small",
)

go_test(
    name = "roundtrip",
    srcs = ["roundtrip_test.go"],
    deps = [
        "//:go_default_library",
    ],
    size = "medium",
)

load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "cryptdo",
    extension = "tgz",
    files = [
        "//cmd/cryptdo:cryptdo",
        "//cmd/cryptdo-rekey:cryptdo-rekey",
        "//cmd/cryptdo-bootstrap:cryptdo-bootstrap",
    ],
    mode = "0755",
)

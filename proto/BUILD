load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_library")

filegroup(
    name = "protos",
    srcs = glob(["*.proto"]),
)

go_proto_library(
    name = "go_default_library",
    srcs = ["cryptdo.proto"],
    visibility = ["//visibility:public"],
)

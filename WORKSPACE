http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz",
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
)
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.9/bazel-gazelle-0.9.tar.gz",
    sha256 = "0103991d994db55b3b5d7b06336f8ae355739635e0c2379dea16b8213ea5a223",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(go_version="host")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

load("@io_bazel_rules_go//go:def.bzl", "go_repository")

go_repository(
    name = "com_github_brianvoe_gofakeit",
    importpath = "github.com/brianvoe/gofakeit",
    commit = "b0b2ecfdf447299dd6bcdef91001692fc349ce4c",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    commit = "c6c44e6fdcc30161c7f4480754da7230d01c06e3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    commit = "aafc9e6bc7b7bb53ddaa75a5ef49a17d6e654be5",
)

go_repository(
    name = "com_github_mitchellh_go-homedir",
    importpath = "github.com/mitchellh/go-homedir",
    commit = "b8bc1bf767474819792c23f32d8286a45736f1c6",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    commit = "00c29f56e2386353d58c599509e8dc3801b0d716",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "github.com/go-yaml/yaml/tree/v2.1.1",
    commit = "7f97868eec74b32b0982dd158a51a446d1da7eb5",
)

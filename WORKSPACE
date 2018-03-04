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
go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    commit = "c2828203cd70a50dcccfb2761f8b1f8ceef9a8e9",
)
go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    commit = "23c074d0eceb2b8a5bfdbb271ab780cde70f05a8",
)
go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    commit = "2c9e9502788518c97fe44e8955cd069417ee89df",
)
go_repository(
    name = "com_github_pelletier_go-toml",
    importpath = "github.com/pelletier/go-toml",
    commit = "05bcc0fb0d3e60da4b8dd5bd7e0ea563eb4ca943",
)
go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    commit = "bbf41cb36dffe15dff5bf7e18c447801e7ffe163",
)
go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    commit = "8965335b8c7107321228e3e3702cab9832751bac",
)
go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    commit = "7c0cea34c8ece3fbeb2b27ab9b59511d360fb394",
)
go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    commit = "ee5fd03fd6acfd43e44aea0b4135958546ed8e73",
)

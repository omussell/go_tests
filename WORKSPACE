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

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(go_version = "host")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

go_repository(
    name = "com_github_brianvoe_gofakeit",
    importpath = "github.com/brianvoe/gofakeit",
    commit = "b0b2ecfdf447299dd6bcdef91001692fc349ce4c",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    commit = "7b2c5ac9fc04fc5efafb60700713d4fa609b777b",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    commit = "25b30aa063fc18e48662b86996252eabdcf2f0c7",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
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
    importpath = "gopkg.in/yaml.v2",
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
    commit = "c3beff4c2358b44d0493c7dda585e7db7ff28ae6",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    commit = "acdc4509485b587f5e675510c4f2c63e90ff68a8",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    commit = "bb8f1927f2a9d3ab41c9340aa034f6b803f4359c",
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
    commit = "e57e3eeb33f795204c1ca35f56c44f83227c6e66",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    commit = "76626ae9c91c4f2a10f34cad8ce83ea42c93bb75",
    importpath = "github.com/inconshreveable/mousetrap",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "dd2ff4accc098aceecb86b36eaa7829b2a17b1c9",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)

# go_tests
Testing Go builds and CI

Uses dep for dependency management

- `dep ensure` - Update dependencies
- `gazelle update-repos -from_file Gopkg.lock` - Updates WORKSPACE with dependencies for Bazel
- `go build -ldflags '-X github.com/omussell/go_tests/cmd.version=1.0'` - Builds the binary using the specific version flag

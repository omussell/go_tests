# go_tests
Testing Go builds and CI

Uses dep for dependency management

`dep ensure` - Update dependencies
`gazelle update-repos -from_file Gopkg.lock` - Updates WORKSPACE with dependencies for Bazel

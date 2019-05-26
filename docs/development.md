# Development

## Go

In order to work on the provider, [Go](http://www.golang.org) should be installed first (version 1.11+ is *required*).
[goenv](https://github.com/syndbg/goenv) and [gvm](https://github.com/moovweb/gvm) are great utilities that can help a
lot with that and simplify setup tremendously. 
[GOPATH](http://golang.org/doc/code.html#GOPATH) should be setup correctly and `$GOPATH/bin` should be
added `$PATH`.

This plugin uses Go modules available starting from Go `1.11` and therefore it **should not** be checked out within `$GOPATH` tree.

## Source Code

Source code can be retrieved with `git`
```bash
$ git clone git@github.com:ashald/terraform-provider-yaml.git .
```

## Dependencies

This project uses `go mod` to manage its dependencies and it's expected that all dependencies are vendored so that
it's buildable without internet access. When adding/removing a dependency run following commands:
```bash
$ go mod venndor
$ go mod tidy
```

### Test

```bash
$ make test
  GOPROXY="off" GOFLAGS="-mod=vendor" go test -v ./...
  ?   	github.com/ashald/terraform-provider-yaml	[no test files]
  === RUN   TestListOfStringsDataSource
  --- PASS: TestListOfStringsDataSource (0.03s)
  === RUN   TestMapOfStringsDataSource
  --- PASS: TestMapOfStringsDataSource (0.06s)
  === RUN   TestYamlToJsonDataSource
  --- PASS: TestYamlToJsonDataSource (0.02s)
  === RUN   TestProvider
  --- PASS: TestProvider (0.00s)
  PASS
  ok  	github.com/ashald/terraform-provider-yaml/yaml	(cached)
  GOPROXY="off" GOFLAGS="-mod=vendor" go vet ./...
```

## Build
In order to build plugin for the current platform use [GNU]make:
```bash
$ make build
  GOPROXY="off" GOFLAGS="-mod=vendor" go build -o terraform-provider-yaml_v2.1.0

```

it will build provider from sources and put it into current working directory.

If Terraform was installed (as a binary) or via `go get -u github.com/hashicorp/terraform` it'll pick up the plugin if 
executed against a configuration in the same directory.

## Release

In order to prepare provider binaries for all platforms:
```bash
$ make release
  GOPROXY="off" GOFLAGS="-mod=vendor" GOOS=darwin GOARCH=amd64 go build -o './release/terraform-provider-yaml_v2.1.0-darwin-amd64'
  GOPROXY="off" GOFLAGS="-mod=vendor" GOOS=linux GOARCH=amd64 go build -o './release/terraform-provider-yaml_v2.1.0-linux-amd64'
  GOPROXY="off" GOFLAGS="-mod=vendor" GOOS=windows GOARCH=amd64 go build -o './release/terraform-provider-yaml_v2.1.0-windows-amd64'
```

## Versioning

This project follow [Semantic Versioning](https://semver.org/)

## Changelog

This project follows [keep a changelog](https://keepachangelog.com/en/1.0.0/) guidelines for changelog.
# Terraform YAML Data Source

## Overview

This provider defines a Terraform data source that can consume YAML input [as a string] and parse a map out of it.
Map keys and values should be both strings.
If map values are not strings they are going to be serialized using YAML flow style (resembles JSON but less restrictive).
Nested maps can be optionally flattened with a given separator.

Please note that JSON is subset of YAML and therefore this data source can be used to parse arbitrary JSON as well.

## Data Sources

This plugin defines following data sources:
* `yaml`

## Reference

### Arguments

The following arguments are supported:

* `input` - (Required) A string that will be parsed. Should be a valid YAML, YAML flow or JSON.
* `flatten` - (Optional) A string that should be used as a separator in order to flatten nested maps. If set to a
non-empty value nested maps going to be flattened.  

### Attributes

The following attribute is exported:

* `output` - Map with keys and values as strings. 

## Limitations

### Boolean Conversion

The underlying parser converts `y` into a boolean that can be serialized as a true. For instance `{foo: y}` will be parsed
as `{foo: true}`. In order to avoid such behavior prefer using explicit string declaration with quotes: `{foo: "y"}`.  


## Installation

> Terraform automatically discovers the Providers when it parses configuration files.
> This only occurs when the init command is executed.

Currently Terraform is able to automatically download only
[official plugins distributed by HashiCorp](https://github.com/terraform-providers).

[All other plugins](https://www.terraform.io/docs/providers/type/community-index.html) should be installed manually.

> Terraform will search for matching Providers via a
> [Discovery](https://www.terraform.io/docs/extend/how-terraform-works.html#discovery) process, **including the current
> local directory**.

This means that the plugin should either be placed into current working directory where Terraform will be executed from
or it can be [installed system-wide](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

## Usage

### main.tf
```hcl
locals {input="input.yaml"}

data "yaml" "normal" { input = "${file(local.input)}"             }
data "yaml" "flat"   { input = "${file(local.input)}" flatten="/" }

output "normal" { value = "${data.yaml.normal.output}" }
output "flat"   { value = "${data.yaml.flat.output}"   }
```

### Download
```bash
$ wget "https://github.com/ashald/terraform-provider-yaml/releases/download/v1.0.0/terraform-provider-yaml_v1.0.0-$(uname -s | tr '[:upper:]' '[:lower:]')-amd64"
$ chmod +x ./terraform-provider-yaml*
```

### Init
```bash
$ ls -1
  main.tf
  terraform-provider-yaml_v1.0.0-linux-amd64

$ terraform init
  
  Initializing provider plugins...
  
  Terraform has been successfully initialized!
  
  You may now begin working with Terraform. Try running "terraform plan" to see
  any changes that are required for your infrastructure. All Terraform commands
  should now work.
  
  If you ever set or change modules or backend configuration for Terraform,
  rerun this command to reinitialize your working directory. If you forget, other
  commands will detect it and remind you to do so if necessary.
```

### Apply

```bash
$ cat > input.yaml << EOF
a:
  b:
    c: foobar
list:
 - foo
 - bar
EOF

$ terraform apply
  data.yaml.normal: Refreshing state...
  data.yaml.flat: Refreshing state...
  
  Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
  
  Outputs:
  
  flat = {
    a/b/c = foobar
    list = [foo, bar]
  }
  normal = {
    a = {b: {c: foobar}}
    list = [foo, bar]
  }

```


## Development

### Go

In order to work on the provider, [Go](http://www.golang.org) should be installed first (version 1.8+ is *required*).
[goenv](https://github.com/syndbg/goenv) and [gvm](https://github.com/moovweb/gvm) are great utilities that can help a
lot with that and simplify setup tremendously. 
[GOPATH](http://golang.org/doc/code.html#GOPATH) should be setup correctly and as long as `$GOPATH/bin` should be
added `$PATH`.

### Source Code

Source code can be retrieved either with `go get`

```bash
$ go get -u -d github.com/ashald/terraform-provider-yaml
```

or with `git`
```bash
$ mkdir -p ${GOPATH}/src/github.com/ashald/terraform-provider-yaml
$ cd ${GOPATH}/src/github.com/ashald/terraform-provider-yaml
$ git clone git@github.com:ashald/terraform-provider-yaml.git .
```

### Test

```bash
$ make test
  go test -v ./...
  ?   	github.com/ashald/terraform-provider-yaml	[no test files]
  === RUN   TestYamlDataSource
  --- PASS: TestYamlDataSource (0.06s)
  === RUN   TestProvider
  --- PASS: TestProvider (0.00s)
  PASS
  ok  	github.com/ashald/terraform-provider-yaml/yaml	0.092s
  go vet ./...
```

### Build
In order to build plugin for the current platform use [GNU]make:
```bash
$ make build
  go build -o terraform-provider-yaml.0.0

```

it will build provider from sources and put it into current working directory.

If Terraform was installed (as a binary) or via `go get -u github.com/hashicorp/terraform` it'll pick up the plugin if 
executed against a configuration in the same directory.

### Release

In order to prepare provider binaries for all platforms:
```bash
$ make release
  GOOS=darwin GOARCH=amd64 go build -o './release/terraform-provider-yaml_v1.0.0-darwin-amd64'
  GOOS=linux GOARCH=amd64 go build -o './release/terraform-provider-yaml_v1.0.0-linux-amd64'
```

### Versioning

This project follow [Semantic Versioning](https://semver.org/)

### Changelog

This project follows [keep a changelog](https://keepachangelog.com/en/1.0.0/) guidelines for changelog.

### Contributors

Please see [CONTRIBUTORS.md](./CONTRIBUTORS.md)

## License

This is free and unencumbered software released into the public domain. See [LICENSE](./LICENSE)

# Terraform YAML Data Source (OBSOLETE)

> As of [Terraform v0.12.2](https://github.com/hashicorp/terraform/releases/tag/v0.12.2) there is a native function `yamldecode` that obsoletes the plugin.
> No further development and/or support is planned for this repo. Feel free to open an issue if you have development ideas related to YAML processing that
> were not included into Terraform core.

## Overview

This provider defines a Terraform data sources that can consume YAML input [as a string] and covert it into another
format that can be used with Terraform.

Please note that JSON is subset of YAML and therefore this data source can be used to parse arbitrary JSON as well.

As of **Terraform 0.12** it's trivial to process YAML documents of arbitrary complexity with
[`data "yaml_to_json"`](./docs/data_source_yaml_to_json.md) and [`jsondecode`](https://www.terraform.io/docs/configuration/functions/jsondecode.html).

### Data Sources:
* [data "yaml_to_json"](./docs/data_source_yaml_to_json.md) - converts YAML to JSON for processing with [`jsondecode`](https://www.terraform.io/docs/configuration/functions/jsondecode.html). 
* [data "yaml_map_of_strings"](./docs/data_source_yaml_map_of_strings.md) - parse a YAML map as a map of string keys and values.
* [data "yaml_list_of_strings"](./docs/data_source_yaml_list_of_strings.md) - parse a YAML list as a list of strings.

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

The simplest way to get started is:
```bash
wget "https://github.com/ashald/terraform-provider-yaml/releases/download/v2.1.0/terraform-provider-yaml_v2.1.0-$(uname -s | tr '[:upper:]' '[:lower:]')-amd64"
chmod +x ./terraform-provider-yaml*
```

## Development

Provider is written and maintained by [Borys Pierov](https://github.com/Ashald).
Contributions are welcome and should follow [development guidelines](./docs/development.md).
All contributors are honored in [CONTRIBUTORS.md](./CONTRIBUTORS.md).

## License

This is free and unencumbered software released into the public domain. See [LICENSE](./LICENSE)

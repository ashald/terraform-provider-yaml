# data "yaml_map_of_strings"

## Overview
A data source to parse a YAML map as a map of strings.
Complex values are serialized in flow-style YAML and in case map contains nested maps they can be "flattened".  

Originally created before release of Terraform `v0.12` in combination [`data "yaml_list_of_strings"`](./data_source_yaml_list_of_strings.md)
can be used to process YAML documents of arbitrary complexity.

As of Terraform `v0.12` this data source may appear of less use as it can be superseded by combination of
[`data "yaml_to_json"`](./data_source_yaml_to_json.md) and [`jsondecode`](https://www.terraform.io/docs/configuration/functions/jsondecode.html). 
    

## Arguments

The following arguments are supported:

* `input` - (Required, String) An input be parsed. Should be a valid YAML, YAML flow or JSON.
* `flatten` - (Optional, String) A separator to use in order to flatten nested maps. If set to a non-empty value nested maps going to be flattened.

## Attributes

The following attribute is exported:

* `output` - Map with keys and values as strings.

## Limitations

### Boolean Conversion

The underlying parser converts `y` into a boolean that can be serialized as a true. For instance `{foo: y}` will be parsed
as `{foo: true}`. In order to avoid such behavior prefer using explicit string declaration with quotes: `{foo: "y"}`.

## Usage Example

### Configure
```hcl
locals {
  doc = <<EOF
a:
  b:
    c: foobar
list:
 - foo
 - bar
EOF
}

data "yaml_map_of_strings" "normal" {
  input = "${local.doc}"
}

data "yaml_map_of_strings" "flat" {
  input = "${local.doc}"
  flatten =" /"
}

output "normal" {
  value = "${data.yaml_map_of_strings.normal.output}"
}

output "flat" {
  value = "${data.yaml_map_of_strings.flat.output}"
}

```

### Apply
```bash
$ terraform apply
  data.yaml_map_of_strings.normal: Refreshing state...
  data.yaml_map_of_strings.flat: Refreshing state...
  
  Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
  
  Outputs:
  
  flat = {
    "a /b /c" = "foobar"
    "list" = "[foo, bar]"
  }
  normal = {
    "a" = "{b: {c: foobar}}"
    "list" = "[foo, bar]"
  }
```

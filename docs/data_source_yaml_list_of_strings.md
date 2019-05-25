# data "yaml_list_of_strings"

## Overview
A data source to parse a `yaml` list as a list of strings.
Complex values are serialized in flow-style YAML.  

Originally created before release of Terraform `v0.12` in combination [`data "yaml_map_of_strings"`](./data_source_yaml_map_of_strings.md)
can be used to process YAML documents of arbitrary complexity.

As of Terraform `v0.12` this data source may appear of less use as it can be superseded by combination of
[`data "yaml_to_json"`](./data_source_yaml_to_json.md) and [`jsondecode`](https://www.terraform.io/docs/configuration/functions/jsondecode.html). 


## Arguments

The following arguments are supported:

* `input` - (Required, String) An input be parsed. Should be a valid YAML, YAML flow or JSON.

## Attributes

The following attribute is exported:

* `output` - List of strings.

## Limitations

### Boolean Conversion

The underlying parser converts `y` into a boolean that can be serialized as a true. For instance `{foo: y}` will be parsed
as `{foo: true}`. In order to avoid such behavior prefer using explicit string declaration with quotes: `{foo: "y"}`.

## Usage Example

### Configure
```hcl
data "yaml_list_of_strings" "doc" {
  input = <<EOF
 - foo
 - 123
 - bar: 456
EOF
}

output "result" {
  value = data.yaml_list_of_strings.doc.output
}
```

### Apply
```bash
$ terraform apply
  data.yaml_list_of_strings.doc: Refreshing state...
  
  Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
  
  Outputs:
  
  result = [
    "foo",
    "123",
    "{bar: 456}",
  ]
```

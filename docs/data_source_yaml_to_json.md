# data "yaml_to_json"

## Overview
A data source to convert `yaml` to `json` that is meant to be used with
[`jsondecode`](https://www.terraform.io/docs/configuration/functions/jsondecode.html) in Terraform 0.12.

Terraform v0.12 comes with a support for complex types but there doesn't seem to be a way to leverage dynamic types on
the level of providers as they have to define static schemas for all resources and data sources. This restriction doesn't apply
to built-in functions though and therefore one can generate a "dynamic type" with `jsondecode`. The `yaml_to_json` resources
bridges the gap all the way to YAML by providing means of converting YAML to JSON so that it can be parsed with `jsondecode`.  

The `yaml_to_json` data source works even with earlier versions of Terraform but one can leverage `jsondecode` only starting from `v0.12`.

## Arguments

The following arguments are supported:

* `input` - (Required, String) YAML document to process

## Attributes

The following attribute is exported:

* `output` - (String) JSON representation of the YAML input

## Limitations

### Boolean Conversion

The underlying parser converts `y` into a boolean that can be serialized as a true. For instance `{foo: y}` will be parsed
as `{foo: true}`. In order to avoid such behavior prefer using explicit string declaration with quotes: `{foo: "y"}`.

## Usage Example

### Configure
```hcl
data "yaml_to_json" "doc" {
  input = <<EOF
foo: 123
456: bar
columns:
  - name: column0
    type: boolean
  - name: column1
    type: integer
EOF

}

output "json" {
  value = "${data.yaml_to_json.doc.output}
}

output "result" {
  value = "jsondecode${(data.yaml_to_json.doc.output)}"
}

```

### Apply
```bash
$ terraform apply
  data.yaml_to_json.doc: Refreshing state...
  
  Apply complete! Resources: 0 added, 0 changed, 0 destroyed.
  
  Outputs:
  
  json = {"456":"bar","columns":[{"name":"column0","type":"boolean"},{"name":"column1","type":"integer"}],"foo":123}
  result = {
    "456" = "bar"
    "columns" = [
      {
        "name" = "column0"
        "type" = "boolean"
      },
      {
        "name" = "column1"
        "type" = "integer"
      },
    ]
    "foo" = 123
  }
```

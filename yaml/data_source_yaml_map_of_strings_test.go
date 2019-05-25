package yaml

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
)

const mapInputNoFlatten = `
output "result" { value="${data.yaml_map_of_strings.doc.output}" }

data "yaml_map_of_strings" "doc" {
      input = <<EOF
a:
  b:
    c: foobar
list:
 - foo
 - bar
EOF
}
`

const mapInputEmptyFlatten = `
output "result" { value="${data.yaml_map_of_strings.doc.output}" }

data "yaml_map_of_strings" "doc" {
      flatten = ""
      input = <<EOF
a:
  b:
    c: foobar
list:
 - foo
 - bar
EOF
}
`

const mapInputFlattenBySlash = `
output "result" { value="${data.yaml_map_of_strings.doc.output}" }

data "yaml_map_of_strings" "doc" {
      flatten = "/"
      input = <<EOF
a:
  b:
    c: foobar
list:
 - foo
 - bar
EOF
}
`

const mapInputKeyWithMultiLineString = `
output "result" { value="${data.yaml_map_of_strings.doc.output}" }

data "yaml_map_of_strings" "doc" {
      input = <<EOF
foo: |
  foo
  bar
  baz
EOF
}
`

const mapInputNil = `
output "result" { value="${data.yaml_map_of_strings.doc.output}" }

data "yaml_map_of_strings" "doc" {
      input = <<EOF
empty_key: 
EOF
}
`

func TestMapOfStringsDataSource(t *testing.T) {
	flattenedOutput := map[string]string{"a/b/c": "foobar", "list": "[foo, bar]"}
	nonFlattenedOutput := map[string]string{"a": "{b: {c: foobar}}", "list": "[foo, bar]"}
	keyMultipleLineVal := map[string]string{"foo": "foo\nbar\nbaz\n"}
	inputNilOutput := map[string]string{"empty_key": ""}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: mapInputNoFlatten,
				Check: resource.ComposeTestCheckFunc(
					testMapOutputEquals("result", nonFlattenedOutput),
				),
			},
			{
				Config: mapInputEmptyFlatten,
				Check: resource.ComposeTestCheckFunc(
					testMapOutputEquals("result", nonFlattenedOutput),
				),
			},
			{
				Config: mapInputFlattenBySlash,
				Check: resource.ComposeTestCheckFunc(
					testMapOutputEquals("result", flattenedOutput),
				),
			},
			{
				Config: mapInputKeyWithMultiLineString,
				Check: resource.ComposeTestCheckFunc(
					testMapOutputEquals("result", keyMultipleLineVal),
				),
			},
			{
				Config: mapInputNil,
				Check: resource.ComposeTestCheckFunc(
					testMapOutputEquals("result", inputNilOutput),
				),
			},
		},
	})
}

func testMapOutputEquals(name string, expected map[string]string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		output := state.RootModule().Outputs[name]

		if output == nil {
			return fmt.Errorf("missing '%s' output", name)
		}

		outputMap := make(map[string]string)

		for k, v := range output.Value.(map[string]interface{}) {
			outputMap[k] = v.(string)
		}

		if !reflect.DeepEqual(outputMap, expected) {
			return fmt.Errorf("output '%s' value '%v' does not match expected '%v'", name, output.Value, expected)
		}
		return nil
	}
}

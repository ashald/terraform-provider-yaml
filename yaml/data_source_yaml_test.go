package yaml

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
)

const inputNoFlatten = `
output "result" { value="${data.yaml.doc.output}" }

data "yaml" "doc" {
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

const inputEmptyFlatten = `
output "result" { value="${data.yaml.doc.output}" }

data "yaml" "doc" {
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

const inputFlattenBySlash = `
output "result" { value="${data.yaml.doc.output}" }

data "yaml" "doc" {
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

func TestYamlDataSource(t *testing.T) {
	flattenedOutput := map[string]string{"a/b/c": "foobar", "list": "[foo, bar]"}
	nonFlattenedOutput := map[string]string{"a": "{b: {c: foobar}}", "list": "[foo, bar]"}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: inputNoFlatten,
				Check: resource.ComposeTestCheckFunc(
					testOutputEquals("result", nonFlattenedOutput),
				),
			},
			{
				Config: inputEmptyFlatten,
				Check: resource.ComposeTestCheckFunc(
					testOutputEquals("result", nonFlattenedOutput),
				),
			},
			{
				Config: inputFlattenBySlash,
				Check: resource.ComposeTestCheckFunc(
					testOutputEquals("result", flattenedOutput),
				),
			},
		},
	})
}

func testOutputEquals(name string, expected map[string]string) resource.TestCheckFunc {
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

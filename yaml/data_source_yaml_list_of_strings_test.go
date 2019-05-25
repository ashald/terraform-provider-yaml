package yaml

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
)

const inputListOfStrings = `
output "result" { value="${data.yaml_list_of_strings.doc.output}" }

data "yaml_list_of_strings" "doc" {
      input = <<EOF
 - foo
 - bar
EOF
}
`

const inputListOfMaps = `
output "result" { value="${data.yaml_list_of_strings.doc.output}" }

data "yaml_list_of_strings" "doc" {
      input = <<EOF
 - foo: 123
 - bar: 456
EOF
}
`

func TestListOfStringsDataSource(t *testing.T) {
	expectedListOfStrings := []string{"foo", "bar"}
	expectedListOfMaps := []string{"{foo: 123}", "{bar: 456}"}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: inputListOfStrings,
				Check: resource.ComposeTestCheckFunc(
					testListOutputEquals("result", expectedListOfStrings),
				),
			},
			{
				Config: inputListOfMaps,
				Check: resource.ComposeTestCheckFunc(
					testListOutputEquals("result", expectedListOfMaps),
				),
			},
		},
	})
}

func testListOutputEquals(name string, expected []string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		output := state.RootModule().Outputs[name]

		if output == nil {
			return fmt.Errorf("missing '%s' output", name)
		}

		var outputList []string

		for _, v := range output.Value.([]interface{}) {
			outputList = append(outputList, v.(string))
		}

		if !reflect.DeepEqual(outputList, expected) {
			return fmt.Errorf("output '%s' value '%v' does not match expected '%v'", name, output.Value, expected)
		}
		return nil
	}
}

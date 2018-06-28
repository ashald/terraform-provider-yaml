package yaml

import (
	"testing"

	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
)

const listInput = `
output "result" { value="${data.yaml_list_of_strings.doc.output}" }

data "yaml_list_of_strings" "doc" {
      input = <<EOF
 - foo
 - bar
EOF
}
`

func TestListOfStringsDataSource(t *testing.T) {
	expected := []string{"foo", "bar"}

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: listInput,
				Check: resource.ComposeTestCheckFunc(
					testListOutputEquals("result", expected),
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

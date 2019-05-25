package yaml

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

const yamlToJsonInputList = `
output "result" { value=jsondecode(data.yaml_to_json.doc.output) }

data "yaml_to_json" "doc" {
      input = <<EOF
- foo
- bar

EOF
}
`

const yamlToJsonInputMap = `
output "result" { value=jsondecode(data.yaml_to_json.doc.output) }

data "yaml_to_json" "doc" {
      input = <<EOF
foo: 123
456: bar

EOF
}
`

func TestYamlToJsonDataSource(t *testing.T) {
	expectedOutputList := `["foo","bar"]`
	expectedOutputMap := `{"456":"bar","foo":123}`  // likely unstable, consider comparing after jesondecode

	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: yamlToJsonInputList,
				Check:  resource.TestCheckResourceAttr("data.yaml_to_json.doc", "output", expectedOutputList),
			},
			{
				Config: yamlToJsonInputMap,
				Check:  resource.TestCheckResourceAttr("data.yaml_to_json.doc", "output", expectedOutputMap),
			},
		},
	})
}

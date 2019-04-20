package yaml

import (
	yml "github.com/ashald/yaml"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceList() *schema.Resource {
	return &schema.Resource{
		Read: readYamlList,

		Schema: map[string]*schema.Schema{
			// "Inputs"
			FieldInput: {
				Type:     schema.TypeString,
				Required: true,
			},
			// "Outputs"
			FieldOutput: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readYamlList(d *schema.ResourceData, m interface{}) error {
	input := d.Get(FieldInput).(string)

	var parsed []interface{}

	err := yml.Unmarshal([]byte(input), &parsed)
	if err != nil {
		return err
	}

	var result []string

	for _, value := range parsed {
		serialized, err := serializeToFlowStyleYaml(value)
		if err != nil {
			return err
		}
		result = append(result, serialized)
	}

	d.Set(FieldOutput, result)
	d.SetId(getSHA256(input))

	return nil
}

package yaml

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	yml "gopkg.in/ashald/yaml.v2"
)

func dataSourceYamlToJson() *schema.Resource {
	return &schema.Resource{
		Read: readYamToJson,

		Schema: map[string]*schema.Schema{
			// "Inputs"
			FieldInput: {
				Type:     schema.TypeString,
				Required: true,
			},
			// "Outputs"
			FieldOutput: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readYamToJson(d *schema.ResourceData, m interface{}) error {
	input := d.Get(FieldInput).(string)

	var parsed interface{}

	err := yml.Unmarshal([]byte(input), &parsed)
	if err != nil {
		return err
	}

	result, err := json.Marshal(parsed)
	if err != nil {
		return err
	}

	err = d.Set(FieldOutput, string(result))
	if err != nil {
		return err
	}

	d.SetId(getSHA256(input))

	return nil
}

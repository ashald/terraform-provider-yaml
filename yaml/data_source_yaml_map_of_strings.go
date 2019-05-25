package yaml

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	yml "gopkg.in/ashald/yaml.v2"
	"reflect"
)

func dataSourceMap() *schema.Resource {
	return &schema.Resource{
		Read: readYamlMap,

		Schema: map[string]*schema.Schema{
			// "Inputs"
			FieldInput: {
				Type:     schema.TypeString,
				Required: true,
			},
			FieldFlatten: {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "Outputs"
			FieldOutput: {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func readYamlMap(d *schema.ResourceData, m interface{}) error {
	input := d.Get(FieldInput).(string)
	separatorRaw, shouldFlatten := d.GetOk(FieldFlatten)
	separator := separatorRaw.(string)

	parsed := make(map[string]interface{})

	err := yml.Unmarshal([]byte(input), &parsed)
	if err != nil {
		return err
	}

	result := make(map[string]string)

	if shouldFlatten {
		for key, value := range parsed {
			err = flattenValue(result, reflect.ValueOf(value), key, separator)
			if err != nil {
				return err
			}
		}
	} else {
		for key, value := range parsed {
			serialized, err := serializeToFlowStyleYaml(value)
			if err != nil {
				return err
			}
			result[key] = serialized
		}
	}

	d.Set(FieldOutput, result)
	d.SetId(getSHA256(input))

	return nil
}

func flattenValue(result map[string]string, v reflect.Value, prefix string, separator string) error {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	// For empty values we don't have anything to flatten and bail out early to
	// prevent panic when calling v.Interface(). We still create a value in result
	// map, but set its value to nil leaving a question how to represent null
	// value to Terraform. From what we see, currently Terraform represents it
	// in the same way as the empty string.
	if v.Kind() == reflect.Invalid {
		result[prefix] = ""
		return nil
	}

	switch v.Kind() {
	case reflect.Map:
		flattenMap(result, v, prefix, separator)
	default:
		serialized, err := serializeToFlowStyleYaml(v.Interface())
		if err != nil {
			return err
		}
		result[prefix] = serialized
	}
	return nil
}

func flattenMap(result map[string]string, v reflect.Value, prefix string, separator string) {
	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}

		if k.Kind() != reflect.String {
			panic(fmt.Sprintf("%s: map key is not string: %s", prefix, k))
		}

		newPrefix := fmt.Sprintf("%s%s%s", prefix, separator, k.String())
		flattenValue(result, v.MapIndex(k), newPrefix, separator)
	}
}

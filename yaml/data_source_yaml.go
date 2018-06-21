package yaml

import (
	"crypto/sha256"
	"fmt"
	yml "github.com/ashald/yaml"
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
	"strings"
)

const FieldInput = "input"
const FieldFlatten = "flatten"

const FieldOutput = "output"

func dataSourceYAML() *schema.Resource {
	return &schema.Resource{
		Read: readYaml,

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

func readYaml(d *schema.ResourceData, m interface{}) error {
	input := d.Get(FieldInput).(string)
	separatorRaw, shouldFlatten := d.GetOk(FieldFlatten)
	separator := separatorRaw.(string)

	parsed, err := deserializeYaml(input)
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

func deserializeYaml(input string) (map[string]interface{}, error) {
	parsed := make(map[string]interface{})

	err := yml.Unmarshal([]byte(input), &parsed)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

func serializeToFlowStyleYaml(input interface{}) (string, error) {
	var builder strings.Builder
	encoder := yml.NewEncoder(&builder)
	encoder.SetFlowStyle(true)
	encoder.SetLineWidth(-1)

	err := encoder.Encode(input)
	if err != nil {
		return "", err
	}

	err = encoder.Close()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(builder.String()), nil
}

func flattenValue(result map[string]string, v reflect.Value, prefix string, separator string) error {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
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

func getSHA256(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return fmt.Sprintf("%x", h.Sum(nil))
}

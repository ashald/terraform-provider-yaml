package yaml

import (
	"crypto/sha256"
	"fmt"
	yml "github.com/ashald/yaml"
	"reflect"
	"strings"
)

const FieldInput = "input"
const FieldFlatten = "flatten"

const FieldOutput = "output"

func serializeToFlowStyleYaml(input interface{}) (string, error) {
	inputRef := reflect.ValueOf(input)

	if inputRef.Kind() == reflect.Interface {
		inputRef = inputRef.Elem()
	}

	if inputRef.Kind() == reflect.String {
		return input.(string), nil
	}

	if inputRef.Kind() == reflect.Invalid {
		return "", nil
	}

	var builder strings.Builder
	encoder := yml.NewEncoder(&builder)
	encoder.SetFlowStyle(true)
	encoder.SetLineWidth(-1)

	err := encoder.Encode(inputRef.Interface())
	if err != nil {
		return "", err
	}

	err = encoder.Close()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(builder.String()), nil
}

func getSHA256(src string) string {
	h := sha256.New()
	h.Write([]byte(src))
	return fmt.Sprintf("%x", h.Sum(nil))
}

package yaml

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testProviders map[string]terraform.ResourceProvider

var yamlProvider *schema.Provider

func init() {
	yamlProvider = Provider().(*schema.Provider)
	testProviders = map[string]terraform.ResourceProvider{
		"yaml": yamlProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := yamlProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

package yaml

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"yaml_map_of_strings":  dataSourceMap(),
			"yaml_list_of_strings": dataSourceList(),
		},
	}
}

package terrain

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"strings"
)

type providerConfiguration struct {
	installPkgCommand   string
	installPkgArgs      []string
	installPkgNameIndex int
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"install_command": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Command to install a package",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"terrain_package": resourcePackage(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	p := d.Get("install_command").(string)
	parts := strings.Split(p, " ")

	// find the argument with the name in it
	var idx int
	for i, v := range parts[1:] {
		if strings.Contains(v, "%s") {
			idx = i
			break
		}
	}

	return providerConfiguration{
		parts[0],
		parts[1:],
		idx,
	}, nil
}

package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/webdevwilson/terraform-provider-terrain/terrain"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: terrain.Provider,
	}
	plugin.Serve(&opts)
}

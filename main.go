package main

import (
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: netscaler.Provider,
	}

	plugin.Serve(&opts)
}

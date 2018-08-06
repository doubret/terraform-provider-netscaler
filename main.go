package main

import (
	"github.com/doubret/terraform-provider-netscaler/netscaler"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: netscaler.Provider,
	}

	plugin.Serve(&opts)
}

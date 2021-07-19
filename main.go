package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jameswoolfenden/terraform-providerbridgecrew/bridgecrew"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bridgecrew.Provider,
	})
}

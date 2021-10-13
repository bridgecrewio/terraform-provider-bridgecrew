package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/jameswoolfenden/terraform-provider-bridgecrew/bridgecrew"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return bridgecrew.Provider()
		},
	})
}

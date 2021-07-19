package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	//"github.com/jameswoolfenden/terraform-provider-bridgecrew/bridgecrew/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Default:     "https://www.bridgecrew.cloud/",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIDGECREW_URL", nil),
				Description: "url for Bridgecrew",
			},
			"token": {
				Type:        schema.TypeString,
				Default:     "",
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIDGECREW_TOKEN", nil),
				Description: "Api Token for Bridgecrew",
			},
		},
		ResourcesMap:  map[string]*schema.Resource{},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	token := d.Get("token").(string)
	return nil, nil
	//return bridgecrew.New(url, token)
}

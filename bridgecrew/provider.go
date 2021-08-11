package bridgecrew

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Client To be able to customise the url
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

//Provider main object
func Provider() *schema.Provider {
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
				Description: "API Token for Bridgecrew",
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"bridgecrew_repositories":        dataSourceRepositories(),
			"bridgecrew_repository_branches": dataSourceRepositoryBranches(),
			"bridgecrew_suppressions":        dataSourceSuppressions(),
			"bridgecrew_policies":            dataSourcePolicies(),
		},
	}
}

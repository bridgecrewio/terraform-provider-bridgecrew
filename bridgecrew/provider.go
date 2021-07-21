package bridgecrew

import (
	//"context"
	"net/http"
	//	"time"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

const HostURL string = "https://www.bridgecrew.cloud/api/v1/"

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
		},
		//ConfigureContextFunc: configureProvider,
	}
}

//func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
//	token := d.Get("token").(string)

// Warning or errors can be collected in a slice type
//	var diags diag.Diagnostics

//	if (token != "") {
//		c, err := NewClient(&token)
//		if err != nil {
//			diags = append(diags, diag.Diagnostic{
//				Severity: diag.Error,
//				Summary:  "Unable to create Bridgecrew client",
//				Detail:   "Unable to authenticate api key for Bridgecrew client",
//			})

//			return nil, diags
//		}

//		return c, diags
//	}

//	diags = append(diags, diag.Diagnostic{
//		Severity: diag.Error,
//		Summary:  "Unable to create Bridgecrew client",
//		Detail:   "Unable to authenticate api key for Bridgecrew client",
//	})

//	return nil, diags
//}

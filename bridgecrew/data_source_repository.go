package bridgecrew

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"repositories": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"repository": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"defaultBranch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"isPublic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"creationDate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repositories", "https://www.bridgecrew.cloud/api/v1/"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)

	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	repositories := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&repositories)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repositories", repositories); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

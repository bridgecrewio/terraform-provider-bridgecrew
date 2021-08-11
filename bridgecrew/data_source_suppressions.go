package bridgecrew

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSuppressions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSuppressionRead,
		Schema: map[string]*schema.Schema{
			"suppressions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"suppressiontype": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationdate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policyid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accountid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resourceid": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSuppressionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := "%s/suppressions"

	client, diags, req, diagnostics, done, err := authClient(path)

	if done {
		return diagnostics
	}

	r, err := client.Do(req)

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	log.Print("All data obtained")
	Suppressions := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&Suppressions)

	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	flatRepos := flattenSuppressionData(&Suppressions)

	if err := d.Set("suppressions", flatRepos); err != nil {
		log.Fatal(reflect.TypeOf(Suppressions))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenSuppressionData(Suppressions *[]map[string]interface{}) []interface{} {
	if Suppressions != nil {
		ois := make([]interface{}, len(*Suppressions), len(*Suppressions))

		for i, Suppression := range *Suppressions {
			oi := make(map[string]interface{})
			oi["suppressiontype"] = Suppression["suppressionType"]
			oi["creationdate"] = Suppression["creationDate"]
			oi["id"] = Suppression["id"]
			oi["policyid"] = Suppression["policyId"]
			oi["comment"] = Suppression["comment"]
			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

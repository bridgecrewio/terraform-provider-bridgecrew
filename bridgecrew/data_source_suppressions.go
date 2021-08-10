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
			"suppressions": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"suppressiontype": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"creationdate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"policyid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
								"accountid": &schema.Schema{
									Type:     schema.TypeString,
									Computed: true,
								},
								"resourceid": &schema.Schema{
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

	client, diags, req, err, diagnostics, done := authClient(path)

	if done {
		return diagnostics
	}

	log.Print("Added Header")

	r, err := client.Do(req)

	log.Print("Queried")

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	log.Print("All data obtained")
	Suppressions := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&Suppressions)

	log.Print("Decoded data")

	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	log.Print(Suppressions)
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
			oi["creationdate"]    = Suppression["creationDate"]
			oi["id"]        = Suppression["id"]
			oi["policyid"]  = Suppression["policyId"]
			oi["comment"]   = Suppression["comment"]
			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

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

func dataSourcePolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyRead,

		Schema: map[string]*schema.Schema{
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Computed: true},
						"id": {
							Type:     schema.TypeString,
							Computed: true},
						"title": {
							Type:     schema.TypeString,
							Computed: true},
						"descriptivetitle": {
							Type:     schema.TypeString,
							Computed: true},
						"constructivetitle": {
							Type:     schema.TypeString,
							Computed: true},
						"severity": {
							Type:     schema.TypeString,
							Computed: true},
						"category": {
							Type:     schema.TypeString,
							Computed: true},
					},
				},
			},
			"filters": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resourcetypes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"severity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"createdby": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"provider": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"benchmarks": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := "%s/policies/table/data"
	client, diags, req, diagnostics, done, err := authClient(path)
	if done {
		return diagnostics
	}

	r, err := client.Do(req)
	log.Print("Queried")

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	policies := make([]map[string]interface{}, 0)

	log.Print("***** predecode ****")
	err = json.NewDecoder(r.Body).Decode(&policies)
	log.Print("***** postdecode ****")
	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	flatPolicies := flattenPolicyData(&policies)

	if err := d.Set("policies", flatPolicies); err != nil {
		log.Fatal(reflect.TypeOf(policies))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenPolicyData(Policies *[]map[string]interface{}) interface{} {
	//if Policies != nil {
	//ois := make([]interface{}, len(*Policies), len(*Policies))
	//filters := make(map[string]interface{})
	//filters["category"] = Policies[]
	//filters["resourcetypes"] = Policies.Filters.ResourceTypes
	//filters["severity"] = Policies.Filters.Severity
	//log.Print(Policies.Filters)
	//return filters
	//	for i, Repository := range *Policies {
	//		oi := make(map[string]interface{})
	//		oi["provider"] = Repository["Provider"]
	//		oi["id"] = Repository["ID"]
	//		oi["title"] = Repository["Title"]
	//		oi["descriptivetitle"] = Repository["DescriptiveTitle"]
	//		oi["constructivetitle"] = Repository["ConstructiveTitle"]
	//		oi["severity"] = Repository["Severity"]
	//		oi["category"] = Repository["Category"]
	//		ois[i] = oi
	//	}
	//}

	return make(map[string]interface{}, 0)
}

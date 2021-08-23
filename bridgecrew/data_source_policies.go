package bridgecrew

import (
	"context"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/karlseguin/typed"
)

func dataSourcePolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePolicyRead,

		Schema: map[string]*schema.Schema{
			"policies": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider": {
							Type:     schema.TypeString,
							Computed: false,
							Required: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"descriptivetitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"constructivetitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Required: true,
						},
						"category": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resourcetypes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type:    schema.TypeString,
								Default: "",
							},
						},
						"accountsdata": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"guideline": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iscustom": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"conditionquery": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operator": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"attribute": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cond_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_types": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"benchmarks": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"createdby": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"code": {
							Type:     schema.TypeString,
							Required: true,
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

	if err != nil {
		log.Fatal("Failed at authClient")
		return diags
	}

	if done {
		return diagnostics
	}

	r, err := client.Do(req)

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	typed, err := typed.Json(body)

	//var filters =typed.Object("filters")
	var data = typed.Maps("data")

	flatPolicies := flattenPolicyData(&data)

	if err := d.Set("policies", flatPolicies); err != nil {
		log.Fatal(reflect.TypeOf(data))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diagnostics
}

func flattenPolicyData(Policies *[]map[string]interface{}) []interface{} {
	if Policies != nil {
		ois := make([]interface{}, len(*Policies), len(*Policies))
		for i, Policy := range *Policies {
			oi := make(map[string]interface{})

			oi["provider"] = Policy["provider"] // AWS
			oi["id"] = Policy["id"]             // james_AWS_1620660945849
			oi["title"] = Policy["title"]       // new policy
			oi["descriptivetitle"] = Policy["descriptiveTitle"]
			oi["constructivetitle"] = Policy["constructiveTitle"]
			oi["severity"] = Policy["severity"] //CRITICAL
			oi["category"] = Policy["category"] // General
			oi["guideline"] = Policy["guideline"]
			oi["iscustom"] = Policy["isCustom"]
			//oi["accountsdata"] = Policy["accountsData"]     //slice of map
			//oi["conditionquery"] = Policy["conditionQuery"] //slice of map
			oi["benchmarks"] = []string{"A", "B", "C"}

			oi["resourcetypes"] = Policy["resourceTypes"]
			oi["createdby"] = Policy["createdBy"]
			oi["code"] = Policy["code"]
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

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
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"repository": {
										Required: true,
										Type:     schema.TypeString,
									},
									"amounts": {
										Required: true,
										Type:     schema.TypeMap,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"lastupdatedate": {
										Computed: true,
										Type:     schema.TypeString,
									},
								},
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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"value": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"resource_types": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type:    schema.TypeString,
											Default: "",
										},
									},
								},
							},
						},
						"benchmarks": {
							Type:     schema.TypeMap,
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
		//return diag.FromErr(err)
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	typedjson, err := typed.Json(body)

	if err != nil {
		log.Fatal("Failed at unmarshalling with typed")
	}

	//var filters =typed.Object("filters")
	var data = typedjson.Maps("data")

	flatPolicies := flattenPolicyData(&data)

	if err := d.Set("policies", flatPolicies); err != nil {
		log.Fatal(reflect.TypeOf(data))
		//return diag.FromErr(err)
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
			//accountsData:=Policy["accountsData"]

			condition := make(map[string]interface{})
			condition["value"] = true
			condition["operator"] = "operator"
			condition["attribute"] = "attribute"
			condition["cond_type"] = "cond_type"
			condition["resource_types"] = []string{"aws_api_gateway_api_key"}
			conditions := make([]interface{}, 1, 1)
			conditions[0] = condition
			oi["conditionquery"] = conditions
			oi["resourcetypes"] = Policy["resourceTypes"]
			oi["createdby"] = Policy["createdBy"]
			oi["code"] = Policy["code"]
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func printtypes(accountsData interface{}) {
	m := accountsData.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			log.Println(k, "is string", vv)
		case float64:
			log.Println(k, "is float64", vv)
		case []interface{}:
			log.Println(k, "is an array:")
			for i, u := range vv {
				log.Println(i, u)
			}
		default:
			log.Println(k, "is of a type I don't know how to handle")
		}
	}
}

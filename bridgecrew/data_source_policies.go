package bridgecrew

import (
	"context"
	"encoding/json"
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
											Type: schema.TypeInt,
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
							Type:     schema.TypeString,
							Required: true,
						},
						"benchmarks": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"benchmark": {
										Type:     schema.TypeString,
										Required: true,
									},
									"version": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
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

			//TODO:accountsdata
			//log.Print(Policy["accountsData"])
			var accounts []interface{}

			accountsData := Policy["accountsData"].(map[string]interface{})
			//log.Print(accountsData)
			for key, element := range accountsData {
				account := make(map[string]interface{})
				account["repository"] = key
				temp := element.(map[string]interface{})
				account["amounts"] = temp["amounts"]
				account["lastupdatedate"] = temp["lastUpdateDate"]
				log.Print(account)
				accounts = append(accounts, account)
			}

			oi["accountsdata"] = accounts

			var marks []interface{}

			benchmarks := Policy["benchmarks"].(map[string]interface{})
			for key, element := range benchmarks {
				bench := make(map[string]interface{})
				bench["benchmark"] = key
				bench["version"] = element
				marks = append(marks, bench)
			}

			oi["benchmarks"] = marks

			// this was hard... dynamic schema!
			u, err := json.Marshal(Policy["conditionQuery"])
			if err != nil {
				panic(err)
			}

			oi["conditionquery"] = string(u)
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
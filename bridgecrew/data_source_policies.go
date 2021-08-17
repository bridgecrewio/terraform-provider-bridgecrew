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
							Required: false,
							Optional: true,
							Default:  "aws"},
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
						"resourcetypes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type:    schema.TypeString,
								Default: "",
							},
						},
						//"accountsData":{},struct
						"guideline": {Type: schema.TypeString,
							Computed: true},
						"iscustom": {Type: schema.TypeBool,
							Computed: true},
						//"conditionQuery":struct
						//benchmarks: struct
						"createdby": {Type: schema.TypeString,
							Computed: true},
						"code": {Type: schema.TypeString,
							Computed: true},
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

	log.Print("********PREQUEST********")
	r, err := client.Do(req)

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	log.Print("********PRETYPE********")
	typed, err := typed.Json(body)

	//var filters =typed.Object("filters")
	var data = typed.Maps("data")

	flatPolicies := flattenPolicyData(&data)

	log.Print("********PRESET********")
	log.Print(reflect.TypeOf(data))
	log.Print("********FLATTEN********")
	log.Print(reflect.TypeOf(flatPolicies))
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

			oi["provider"] = Policy["provider"]
			oi["id"] = Policy["id"]
			oi["title"] = Policy["title"]
			oi["descriptivetitle"] = Policy["descriptiveTitle"]
			oi["constructivetitle"] = Policy["constructiveTitle"]
			oi["severity"] = Policy["severity"]
			oi["category"] = Policy["category"]
			oi["resourcetypes"] = []string{"a", "b", "c"}
			oi["guideline"] = Policy["guideline"]
			oi["iscustom"] = true //Policy["isCustom"]

			//oi["resourcetypes"]=Policy["resourceTypes"]
			oi["createdby"] = Policy["createdBy"]
			oi["code"] = Policy["code"]
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

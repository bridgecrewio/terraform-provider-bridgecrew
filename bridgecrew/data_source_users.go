package bridgecrew

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,

		Schema: map[string]*schema.Schema{
			"users": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allaccountsaccess": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"alias": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"lastmodified": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"customername": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

//goland:noinspection GoUnusedParameter
func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := RequestParams{"%s/manage/users", "v1", "GET"}

	configure := m.(ProviderConfig)
	client, req, diagnostics, done := authClient(params, configure, nil)

	if done {
		return diagnostics
	}

	r, err := client.Do(req)

	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed at client.Do %s \n", err.Error()),
		})
		return diagnostics
	}

	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	var jsonMap []map[string]interface{}
	err = json.Unmarshal(body, &jsonMap)

	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed at unmarshalling with typed %s \n", err.Error()),
		})
		return diagnostics
	}

	flatUsers := flattenUserData(&jsonMap)

	if err := d.Set("users", flatUsers); err != nil {
		log.Fatal(reflect.TypeOf(jsonMap))
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diagnostics
}

func flattenUserData(users *[]map[string]interface{}) []interface{} {
	if users != nil {
		ois := make([]interface{}, len(*users))
		for i, User := range *users {
			oi := make(map[string]interface{})
			oi["role"] = User["role"]
			oi["allaccountsaccess"] = User["all_accounts_access"]
			oi["email"] = User["email"]
			oi["lastmodified"] = User["last_modified"]

			if User["accounts"] != nil {
				var accounts []interface{}
				accountsData := User["accounts"].([]interface{})
				if len(accountsData) > 0 {
					for _, element := range accountsData {
						account := make(map[string]interface{})
						temp := element.(map[string]interface{})
						account["alias"] = temp["alias"].(string)
						account["id"] = temp["id"].(string)
						accounts = append(accounts, account)
					}

					oi["accounts"] = accounts
				}
			}

			oi["customername"] = User["customer_name"]
			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

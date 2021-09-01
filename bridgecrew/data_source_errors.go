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

func dataSourceErrors() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceErrorRead,
		Schema: map[string]*schema.Schema{
			"errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationdate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"repository": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},
						"owner": {
							Type:     schema.TypeString,
							Required: true,
						},
						"defaultbranch": {
							Type:     schema.TypeString,
							Required: true,
						},
						"ispublic": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}

func dataSourceErrorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := "%s/errors/gitBlameAuthors"
	configure := m.(ProviderConfig)
	client, req, diagnostics, done, err := authClient(path, configure)

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
	errors := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&errors)

	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	flatErrors := flattenErrorData(&errors)

	if err := d.Set("errors", flatErrors); err != nil {
		log.Fatal(reflect.TypeOf(errors))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diagnostics
}

func flattenErrorData(Errors *[]map[string]interface{}) []interface{} {
	if Errors != nil {
		ois := make([]interface{}, len(*Errors), len(*Errors))

		for i, Error := range *Errors {
			oi := make(map[string]interface{})
			log.Print(Error)
			//todo
			/*oi["id"] = Repository["id"]
			oi["repository"] = Repository["repository"]
			oi["source"] = Repository["source"]
			oi["owner"] = Repository["owner"]
			oi["creationdate"] = Repository["creationDate"]
			oi["defaultbranch"] = Repository["defaultBranch"]
			oi["ispublic"] = Repository["ispublic"]*/

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

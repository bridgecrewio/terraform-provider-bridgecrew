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

func dataSourceRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"repositories": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationdate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
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
						"defaultbranch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ispublic": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := "%s/repositories"
	

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
	repositories := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&repositories)

	log.Print("Decoded data")

	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	log.Print(repositories)
	flatRepos := flattenRepositoryData(&repositories)

	if err := d.Set("repositories", flatRepos); err != nil {
		log.Fatal(reflect.TypeOf(repositories))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRepositoryData(Repositories *[]map[string]interface{}) []interface{} {
	if Repositories != nil {
		ois := make([]interface{}, len(*Repositories), len(*Repositories))

		for i, Repository := range *Repositories {
			oi := make(map[string]interface{})
			oi["id"] = Repository["id"]
			oi["repository"] = Repository["repository"]
			oi["source"] = Repository["source"]
			oi["owner"] = Repository["owner"]
			oi["creationdate"] = Repository["creationDate"]
			oi["defaultbranch"] = Repository["defaultBranch"]
			oi["ispublic"] = Repository["ispublic"]

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}

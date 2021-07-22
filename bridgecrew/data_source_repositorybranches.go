package bridgecrew

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepositoryBranches() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoryBranchRead,
		Schema: map[string]*schema.Schema{
			"repositoriesbranches": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{},
			},
		},
	}
}

func dataSourceRepositoryBranchRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	api := os.Getenv("BRIDGECREW_API")

	if api == "" {
		log.Fatal("BRIDGECREW_API is missing")
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + api

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repositories/branches", "https://www.bridgecrew.cloud/api/v1"), nil)

	if err != nil {
		log.Fatal("Failed at http")
		return diag.FromErr(err)
	}

	log.Print("Passed http Request")

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	log.Print("Added Header")

	r, err := client.Do(req)

	log.Print("Queried")

	if err != nil {
		log.Fatal("Failed at client.Do")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	log.Print("All data obtained")
	repositoriesbranches := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&repositoriesbranches)

	log.Print("Decoded data")
	log.Print(r.Body)

	if err != nil {
		log.Fatal("Failed to parse data")
		return diag.FromErr(err)
	}

	log.Print(repositoriesbranches)
	flatBranch := flattenBranchData(&repositoriesbranches)

	if err := d.Set("repositories", flatBranch); err != nil {
		log.Fatal(reflect.TypeOf(repositoriesbranches))
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
func flattenBranchData(Repositories *[]map[string]interface{}) []interface{} {
	if Repositories != nil {
		ois := make([]interface{}, len(*Repositories), len(*Repositories))

		//for i, Repository := range *Repositories {
		//	oi := make(map[string]interface{})
		//oi["name"] = Repository["name"]
		//oi["creationdate"] = Repository["creationDate"]
		//oi["defaultbranch"] = Repository["defaultBranch"]

		//	ois[i] = oi
		//}

		return ois
	}

	return make([]interface{}, 0)
}

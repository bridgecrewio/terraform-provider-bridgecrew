package bridgecrew

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/karlseguin/typed"
)

func dataSourceIncidents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIncidentRead,

		Schema: map[string]*schema.Schema{
			"incidents": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloud_provider": {
							Type:        schema.TypeString,
							Computed:    false,
							Required:    true,
							Description: "The name of the Cloud Provider",
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"category": {
							Type:     schema.TypeString,
							Required: true,
						},
						"guideline": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Required: true,
						},
						"constructive_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"descriptive_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pc_severity": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iscustom": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIncidentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	configure := m.(ProviderConfig)
	var diagnostics diag.Diagnostics

	limit, offset := 100, 0
	var data []map[string]interface{}

	for {
		path := fmt.Sprintf("/incidents?limit=%d&offset=%d", limit, offset)
		params := RequestParams{"%s" + path, "v2", "POST"}

		client, req, tempDiagnostics, done, err := authClient(params, configure)
		diagnostics = tempDiagnostics

		if err != nil {
			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failed at authClient %s \n", err.Error()),
			})
			return diagnostics
		}

		if done {
			return diagnostics
		}

		req.Header.Add("Content-Type", "application/json")
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
		typedjson, err := typed.Json(body)

		if err != nil {
			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Failed at unmarshalling with typed %s \n", err.Error()),
			})
			return diagnostics
		}

		partialData := typedjson.Maps("data")
		data = append(data, partialData...)

		hasNext := typedjson.Bool("hasNext")
		if !hasNext {
			break
		}
		offset += limit
	}

	flatIncidents := flattenIncidentData(&data)

	if err := d.Set("incidents", flatIncidents); err != nil {
		log.Fatal(reflect.TypeOf(data))
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diagnostics
}

func flattenIncidentData(Incidents *[]map[string]interface{}) []interface{} {
	if Incidents != nil {
		ois := make([]interface{}, len(*Incidents), len(*Incidents))
		for i, Incident := range *Incidents {
			oi := make(map[string]interface{})

			oi["cloud_provider"] = Incident["provider"]
			oi["id"] = Incident["incidentId"]
			oi["title"] = Incident["title"]
			oi["constructive_title"] = Incident["constructiveTitle"]
			oi["descriptive_title"] = Incident["descriptiveTitle"]
			oi["severity"] = Incident["severity"]
			oi["pc_severity"] = Incident["pcSeverity"]
			oi["category"] = Incident["category"]
			oi["guideline"] = Incident["guideline"]
			oi["iscustom"] = Incident["isCustom"]

			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

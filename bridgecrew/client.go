package bridgecrew

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

//use basic auth client
func authClient(path string, configure ProviderConfig) (*http.Client, *http.Request, diag.Diagnostics, bool, error) {

	var diags diag.Diagnostics
	api := configure.Token
	url := configure.URL

	var baseurl string
	baseurl = fmt.Sprintf(url + "/api/v1")

	if api == "" {
		log.Fatal("BRIDGECREW_API is missing")
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + api

	client := &http.Client{Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf(path, baseurl), nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Get request failed %s \n", err.Error()),
		})
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	return client, req, diags, false, err
}

// CheckStatus confirms returns codes are 200
func CheckStatus(res *http.Response) (diag.Diagnostics, bool) {
	if res.StatusCode != http.StatusOK {
		errStr := fmt.Errorf("Non-OK HTTP status: %d", res.StatusCode)
		return diag.FromErr(errStr), true
	}

	return nil, false
}

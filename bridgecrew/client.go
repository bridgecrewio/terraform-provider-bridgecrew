package bridgecrew

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

//use basic auth client
func authClient(path string) (*http.Client, diag.Diagnostics, *http.Request, error, diag.Diagnostics, bool) {
	api := os.Getenv("BRIDGECREW_API")

	if api == "" {
		log.Fatal("BRIDGECREW_API is missing")
	}

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + api

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	log.Print(fmt.Sprintf(path, "https://www.bridgecrew.cloud/api/v1"))
	req, err := http.NewRequest("GET", fmt.Sprintf(path, "https://www.bridgecrew.cloud/api/v1"), nil)

	if err != nil {
		log.Fatal("Failed at http")
		return nil, nil, nil, nil, diag.FromErr(err), true
	}

	log.Print("Passed http Request")

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	return client, diags, req, err, nil, false
}

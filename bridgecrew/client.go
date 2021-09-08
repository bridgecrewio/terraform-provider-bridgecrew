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

	highlight(fmt.Sprintf(path, baseurl))

	if err != nil {
		log.Fatal("Failed at http")
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	return client, req, nil, false, err
}

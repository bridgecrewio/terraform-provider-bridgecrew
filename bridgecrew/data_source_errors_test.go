package bridgecrew

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceErrors(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceErrors(),
			},
		},
	})
}

func testAccDataSourceErrors() string {
	return fmt.Sprintf(
		`
data "bridgecrew_errors" "test" {
}`)
}

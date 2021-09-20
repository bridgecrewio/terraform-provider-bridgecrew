package bridgecrew

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicies(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicies(),
			},
		},
	})
}

func testAccDataSourcePolicies() string {
	return fmt.Sprintf(
		`
data "bridgecrew_policies" "test" {
}`)
}

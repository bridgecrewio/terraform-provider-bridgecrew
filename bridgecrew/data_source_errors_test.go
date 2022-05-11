package bridgecrew

//func TestAccDataSourceErrors(t *testing.T) {
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { testAccPreCheck(t) },
//		ProviderFactories: testAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccDataSourceErrors(),
//				Check:  resource.ComposeAggregateTestCheckFunc(
//				// resource.TestCheckResourceAttrSet("data.bridgecrew_errors.test", "id"),
//				),
//			},
//		},
//	})
//}

func testAccDataSourceErrors() string {
	return `
	data "bridgecrew_errors" "test" {
	}`
}

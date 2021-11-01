

resource "bridgecrew_policy" "new" {
  cloud_provider = "bogus"

  file = "${path.module}/policy/isnot.yaml"

  // although benchmarks take a free text this is total ***, as it needs to be an existing benchmark as
  // does the version, and that more like a category than anything
  benchmarks {
    cis_aws_v12 = ["1.1", "2.1"]
    //cis_aws_v13 = ["1.3", "2.4"]
  }

}



resource "bridgecrew_simple_policy" "new" {
  cloud_provider = "bogus"
  title          = "is too short"
  severity       = "safe"
  category       = "jimmy"

  // For now only one condition block is valid
  conditions {
    resource_types = ["aws_s3_bucket", "aws_instance"]
    cond_type      = "attribute"
    attribute      = "bucket"
    operator       = "not_equals"
    value          = "jimbo2"
  }

  guidelines = "Way too lazy to write docs"

  // although benchmarks take a free text this is total ***, as it needs to be an existing benchmark as
  // does the version, and that more like a category than anything
  benchmarks {
    cis_aws_v12 = ["1.1", "2.1"]
    //cis_aws_v13 = ["1.3", "2.4"]
  }

}


resource "random_string" "new" {
  length  = 8
  special = false
}

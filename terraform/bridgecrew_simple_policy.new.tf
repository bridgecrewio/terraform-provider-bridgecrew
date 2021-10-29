

resource "bridgecrew_simple_policy" "new" {
  count          = 1
  cloud_provider = "aws"
  title          = "my first test ${count.index} ${random_string.new.id}"
  severity       = "critical"
  category       = "logging"

  // For now only one condition block is valid
  conditions {
    resource_types = ["aws_s3_bucket", "aws_instance"]
    cond_type      = "attribute"
    attribute      = "bucket"
    operator       = "not_equals"
    value          = "jimbo2"
  }

  guidelines = "This should explain a lot more"

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

resource "bridgecrew_simple_policy" "tagging" {
  cloud_provider = "all"
  title          = "Check that all resources have a yor tag"
  severity       = "critical"
  category       = "general"
  guidelines     = "Use (yor.io)[yor.io](yor.io)!"

  conditions {
    resource_types = ["all"]
    cond_type      = "attribute"
    attribute      = "tags.yor_trace"
    operator       = "exists"
  }
}

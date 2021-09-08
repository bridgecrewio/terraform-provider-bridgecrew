provider "bridgecrew" {
  url   = "https://9vk79w5qvc.execute-api.us-west-2.amazonaws.com/v1"
  token = "af7147be-f48a-4c4e-9e50-c32356cde608"
}

//data "bridgecrew_repositories" "all" {}
//data "bridgecrew_suppressions" "all" {}
data "bridgecrew_policies" "all" {}
//data "bridgecrew_repository_branches" "all" {}

terraform {
  required_providers {
    bridgecrew = {
      version = "0.1.1"
      source  = "jameswoolfenden/dev/bridgecrew"
    }
  }
}

# output "repos" {
#   value = data.bridgecrew_repositories.all
# }


//output "suppression" {
//  value = data.bridgecrew_suppressions.all
//}

output "polices" {
  value = data.bridgecrew_policies.all
}

//output "branches" {
// value = data.bridgecrew_repository_branches.all
//}

resource "random_string" "new" {
  length  = 8
  special = false

}

resource "bridgecrew_policy" "new" {
  count          = 1
  cloud_provider = "aws"
  title          = "my first test ${count.index} ${random_string.new.id}"
  severity       = "critical"
  category       = "logging"

  //still fails for now
  code = ""

  // For now only one condition block is valid
  conditions {
    resource_types = ["aws_s3_bucket", "aws_instance"]
    cond_type      = "attribute"
    attribute      = "bucket"
    operator       = "not_equals"
    value          = "jimbo"
  }

  guidelines = "This should explain a little"

  // although benchmarks take a free text this is total BS, as it needs to be an existing benchmark as
  // does the version, and that more like a category than anything
  benchmarks {
    cis_aws_v12 = ["1.1", "2.1"]
    //cis_aws_v13 = ["1.3", "2.4"]
  }

}

output "policy" {
  value = bridgecrew_policy.new
}

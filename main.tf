provider "bridgecrew" {
  url   = "https://www.bridgecrew.cloud"
  token = "e1debacc-fb6d-5230-89f4-ec76f383d092"
}

data "bridgecrew_repositories" "all" {}
data "bridgecrew_suppressions" "all" {}
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

//output "repos" {
//  value = data.bridgecrew_repositories.all
//}


//output "suppression" {
//  value = data.bridgecrew_suppressions.all
//}

//output "polices" {
//  value = data.bridgecrew_policies.all
//}

//output "branches" {
// value = data.bridgecrew_repository_branches.all
//}


resource "bridgecrew_policy" "new" {
  count           = 2
  cloud_provider  = "aws"
  title           = "my first test ${count.index}"
  severity        = "CRITICAL"
  category        = "LOGGING"
  resource_types  = ["aws_instance", "aws_s3_bucket"]
  iscustom        = true
  code            = "echo 'hello world'"
  condition_query = "some shizzle"
  benchmarks {
    benchmark = "free text for now"
    version   = ["1.1", "1.2"]
  }
}

output "policy" {
  value = bridgecrew_policy.new
}

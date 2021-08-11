provider "bridgecrew" {
  url   = "https://www.bridgecrew.cloud/api/v1"
  token = "e1debacc-fb6d-5230-89f4-ec76f383d092"
}

data "bridgecrew_repositories" "all" {}
data "bridgecrew_suppressions" "all" {}
//data "bridgecrew_policies" "all" {}
//data "bridgecrew_repository_branches" "all" {}

terraform {
  required_providers {
    bridgecrew = {
      version = "0.1.1"
      source  = "jameswoolfenden/dev/bridgecrew"
    }
  }
}

output "repos" {
  value = data.bridgecrew_repositories.all
}


output "suppression" {
  value = data.bridgecrew_suppressions.all
}

//output "polices" {
//  value = data.bridgecrew_policies.all
//}

//output "branches" {
//  value = data.bridgecrew_repository_branches.all
//}

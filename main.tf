provider "bridgecrew" {
  url   = "https://www.bridgecrew.cloud/api/v1"
  token = "e1debacc-fb6d-5230-89f4-ec76f383d092"
}

data "bridgecrew_repositories" "all" {}


terraform {
  required_providers {
    bridgecrew = {
      version = "0.1.1"
      source  = "jameswoolfenden/dev/bridgecrew"
    }
  }
}
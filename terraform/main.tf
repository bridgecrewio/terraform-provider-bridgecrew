
data "bridgecrew_suppressions" "all" {}
data "bridgecrew_policies" "all" {}

output "suppression" {
  value = data.bridgecrew_suppressions.all
}

output "polices" {
  value = data.bridgecrew_policies.all
}

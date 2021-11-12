output "apitokens" {
  value = data.bridgecrew_apitokens.all
}

#output "branches" {
#  value = data.bridgecrew_repository_branches.all
#}

output "errors" {
  value = data.bridgecrew_errors.all
}

output "policy" {
  value = bridgecrew_policy.new
}

output "simple_policy" {
  value = bridgecrew_simple_policy.new
}


output "repos" {
  value = data.bridgecrew_repositories.all
}

output "suppression" {
  value = data.bridgecrew_suppressions.all
}

output "polices" {
  value = data.bridgecrew_policies.all
}

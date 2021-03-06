output "enforcement_rules" {
  value = data.bridgecrew_enforcement_rules.all
}

output "enforcement_rule" {
  value = data.bridgecrew_enforcement_rule.all
}

output "enforcement_accounts" {
  value = data.bridgecrew_enforcement_accounts.all
}

output "mappings" {
  value = data.bridgecrew_mappings.new
}

output "integrations" {
  value = data.bridgecrew_integrations.all
}

output "apitokens" {
  value = data.bridgecrew_apitokens.all
}

output "customer_apitokens" {
  value = data.bridgecrew_apitokens_customer.all
}

output "users" {
  value = data.bridgecrew_users.all
}

output "branches" {
  value = data.bridgecrew_repository_branches.all
}


output "authors" {
  value = data.bridgecrew_authors.all
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

output "incidents" {
  value = data.bridgecrew_incidents.all
}

output "presets" {
  value = data.bridgecrew_incidents_preset.all
}

output "info" {
  value = data.bridgecrew_incidents_info.all
}

output "organisation" {
  value = data.bridgecrew_organisation.mine
}

output "tag" {
  value = data.bridgecrew_tag.found
}

output "tags" {
  value = data.bridgecrew_tags.found
}

output "justifications" {
  value = data.bridgecrew_justifications.given
}

output "resource_tag" {
  value = bridgecrew_tag.new
}

output "policy" {
  value = bridgecrew_policy.new
}

output "simple_policy" {
  value = bridgecrew_simple_policy.new
}

output "complex_policy" {
  value = bridgecrew_complex_policy.new
}

output "code_reviews" {
  value = data.bridgecrew_code_reviews.all
}

//when it works
output "made_enforcement_rule" {
  value = bridgecrew_enforcement_rule.example
}

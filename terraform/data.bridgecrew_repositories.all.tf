data "bridgecrew_repositories" "all" {}

output "repos" {
  value = data.bridgecrew_repositories.all
}

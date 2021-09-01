{{define "data_source_repositories"}}---
layout: "bridgecrew"
page_title: "Bridgecrew: data_source_repositories"
sidebar_current: "docs-bridgecrew-data_source_repositories"

description: |-
Gets a list of all your managed repositories
---

# bridgecrew_repositories (Data Source)

Use this datasource to get the details of your managed repositories from Bridgecrew.




## Example Usage
```hcl
data "bridgecrew_repositories" "myrepositories" {
}
```
{{end}}

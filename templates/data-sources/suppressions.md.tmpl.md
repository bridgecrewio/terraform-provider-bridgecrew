{{define "data_source_suppressions"}}---
layout: "bridgecrew"
page_title: "Bridgecrew: data_source_suppressions"
sidebar_current: "docs-bridgecrew-data_source_suppressions"

description: |-
Get a list of all your policies suppressions
---

# bridgecrew_suppressions (Data Source)

Use this datasource to get the details of your policy suppressions from Bridgecrew.




## Example Usage
```hcl
data "bridgecrew_suppressions" "mysuppressions" {
}
```
{{end}}

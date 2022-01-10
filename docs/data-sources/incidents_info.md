---
layout: "bridgecrew"
page_title: "Bridgecrew: data_source_incidents_info"
sidebar_current: "docs-bridgecrew-data_source_incidents_info"

description: |-
Gets all the info and counters of the incidents and violations.
More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getinfo>.

---

# bridgecrew_incidents_info

Use this datasource to get the details of your incidents counters from Bridgecrew.




## Example Usage
```hcl
data "bridgecrew_incidents_info" "all" {}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **encryption** (Set of Object) (see [below for nested schema](#nestedatt--encryption))
- **reachability** (Set of Object) (see [below for nested schema](#nestedatt--reachability))
- **status** (Set of Object) (see [below for nested schema](#nestedatt--status))
- **total** (Number)
- **traced** (Set of Object) (see [below for nested schema](#nestedatt--traced))

<a id="nestedatt--encryption"></a>
### Nested Schema for `encryption`

Read-Only:

- **encrypted** (Number)
- **noencryption** (Number)
- **unencrypted** (Number)


<a id="nestedatt--reachability"></a>
### Nested Schema for `reachability`

Read-Only:

- **noreachability** (Number)
- **private** (Number)
- **public** (Number)


<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- **open** (Number)
- **passed** (Number)
- **suppressed** (Number)


<a id="nestedatt--traced"></a>
### Nested Schema for `traced`

Read-Only:

- **nottraced** (Number)
- **traced** (Number)
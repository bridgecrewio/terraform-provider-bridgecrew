---
layout: "bridgecrew"
page_title: "Bridgecrew: data_source_enforcement_rules"
sidebar_current: "docs-bridgecrew-data_source_enforcement_rules"

description: |-
Get a list of all your Bridgecrew platform enforcement rules.
More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/getschemeforcustomer>.

---

# bridgecrew_enforcement_rules

Use this datasource to get the details of all your enforcement rules from Bridgecrew.




## Example Usage
```hcl
data "bridgecrew_enforcement_rules" "all" {
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **accountsnotinmainrule** (Block List) (see [below for nested schema](#nestedblock--accountsnotinmainrule))
- **id** (String) The ID of this resource.

### Read-Only

- **rules** (List of Object) (see [below for nested schema](#nestedatt--rules))

<a id="nestedblock--accountsnotinmainrule"></a>
### Nested Schema for `accountsnotinmainrule`

Required:

- **account_id** (String)
- **account_name** (String)


<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- **codecategories** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories))
- **createdby** (String)
- **creationdate** (String)
- **editable** (Boolean)
- **id** (String)
- **mainrule** (Boolean)
- **name** (String)
- **repositories** (List of String)

<a id="nestedobjatt--rules--codecategories"></a>
### Nested Schema for `rules.codecategories`

Read-Only:

- **iac** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories--iac))
- **images** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories--images))
- **open_source** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories--open_source))
- **secrets** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories--secrets))
- **supply_chain** (Set of Object) (see [below for nested schema](#nestedobjatt--rules--codecategories--supply_chain))

<a id="nestedobjatt--rules--codecategories--iac"></a>
### Nested Schema for `rules.codecategories.iac`

Read-Only:

- **comments_bot_threshold** (String)
- **hard_fail_threshold** (String)
- **soft_fail_threshold** (String)


<a id="nestedobjatt--rules--codecategories--images"></a>
### Nested Schema for `rules.codecategories.images`

Read-Only:

- **comments_bot_threshold** (String)
- **hard_fail_threshold** (String)
- **soft_fail_threshold** (String)


<a id="nestedobjatt--rules--codecategories--open_source"></a>
### Nested Schema for `rules.codecategories.open_source`

Read-Only:

- **comments_bot_threshold** (String)
- **hard_fail_threshold** (String)
- **soft_fail_threshold** (String)


<a id="nestedobjatt--rules--codecategories--secrets"></a>
### Nested Schema for `rules.codecategories.secrets`

Read-Only:

- **comments_bot_threshold** (String)
- **hard_fail_threshold** (String)
- **soft_fail_threshold** (String)


<a id="nestedobjatt--rules--codecategories--supply_chain"></a>
### Nested Schema for `rules.codecategories.supply_chain`

Read-Only:

- **comments_bot_threshold** (String)
- **hard_fail_threshold** (String)
- **soft_fail_threshold** (String)

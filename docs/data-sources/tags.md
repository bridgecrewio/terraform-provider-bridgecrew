---
layout: "bridgecrew"
page_title: "Bridgecrew: data_source_tags"
sidebar_current: "docs-bridgecrew-data_source_tags"

description: |-
Get all the tag rules.
More details on the Bridgecrew API for this datasource are available <https://docs.bridgecrew.io/reference/gettag>.

---

# bridgecrew_tags

Use this datasource to get the all the tag rules.




## Example Usage
```hcl
data "bridgecrew_tags" "custom" {
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **tags** (List of Object) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **candoactions** (Boolean)
- **createdby** (String)
- **creationdate** (String)
- **definition** (String)
- **description** (String)
- **id** (String)
- **isenabled** (Boolean)
- **name** (String)
- **repositories** (List of Object) (see [below for nested schema](#nestedobjatt--tags--repositories))
- **tagruleootbid** (String)

<a id="nestedobjatt--tags--repositories"></a>
### Nested Schema for `tags.repositories`

Read-Only:

- **defaultbranch** (String)
- **id** (String)
- **name** (String)
- **owner** (String)
- **repo** (String)
- **source** (String)

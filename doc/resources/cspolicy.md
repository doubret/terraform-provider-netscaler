# netscaler_cspolicy

Terraform resource name : ```netscaler_cspolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[csaction.name](/doc/resources/csaction.md)|
|domain|No|No|string|
|logaction|No|No|string|
|policyname|No|No|string|
|rule|No|No|string|
|url|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/cspolicy/cspolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cspolicy" "tf_name" {

    action = "${netscaler_csaction.<resource_name>.name}"
    domain = "abc"
    logaction = "abc"
    policyname = "abc"
    rule = "abc"
    url = "abc"
}
```


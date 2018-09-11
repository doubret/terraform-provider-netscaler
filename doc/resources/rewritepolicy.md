# netscaler_rewritepolicy

Terraform resource name : ```netscaler_rewritepolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[rewriteaction.name](/doc/resources/rewriteaction.md)|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|rule|No|No|string|
|undefaction|No|No|[rewriteaction.name](/doc/resources/rewriteaction.md)|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/rewrite/rewritepolicy/rewritepolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_rewritepolicy" "<resource_name>" {

    action = "${netscaler_rewriteaction.<resource_name>.name}"
    comment = "abc"
    logaction = "abc"
    name = "abc"
    rule = "abc"
    undefaction = "${netscaler_rewriteaction.<resource_name>.name}"
}
```


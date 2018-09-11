# netscaler_capolicy

Terraform resource name : ```netscaler_capolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[caaction.name](/doc/resources/caaction.md)|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|rule|No|No|string|
|undefaction|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ca/capolicy/capolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_capolicy" "<resource_name>" {

    action = "${netscaler_caaction.<resource_name>.name}"
    comment = "abc"
    logaction = "abc"
    name = "abc"
    rule = "abc"
    undefaction = "abc"
}
```


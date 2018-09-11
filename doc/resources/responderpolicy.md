# netscaler_responderpolicy

Terraform resource name : ```netscaler_responderpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[responderaction.name](/doc/resources/responderaction.md)|
|appflowaction|No|No|[appflowaction.name](/doc/resources/appflowaction.md)|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|rule|No|No|string|
|undefaction|No|No|NOOP, RESET, DROP|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/responder/responderpolicy/responderpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_responderpolicy" "<resource_name>" {

    action = "${netscaler_responderaction.<resource_name>.name}"
    appflowaction = "${netscaler_appflowaction.<resource_name>.name}"
    comment = "abc"
    logaction = "abc"
    name = "abc"
    rule = "abc"
    undefaction = "NOOP"
}
```


# netscaler_appflowpolicy

Terraform resource name : ```netscaler_appflowpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[appflowaction.name](/doc/resources/appflowaction.md)|
|comment|No|No|string|
|name|No|No|string|
|rule|No|No|string|
|undefaction|No|No|[appflowaction.name](/doc/resources/appflowaction.md)|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowpolicy/appflowpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_appflowpolicy" "<resource_name>" {

    action = "${netscaler_appflowaction.<resource_name>.name}"
    comment = "abc"
    name = "abc"
    rule = "abc"
    undefaction = "${netscaler_appflowaction.<resource_name>.name}"
}
```


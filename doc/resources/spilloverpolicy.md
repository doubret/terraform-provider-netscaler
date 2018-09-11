# netscaler_spilloverpolicy

Terraform resource name : ```netscaler_spilloverpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[spilloveraction.name](/doc/resources/spilloveraction.md)|
|comment|No|No|string|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/spillover/spilloverpolicy/spilloverpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_spilloverpolicy" "<resource_name>" {

    action = "${netscaler_spilloveraction.<resource_name>.name}"
    comment = "abc"
    name = "abc"
    rule = "abc"
}
```


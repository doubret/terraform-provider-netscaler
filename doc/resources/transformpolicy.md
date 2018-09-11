# netscaler_transformpolicy

Terraform resource name : ```netscaler_transformpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|profilename|No|No|[transformprofile.name](/doc/resources/transformprofile.md)|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/transform/transformpolicy/transformpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_transformpolicy" "<resource_name>" {

    comment = "abc"
    logaction = "abc"
    name = "abc"
    profilename = "${netscaler_transformprofile.<resource_name>.name}"
    rule = "abc"
}
```


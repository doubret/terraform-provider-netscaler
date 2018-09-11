# netscaler_appqoepolicy

Terraform resource name : ```netscaler_appqoepolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[appqoeaction.name](/doc/resources/appqoeaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appqoe/appqoepolicy/appqoepolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appqoepolicy" "<resource_name>" {

    action = "${netscaler_appqoeaction.<resource_name>.name}"
    name = "abc"
    rule = "abc"
}
```


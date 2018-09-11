# netscaler_appfwpolicy

Terraform resource name : ```netscaler_appfwpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|profilename|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/application-firewall/appfwpolicy/appfwpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appfwpolicy" "<resource_name>" {

    comment = "abc"
    logaction = "abc"
    name = "abc"
    profilename = "abc"
    rule = "abc"
}
```


# netscaler_transformaction

Terraform resource name : ```netscaler_transformaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|cookiedomainfrom|No|No|string|
|cookiedomaininto|No|No|string|
|name|No|No|string|
|priority|No|No|double|
|profilename|No|No|[transformprofile.name](/doc/resources/transformprofile.md)|
|requrlfrom|No|No|string|
|requrlinto|No|No|string|
|resurlfrom|No|No|string|
|resurlinto|No|No|string|
|state|No|No|ENABLED, DISABLED|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/transform/transformaction/transformaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_transformaction" "<resource_name>" {

    comment = "abc"
    cookiedomainfrom = "abc"
    cookiedomaininto = "abc"
    name = "abc"
    priority = 42
    profilename = "${netscaler_transformprofile.<resource_name>.name}"
    requrlfrom = "abc"
    requrlinto = "abc"
    resurlfrom = "abc"
    resurlinto = "abc"
    state = "ENABLED"
}
```


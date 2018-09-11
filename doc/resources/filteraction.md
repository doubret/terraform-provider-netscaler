# netscaler_filteraction

Terraform resource name : ```netscaler_filteraction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|string|
|page|No|No|string|
|qual|No|No|reset, add, corrupt, forward, errorcode, drop|
|respcode|No|No|double|
|servicename|No|No|[service.name](/doc/resources/service.md)|
|value|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/filter/filteraction/filteraction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_filteraction" "tf_name" {

    name = "abc"
    page = "abc"
    qual = "reset"
    respcode = 42
    servicename = "${netscaler_service.<resource_name>.name}"
    value = "abc"
}
```


# netscaler_policydataset_value_binding

Terraform resource name : ```netscaler_policydataset_value_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|index|No|No|double|
|name|No|No|[policydataset.name](/doc/resources/policydataset.md)|
|value|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policydataset_value_binding/policydataset_value_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_policydataset_value_binding" "<resource_name>" {

    index = 42
    name = "${netscaler_policydataset.<resource_name>.name}"
    value = "abc"
}
```


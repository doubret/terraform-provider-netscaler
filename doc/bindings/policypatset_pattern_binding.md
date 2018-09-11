# netscaler_policypatset_pattern_binding

Terraform resource name : ```netscaler_policypatset_pattern_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|charset|No|No|ASCII, UTF_8|
|index|No|No|double|
|name|No|No|[policypatset.name](/doc/resources/policypatset.md)|
|string|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policypatset_pattern_binding/policypatset_pattern_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_policypatset_pattern_binding" "<resource_name>" {

    charset = "ASCII"
    index = 42
    name = "${netscaler_policypatset.<resource_name>.name}"
    string = "abc"
}
```


# netscaler_policypatset

Terraform resource name : ```netscaler_policypatset```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|indextype|No|No|Auto-generated, User-defined|
|name|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policypatset/policypatset/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_policypatset" "<resource_name>" {

    comment = "abc"
    indextype = "Auto-generated"
    name = "abc"
}
```


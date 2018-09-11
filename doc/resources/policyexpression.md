# netscaler_policyexpression

Terraform resource name : ```netscaler_policyexpression```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|clientsecuritymessage|No|No|string|
|comment|No|No|string|
|description|No|No|string|
|name|No|No|string|
|value|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policyexpression/policyexpression/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_policyexpression" "<resource_name>" {

    clientsecuritymessage = "abc"
    comment = "abc"
    description = "abc"
    name = "abc"
    value = "abc"
}
```


# netscaler_caaction

Terraform resource name : ```netscaler_caaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|accumressize|No|No|double|
|comment|No|No|string|
|lbvserver|No|No|string|
|name|No|No|string|
|type|No|No|nolookup, lookup, noop|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ca/caaction/caaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_caaction" "<resource_name>" {

    accumressize = 42
    comment = "abc"
    lbvserver = "abc"
    name = "abc"
    type = "nolookup"
}
```


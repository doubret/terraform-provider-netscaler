# netscaler_policydataset

Terraform resource name : ```netscaler_policydataset```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|indextype|No|No|Auto-generated, User-defined|
|name|No|No|string|
|type|No|No|ipv4, number, ipv6, ulong, double, mac|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policydataset/policydataset/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_policydataset" "tf_name" {

    comment = "abc"
    indextype = "Auto-generated"
    name = "abc"
    type = "ipv4"
}
```


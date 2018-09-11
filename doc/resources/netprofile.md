# netscaler_netprofile

Terraform resource name : ```netscaler_netprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|string|
|overridelsn|No|No|ENABLED, DISABLED|
|srcip|No|No|ip|
|srcippersistency|No|No|ENABLED, DISABLED|
|td|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/network/netprofile/netprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_netprofile" "<resource_name>" {

    name = "abc"
    overridelsn = "ENABLED"
    srcip = "1.2.3.4"
    srcippersistency = "ENABLED"
    td = 42
}
```


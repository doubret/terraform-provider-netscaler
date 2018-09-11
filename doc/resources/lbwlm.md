# netscaler_lbwlm

Terraform resource name : ```netscaler_lbwlm```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ipaddress|No|No|ip|
|katimeout|No|No|double|
|lbuid|No|No|string|
|port|No|No|int|
|wlmname|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbwlm/lbwlm/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_lbwlm" "<resource_name>" {

    ipaddress = "1.2.3.4"
    katimeout = 42
    lbuid = "abc"
    port = 42
    wlmname = "abc"
}
```


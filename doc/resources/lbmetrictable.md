# netscaler_lbmetrictable

Terraform resource name : ```netscaler_lbmetrictable```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|metrictable|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmetrictable/lbmetrictable/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbmetrictable" "<resource_name>" {

    metrictable = "abc"
}
```


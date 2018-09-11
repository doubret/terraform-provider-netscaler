# netscaler_cmpaction

Terraform resource name : ```netscaler_cmpaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|addvaryheader|No|No|GLOBAL, DISABLED, ENABLED|
|cmptype|No|No|compress, gzip, deflate, nocompress|
|deltatype|No|No|PERURL, PERPOLICY|
|name|No|No|string|
|varyheadervalue|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/compression/cmpaction/cmpaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cmpaction" "<resource_name>" {

    addvaryheader = "GLOBAL"
    cmptype = "compress"
    deltatype = "PERURL"
    name = "abc"
    varyheadervalue = "abc"
}
```


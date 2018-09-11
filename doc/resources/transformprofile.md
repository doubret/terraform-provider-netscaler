# netscaler_transformprofile

Terraform resource name : ```netscaler_transformprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|name|No|No|string|
|onlytransformabsurlinbody|No|No|ON, OFF|
|type|No|No|URL|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/transform/transformprofile/transformprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_transformprofile" "<resource_name>" {

    comment = "abc"
    name = "abc"
    onlytransformabsurlinbody = "ON"
    type = "URL"
}
```


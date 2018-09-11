# netscaler_dospolicy

Terraform resource name : ```netscaler_dospolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|cltdetectrate|No|No|double|
|name|No|No|string|
|qdepth|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/http-dos-protection/dospolicy/dospolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_dospolicy" "<resource_name>" {

    cltdetectrate = 42
    name = "abc"
    qdepth = 42
}
```


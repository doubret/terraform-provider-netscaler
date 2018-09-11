# netscaler_cmppolicy

Terraform resource name : ```netscaler_cmppolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|string|
|resaction|No|No|[cmpaction.name](/doc/resources/cmpaction.md)|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/compression/cmppolicy/cmppolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cmppolicy" "tf_name" {

    name = "abc"
    resaction = 
    rule = "abc"
}
```


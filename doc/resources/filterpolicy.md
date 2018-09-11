# netscaler_filterpolicy

Terraform resource name : ```netscaler_filterpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|string|
|reqaction|No|No|[filteraction.name](/doc/resources/filteraction.md)|
|resaction|No|No|[filteraction.name](/doc/resources/filteraction.md)|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/filter/filterpolicy/filterpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_filterpolicy" "<resource_name>" {

    name = "abc"
    reqaction = "${netscaler_filteraction.<resource_name>.name}"
    resaction = "${netscaler_filteraction.<resource_name>.name}"
    rule = "abc"
}
```


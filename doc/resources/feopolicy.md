# netscaler_feopolicy

Terraform resource name : ```netscaler_feopolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[feoaction.name](/doc/resources/feoaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/front-end-optimization/feopolicy/feopolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_feopolicy" "tf_name" {

    action = "${netscaler_feoaction.<resource_name>.name}"
    name = "abc"
    rule = "abc"
}
```


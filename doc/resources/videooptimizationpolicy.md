# netscaler_videooptimizationpolicy

Terraform resource name : ```netscaler_videooptimizationpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[videooptimizationaction.name](/doc/resources/videooptimizationaction.md)|
|comment|No|No|string|
|logaction|No|No|string|
|name|No|No|string|
|rule|No|No|string|
|undefaction|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/videooptimization/videooptimizationpolicy/videooptimizationpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_videooptimizationpolicy" "tf_name" {

    action = 
    comment = "abc"
    logaction = "abc"
    name = "abc"
    rule = "abc"
    undefaction = "abc"
}
```


# netscaler_spilloveraction

Terraform resource name : ```netscaler_spilloveraction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|SPILLOVER|
|name|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/spillover/spilloveraction/spilloveraction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_spilloveraction" "tf_name" {

    action = "SPILLOVER"
    name = "abc"
}
```


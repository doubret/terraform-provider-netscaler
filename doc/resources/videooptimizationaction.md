# netscaler_videooptimizationaction

Terraform resource name : ```netscaler_videooptimizationaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|name|No|No|string|
|rate|No|No|int|
|type|No|No|clear_text_pd, clear_text_abr, encrypted_abr, trigger_enc_abr, optimize_abr|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/videooptimization/videooptimizationaction/videooptimizationaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_videooptimizationaction" "tf_name" {

    comment = "abc"
    name = "abc"
    rate = 42
    type = "clear_text_pd"
}
```


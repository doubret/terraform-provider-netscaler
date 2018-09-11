# netscaler_videooptimizationpolicylabel

Terraform resource name : ```netscaler_videooptimizationpolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|labelname|No|No|string|
|policylabeltype|No|No|videoopt_req, videoopt_res|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/videooptimization/videooptimizationpolicylabel/videooptimizationpolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_videooptimizationpolicylabel" "<resource_name>" {

    comment = "abc"
    labelname = "abc"
    policylabeltype = "videoopt_req"
}
```


# netscaler_appflowpolicylabel

Terraform resource name : ```netscaler_appflowpolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|labelname|No|No|string|
|policylabeltype|No|No|HTTP, OTHERTCP|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowpolicylabel/appflowpolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appflowpolicylabel" "<resource_name>" {

    labelname = "abc"
    policylabeltype = "HTTP"
}
```


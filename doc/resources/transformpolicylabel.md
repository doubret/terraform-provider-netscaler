# netscaler_transformpolicylabel

Terraform resource name : ```netscaler_transformpolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|labelname|No|No|string|
|policylabeltype|No|No|http_req|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/transform/transformpolicylabel/transformpolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_transformpolicylabel" "<resource_name>" {

    labelname = "abc"
    policylabeltype = "http_req"
}
```


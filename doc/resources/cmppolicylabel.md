# netscaler_cmppolicylabel

Terraform resource name : ```netscaler_cmppolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|labelname|No|No|string|
|type|No|No|REQ, RES|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/compression/cmppolicylabel/cmppolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cmppolicylabel" "<resource_name>" {

    labelname = "abc"
    type = "REQ"
}
```


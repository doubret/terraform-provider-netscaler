# netscaler_servicegroup_servicegroupmember_binding

Terraform resource name : ```netscaler_servicegroup_servicegroupmember_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|port|No|No|int|
|servername|No|No|string|
|servicegroupname|No|No|string|
|weight|No|No|double|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/servicegroup_servicegroupmember_binding/servicegroup_servicegroupmember_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_servicegroup_servicegroupmember_binding" "<resource_name>" {

    port = 42
    servername = "abc"
    servicegroupname = "abc"
    weight = 42
}
```


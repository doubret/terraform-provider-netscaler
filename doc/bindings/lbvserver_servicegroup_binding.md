# netscaler_lbvserver_servicegroup_binding

Terraform resource name : ```netscaler_lbvserver_servicegroup_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|servicegroupname|No|No|[servicegroup.servicegroupname](/doc/resources/servicegroup.md)|
|weight|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_servicegroup_binding/lbvserver_servicegroup_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbvserver_servicegroup_binding" "<resource_name>" {

    name = "${netscaler_lbvserver.<resource_name>.name}"
    servicegroupname = "${netscaler_servicegroup.<resource_name>.servicegroupname}"
    weight = 42
}
```


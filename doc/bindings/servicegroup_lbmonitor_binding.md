# netscaler_servicegroup_lbmonitor_binding

Terraform resource name : ```netscaler_servicegroup_lbmonitor_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|monitor_name|No|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|servicegroupname|No|No|[servicegroup.servicegroupname](/doc/resources/servicegroup.md)|
|weight|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/servicegroup_lbmonitor_binding/servicegroup_lbmonitor_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_servicegroup_lbmonitor_binding" "<resource_name>" {

    monitor_name = "${netscaler_lbmonitor.<resource_name>.monitorname}"
    servicegroupname = "${netscaler_servicegroup.<resource_name>.servicegroupname}"
    weight = 42
}
```


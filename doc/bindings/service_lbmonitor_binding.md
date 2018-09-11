# netscaler_service_lbmonitor_binding

Terraform resource name : ```netscaler_service_lbmonitor_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|monitor_name|No|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|monstate|No|No|string|
|name|No|No|[service.name](/doc/resources/service.md)|
|passive|No|No|bool|
|weight|No|No|double|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service_lbmonitor_binding/service_lbmonitor_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_service_lbmonitor_binding" "<resource_name>" {

    monitor_name = "${netscaler_lbmonitor.<resource_name>.monitorname}"
    monstate = "abc"
    name = "${netscaler_service.<resource_name>.name}"
    passive = true
    weight = 42
}
```


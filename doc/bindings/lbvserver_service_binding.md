# netscaler_lbvserver_service_binding

Terraform resource name : ```netscaler_lbvserver_service_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|servicename|No|No|[service.name](/doc/resources/service.md)|
|weight|No|No|double|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_service_binding/lbvserver_service_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbvserver_service_binding" "<resource_name>" {

    name = "${netscaler_lbvserver.<resource_name>.name}"
    servicename = "${netscaler_service.<resource_name>.name}"
    weight = 42
}
```


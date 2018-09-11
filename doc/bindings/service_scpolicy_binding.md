# netscaler_service_scpolicy_binding

Terraform resource name : ```netscaler_service_scpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|name|No|No|[service.name](/doc/resources/service.md)|
|policyname|No|No|[scpolicy.name](/doc/resources/scpolicy.md)|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service_scpolicy_binding/service_scpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_service_scpolicy_binding" "<resource_name>" {

    name = "${netscaler_service.<resource_name>.name}"
    policyname = "${netscaler_scpolicy.<resource_name>.name}"
}
```


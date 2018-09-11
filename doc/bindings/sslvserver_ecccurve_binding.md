# netscaler_sslvserver_ecccurve_binding

Terraform resource name : ```netscaler_sslvserver_ecccurve_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ecccurvename|No|No|ALL, P_224, P_256, P_384, P_521|
|vservername|No|No|string|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ssl/sslvserver_ecccurve_binding/sslvserver_ecccurve_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_sslvserver_ecccurve_binding" "<resource_name>" {

    ecccurvename = "ALL"
    vservername = "abc"
}
```


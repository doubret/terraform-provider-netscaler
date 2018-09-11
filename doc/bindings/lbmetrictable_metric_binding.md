# netscaler_lbmetrictable_metric_binding

Terraform resource name : ```netscaler_lbmetrictable_metric_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|metric|No|No|string|
|metrictable|No|No|[lbmetrictable.metrictable](/doc/resources/lbmetrictable.md)|
|snmpoid|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmetrictable_metric_binding/lbmetrictable_metric_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbmetrictable_metric_binding" "<resource_name>" {

    metric = "abc"
    metrictable = "${netscaler_lbmetrictable.<resource_name>.metrictable}"
    snmpoid = "abc"
}
```


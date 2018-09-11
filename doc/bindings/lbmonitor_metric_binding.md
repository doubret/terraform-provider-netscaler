# netscaler_lbmonitor_metric_binding

Terraform resource name : ```netscaler_lbmonitor_metric_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|metric|No|No|string|
|metricthreshold|No|No|double|
|metricweight|No|No|double|
|monitorname|No|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmonitor_metric_binding/lbmonitor_metric_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbmonitor_metric_binding" "<resource_name>" {

    metric = "abc"
    metricthreshold = 42
    metricweight = 42
    monitorname = "${netscaler_lbmonitor.<resource_name>.monitorname}"
}
```


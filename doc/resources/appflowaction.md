# netscaler_appflowaction

Terraform resource name : ```netscaler_appflowaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|clientsidemeasurements|No|No|ENABLED, DISABLED|
|collectors|Yes|No|[appflowcollector.name](/doc/resources/appflowcollector.md)|
|comment|No|No|string|
|name|No|No|string|
|pagetracking|No|No|ENABLED, DISABLED|
|securityinsight|No|No|ENABLED, DISABLED|
|webinsight|No|No|ENABLED, DISABLED|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowaction/appflowaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appflowaction" "<resource_name>" {

    clientsidemeasurements = "ENABLED"
    collectors = [ ... ]"${netscaler_appflowcollector.<resource_name>.name[]}"
    comment = "abc"
    name = "abc"
    pagetracking = "ENABLED"
    securityinsight = "ENABLED"
    webinsight = "ENABLED"
}
```


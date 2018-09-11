# netscaler_dnsprofile

Terraform resource name : ```netscaler_dnsprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|cacheecsresponses|No|No|ENABLED, DISABLED|
|cachenegativeresponses|No|No|ENABLED, DISABLED|
|cacherecords|No|No|ENABLED, DISABLED|
|dnsanswerseclogging|No|No|ENABLED, DISABLED|
|dnserrorlogging|No|No|ENABLED, DISABLED|
|dnsextendedlogging|No|No|ENABLED, DISABLED|
|dnsprofilename|No|No|string|
|dnsquerylogging|No|No|ENABLED, DISABLED|
|dropmultiqueryrequest|No|No|ENABLED, DISABLED|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/domain-name-service/dnsprofile/dnsprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_dnsprofile" "<resource_name>" {

    cacheecsresponses = "ENABLED"
    cachenegativeresponses = "ENABLED"
    cacherecords = "ENABLED"
    dnsanswerseclogging = "ENABLED"
    dnserrorlogging = "ENABLED"
    dnsextendedlogging = "ENABLED"
    dnsprofilename = "abc"
    dnsquerylogging = "ENABLED"
    dropmultiqueryrequest = "ENABLED"
}
```


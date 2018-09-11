# netscaler_responderpolicylabel

Terraform resource name : ```netscaler_responderpolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|labelname|No|No|string|
|policylabeltype|No|No|HTTP, OTHERTCP, SIP_UDP, SIP_TCP, MYSQL, MSSQL, NAT, DIAMETER, RADIUS, DNS|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/responder/responderpolicylabel/responderpolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_responderpolicylabel" "<resource_name>" {

    comment = "abc"
    labelname = "abc"
    policylabeltype = "HTTP"
}
```


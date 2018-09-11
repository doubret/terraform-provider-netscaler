# netscaler_sslvserver_sslciphersuite_binding

Terraform resource name : ```netscaler_sslvserver_sslciphersuite_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ciphername|No|No|string|
|vservername|No|No|string|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ssl/sslvserver_sslciphersuite_binding/sslvserver_sslciphersuite_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_sslvserver_sslciphersuite_binding" "<resource_name>" {

    ciphername = "abc"
    vservername = "abc"
}
```


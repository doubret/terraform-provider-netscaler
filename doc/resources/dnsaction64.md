# netscaler_dnsaction64

Terraform resource name : ```netscaler_dnsaction64```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|actionname|No|No|string|
|excluderule|No|No|string|
|mappedrule|No|No|string|
|prefix|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/domain-name-service/dnsaction64/dnsaction64/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_dnsaction64" "<resource_name>" {

    actionname = "abc"
    excluderule = "abc"
    mappedrule = "abc"
    prefix = "abc"
}
```


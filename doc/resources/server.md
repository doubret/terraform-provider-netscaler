# netscaler_server

Terraform resource name : ```netscaler_server```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|domain|No|No|string|
|domainresolveretry|No|No|int|
|ipaddress|No|No|ip|
|ipv6address|No|No|YES, NO|
|name|No|No|string|
|state|No|No|ENABLED, DISABLED|
|td|No|No|double|
|translationip|No|No|ip|
|translationmask|No|No|ip_mask|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/server/server/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_server" "<resource_name>" {

    comment = "abc"
    domain = "abc"
    domainresolveretry = 42
    ipaddress = "1.2.3.4"
    ipv6address = "YES"
    name = "abc"
    state = "ENABLED"
    td = 42
    translationip = "1.2.3.4"
    translationmask = "255.255.255.0"
}
```


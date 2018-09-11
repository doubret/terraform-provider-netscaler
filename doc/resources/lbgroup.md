# netscaler_lbgroup

Terraform resource name : ```netscaler_lbgroup```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|backuppersistencetimeout|No|No|double|
|cookiedomain|No|No|string|
|cookiename|No|No|string|
|name|No|No|string|
|persistencebackup|No|No|SOURCEIP, NONE|
|persistencetype|No|No|SOURCEIP, COOKIEINSERT, RULE, NONE|
|persistmask|No|No|string|
|rule|No|No|string|
|timeout|No|No|double|
|usevserverpersistency|No|No|ENABLED, DISABLED|
|v6persistmasklen|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbgroup/lbgroup/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbgroup" "<resource_name>" {

    backuppersistencetimeout = 42
    cookiedomain = "abc"
    cookiename = "abc"
    name = "abc"
    persistencebackup = "SOURCEIP"
    persistencetype = "SOURCEIP"
    persistmask = "abc"
    rule = "abc"
    timeout = 42
    usevserverpersistency = "ENABLED"
    v6persistmasklen = 42
}
```


# netscaler_tmsessionaction

Terraform resource name : ```netscaler_tmsessionaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|defaultauthorizationaction|No|No|ALLOW, DENY|
|homepage|No|No|string|
|httponlycookie|No|No|YES, NO|
|kcdaccount|No|No|string|
|name|No|No|string|
|persistentcookie|No|No|YES, NO|
|persistentcookievalidity|No|No|double|
|sesstimeout|No|No|double|
|sso|No|No|ON, OFF|
|ssocredential|No|No|PRIMARY, SECONDARY|
|ssodomain|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/traffic-management/tmsessionaction/tmsessionaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_tmsessionaction" "<resource_name>" {

    defaultauthorizationaction = "ALLOW"
    homepage = "abc"
    httponlycookie = "YES"
    kcdaccount = "abc"
    name = "abc"
    persistentcookie = "YES"
    persistentcookievalidity = 42
    sesstimeout = 42
    sso = "ON"
    ssocredential = "PRIMARY"
    ssodomain = "abc"
}
```


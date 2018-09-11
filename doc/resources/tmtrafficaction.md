# netscaler_tmtrafficaction

Terraform resource name : ```netscaler_tmtrafficaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|apptimeout|No|No|double|
|forcedtimeout|No|No|START, STOP, RESET|
|forcedtimeoutval|No|No|double|
|formssoaction|No|No|string|
|initiatelogout|No|No|ON, OFF|
|kcdaccount|No|No|string|
|name|No|No|string|
|passwdexpression|No|No|string|
|persistentcookie|No|No|ON, OFF|
|samlssoprofile|No|No|string|
|sso|No|No|ON, OFF|
|userexpression|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/traffic-management/tmtrafficaction/tmtrafficaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_tmtrafficaction" "<resource_name>" {

    apptimeout = 42
    forcedtimeout = "START"
    forcedtimeoutval = 42
    formssoaction = "abc"
    initiatelogout = "ON"
    kcdaccount = "abc"
    name = "abc"
    passwdexpression = "abc"
    persistentcookie = "ON"
    samlssoprofile = "abc"
    sso = "ON"
    userexpression = "abc"
}
```


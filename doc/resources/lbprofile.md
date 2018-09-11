# netscaler_lbprofile

Terraform resource name : ```netscaler_lbprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|cookiepassphrase|No|No|string|
|dbslb|No|No|ENABLED, DISABLED|
|httponlycookieflag|No|No|ENABLED, DISABLED|
|lbprofilename|No|No|string|
|processlocal|No|No|ENABLED, DISABLED|
|useencryptedpersistencecookie|No|No|ENABLED, DISABLED|
|usesecuredpersistencecookie|No|No|ENABLED, DISABLED|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbprofile/lbprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbprofile" "<resource_name>" {

    cookiepassphrase = "abc"
    dbslb = "ENABLED"
    httponlycookieflag = "ENABLED"
    lbprofilename = "abc"
    processlocal = "ENABLED"
    useencryptedpersistencecookie = "ENABLED"
    usesecuredpersistencecookie = "ENABLED"
}
```


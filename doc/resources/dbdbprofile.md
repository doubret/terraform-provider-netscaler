# netscaler_dbdbprofile

Terraform resource name : ```netscaler_dbdbprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|conmultiplex|No|No|ENABLED, DISABLED|
|enablecachingconmuxoff|No|No|ENABLED, DISABLED|
|interpretquery|No|No|YES, NO|
|kcdaccount|No|No|string|
|name|No|No|string|
|stickiness|No|No|YES, NO|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/db/dbdbprofile/dbdbprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_dbdbprofile" "<resource_name>" {

    conmultiplex = "ENABLED"
    enablecachingconmuxoff = "ENABLED"
    interpretquery = "YES"
    kcdaccount = "abc"
    name = "abc"
    stickiness = "YES"
}
```


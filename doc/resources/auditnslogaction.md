# netscaler_auditnslogaction

Terraform resource name : ```netscaler_auditnslogaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|acl|No|No|ENABLED, DISABLED|
|alg|No|No|ENABLED, DISABLED|
|appflowexport|No|No|ENABLED, DISABLED|
|dateformat|No|No|MMDDYYYY, DDMMYYYY, YYYYMMDD|
|domainresolveretry|No|No|int|
|logfacility|No|No|LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7|
|loglevel|Yes|No|ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE|
|lsn|No|No|ENABLED, DISABLED|
|name|No|No|string|
|serverdomainname|No|No|string|
|serverip|No|No|ip|
|serverport|No|No|int|
|sslinterception|No|No|ENABLED, DISABLED|
|subscriberlog|No|No|ENABLED, DISABLED|
|tcp|No|No|NONE, ALL|
|timezone|No|No|GMT_TIME, LOCAL_TIME|
|userdefinedauditlog|No|No|YES, NO|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/audit/auditnslogaction/auditnslogaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_auditnslogaction" "<resource_name>" {

    acl = "ENABLED"
    alg = "ENABLED"
    appflowexport = "ENABLED"
    dateformat = "MMDDYYYY"
    domainresolveretry = 42
    logfacility = "LOCAL0"
    loglevel = [ ... ]
    lsn = "ENABLED"
    name = "abc"
    serverdomainname = "abc"
    serverip = "1.2.3.4"
    serverport = 42
    sslinterception = "ENABLED"
    subscriberlog = "ENABLED"
    tcp = "NONE"
    timezone = "GMT_TIME"
    userdefinedauditlog = "YES"
}
```


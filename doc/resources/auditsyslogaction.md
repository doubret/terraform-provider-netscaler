# netscaler_auditsyslogaction

Terraform resource name : ```netscaler_auditsyslogaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|acl|No|No|ENABLED, DISABLED|
|alg|No|No|ENABLED, DISABLED|
|appflowexport|No|No|ENABLED, DISABLED|
|dateformat|No|No|MMDDYYYY, DDMMYYYY, YYYYMMDD|
|dns|No|No|ENABLED, DISABLED|
|domainresolveretry|No|No|int|
|lbvservername|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|logfacility|No|No|LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7|
|loglevel|Yes|No|ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE|
|lsn|No|No|ENABLED, DISABLED|
|maxlogdatasizetohold|No|No|double|
|name|No|No|string|
|netprofile|No|No|[netprofile.name](/doc/resources/netprofile.md)|
|serverdomainname|No|No|string|
|serverip|No|No|ip|
|serverport|No|No|int|
|sslinterception|No|No|ENABLED, DISABLED|
|subscriberlog|No|No|ENABLED, DISABLED|
|tcp|No|No|NONE, ALL|
|tcpprofilename|No|No|[nstcpprofile.name](/doc/resources/nstcpprofile.md)|
|timezone|No|No|GMT_TIME, LOCAL_TIME|
|transport|No|No|TCP, UDP|
|userdefinedauditlog|No|No|YES, NO|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/audit/auditsyslogaction/auditsyslogaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_auditsyslogaction" "<resource_name>" {

    acl = "ENABLED"
    alg = "ENABLED"
    appflowexport = "ENABLED"
    dateformat = "MMDDYYYY"
    dns = "ENABLED"
    domainresolveretry = 42
    lbvservername = "${netscaler_lbvserver.<resource_name>.name}"
    logfacility = "LOCAL0"
    loglevel = [ ... ]
    lsn = "ENABLED"
    maxlogdatasizetohold = 42
    name = "abc"
    netprofile = "${netscaler_netprofile.<resource_name>.name}"
    serverdomainname = "abc"
    serverip = "1.2.3.4"
    serverport = 42
    sslinterception = "ENABLED"
    subscriberlog = "ENABLED"
    tcp = "NONE"
    tcpprofilename = "${netscaler_nstcpprofile.<resource_name>.name}"
    timezone = "GMT_TIME"
    transport = "TCP"
    userdefinedauditlog = "YES"
}
```


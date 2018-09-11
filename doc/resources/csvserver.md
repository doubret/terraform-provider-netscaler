# netscaler_csvserver

Terraform resource name : ```netscaler_csvserver```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|appflowlog|No|No|ENABLED, DISABLED|
|authentication|No|No|ON, OFF|
|authenticationhost|No|No|string|
|authn401|No|No|ON, OFF|
|authnprofile|No|No|string|
|authnvsname|No|No|string|
|backupvserver|No|No|string|
|cacheable|No|No|YES, NO|
|casesensitive|No|No|ON, OFF|
|clttimeout|No|No|double|
|comment|No|No|string|
|dbprofilename|No|No|[dbdbprofile.name](/doc/resources/dbdbprofile.md)|
|disableprimaryondown|No|No|ENABLED, DISABLED|
|dnsprofilename|No|No|[dnsprofile.dnsprofilename](/doc/resources/dnsprofile.md)|
|downstateflush|No|No|ENABLED, DISABLED|
|httpprofilename|No|No|[nshttpprofile.name](/doc/resources/nshttpprofile.md)|
|icmpvsrresponse|No|No|PASSIVE, ACTIVE|
|insertvserveripport|No|No|OFF, VIPADDR, V6TOV4MAPPING|
|ipmask|No|No|ip_mask|
|ippattern|No|No|string|
|ipv46|No|No|ip|
|l2conn|No|No|ON, OFF|
|listenpolicy|No|No|string|
|listenpriority|No|No|double|
|mssqlserverversion|No|No|70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014|
|mysqlcharacterset|No|No|double|
|mysqlprotocolversion|No|No|double|
|mysqlservercapabilities|No|No|double|
|mysqlserverversion|No|No|string|
|name|No|No|string|
|netprofile|No|No|string|
|oracleserverversion|No|No|10G, 11G|
|port|No|No|int|
|precedence|No|No|RULE, URL|
|push|No|No|ENABLED, DISABLED|
|pushlabel|No|No|string|
|pushmulticlients|No|No|YES, NO|
|pushvserver|No|No|string|
|range|No|No|double|
|redirectportrewrite|No|No|ENABLED, DISABLED|
|redirecturl|No|No|string|
|rhistate|No|No|PASSIVE, ACTIVE|
|rtspnat|No|No|ON, OFF|
|servicetype|No|No|HTTP, SSL, TCP, FTP, RTSP, SSL_TCP, UDP, DNS, SIP_UDP, SIP_TCP, SIP_SSL, ANY, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, DNS_TCP, ORACLE, SMPP|
|sobackupaction|No|No|DROP, ACCEPT, REDIRECT|
|somethod|No|No|CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE|
|sopersistence|No|No|ENABLED, DISABLED|
|sopersistencetimeout|No|No|double|
|sothreshold|No|No|double|
|state|No|No|ENABLED, DISABLED|
|stateupdate|No|No|ENABLED, DISABLED|
|tcpprofilename|No|No|[nstcpprofile.name](/doc/resources/nstcpprofile.md)|
|td|No|No|double|
|vipheader|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/csvserver/csvserver/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_csvserver" "<resource_name>" {

    appflowlog = "ENABLED"
    authentication = "ON"
    authenticationhost = "abc"
    authn401 = "ON"
    authnprofile = "abc"
    authnvsname = "abc"
    backupvserver = "abc"
    cacheable = "YES"
    casesensitive = "ON"
    clttimeout = 42
    comment = "abc"
    dbprofilename = "${netscaler_dbdbprofile.<resource_name>.name}"
    disableprimaryondown = "ENABLED"
    dnsprofilename = "${netscaler_dnsprofile.<resource_name>.dnsprofilename}"
    downstateflush = "ENABLED"
    httpprofilename = "${netscaler_nshttpprofile.<resource_name>.name}"
    icmpvsrresponse = "PASSIVE"
    insertvserveripport = "OFF"
    ipmask = "255.255.255.0"
    ippattern = "abc"
    ipv46 = "1.2.3.4"
    l2conn = "ON"
    listenpolicy = "abc"
    listenpriority = 42
    mssqlserverversion = "70"
    mysqlcharacterset = 42
    mysqlprotocolversion = 42
    mysqlservercapabilities = 42
    mysqlserverversion = "abc"
    name = "abc"
    netprofile = "abc"
    oracleserverversion = "10G"
    port = 42
    precedence = "RULE"
    push = "ENABLED"
    pushlabel = "abc"
    pushmulticlients = "YES"
    pushvserver = "abc"
    range = 42
    redirectportrewrite = "ENABLED"
    redirecturl = "abc"
    rhistate = "PASSIVE"
    rtspnat = "ON"
    servicetype = "HTTP"
    sobackupaction = "DROP"
    somethod = "CONNECTION"
    sopersistence = "ENABLED"
    sopersistencetimeout = 42
    sothreshold = 42
    state = "ENABLED"
    stateupdate = "ENABLED"
    tcpprofilename = "${netscaler_nstcpprofile.<resource_name>.name}"
    td = 42
    vipheader = "abc"
}
```


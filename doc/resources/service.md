# netscaler_service

Terraform resource name : ```netscaler_service```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|accessdown|No|No|YES, NO|
|appflowlog|No|No|ENABLED, DISABLED|
|cacheable|No|No|YES, NO|
|cachetype|No|No|TRANSPARENT, REVERSE, FORWARD|
|cip|No|No|ENABLED, DISABLED|
|cipheader|No|No|string|
|cka|No|No|YES, NO|
|cleartextport|No|No|int|
|clttimeout|No|No|int|
|cmp|No|No|YES, NO|
|comment|No|No|string|
|customserverid|No|No|string|
|dnsprofilename|No|No|[dnsprofile.dnsprofilename](/doc/resources/dnsprofile.md)|
|downstateflush|No|No|ENABLED, DISABLED|
|hashid|No|No|double|
|healthmonitor|No|No|YES, NO|
|httpprofilename|No|No|[nshttpprofile.name](/doc/resources/nshttpprofile.md)|
|ip|No|No|ip|
|maxbandwidth|No|No|double|
|maxclient|No|No|double|
|maxreq|No|No|double|
|monconnectionclose|No|No|RESET, FIN|
|monthreshold|No|No|double|
|name|No|No|string|
|netprofile|No|No|string|
|pathmonitor|No|No|YES, NO|
|pathmonitorindv|No|No|YES, NO|
|port|No|No|int|
|processlocal|No|No|ENABLED, DISABLED|
|rtspsessionidremap|No|No|ON, OFF|
|sc|No|No|ON, OFF|
|serverid|No|No|double|
|servername|No|No|[server.name](/doc/resources/server.md)|
|servicetype|No|No|HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP|
|sp|No|No|ON, OFF|
|svrtimeout|No|No|int|
|tcpb|No|No|YES, NO|
|tcpprofilename|No|No|[nstcpprofile.name](/doc/resources/nstcpprofile.md)|
|td|No|No|double|
|useproxyport|No|No|YES, NO|
|usip|No|No|YES, NO|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service/service/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_service" "<resource_name>" {

    accessdown = "YES"
    appflowlog = "ENABLED"
    cacheable = "YES"
    cachetype = "TRANSPARENT"
    cip = "ENABLED"
    cipheader = "abc"
    cka = "YES"
    cleartextport = 42
    clttimeout = 42
    cmp = "YES"
    comment = "abc"
    customserverid = "abc"
    dnsprofilename = "${netscaler_dnsprofile.<resource_name>.dnsprofilename}"
    downstateflush = "ENABLED"
    hashid = 42
    healthmonitor = "YES"
    httpprofilename = "${netscaler_nshttpprofile.<resource_name>.name}"
    ip = "1.2.3.4"
    maxbandwidth = 42
    maxclient = 42
    maxreq = 42
    monconnectionclose = "RESET"
    monthreshold = 42
    name = "abc"
    netprofile = "abc"
    pathmonitor = "YES"
    pathmonitorindv = "YES"
    port = 42
    processlocal = "ENABLED"
    rtspsessionidremap = "ON"
    sc = "ON"
    serverid = 42
    servername = "${netscaler_server.<resource_name>.name}"
    servicetype = "HTTP"
    sp = "ON"
    svrtimeout = 42
    tcpb = "YES"
    tcpprofilename = "${netscaler_nstcpprofile.<resource_name>.name}"
    td = 42
    useproxyport = "YES"
    usip = "YES"
}
```


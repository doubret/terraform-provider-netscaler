# netscaler_lbmonitor

Terraform resource name : ```netscaler_lbmonitor```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|NONE, LOG, DOWN|
|alertretries|No|No|int|
|application|No|No|string|
|attribute|No|No|string|
|basedn|No|No|string|
|binddn|No|No|string|
|customheaders|No|No|string|
|database|No|No|string|
|destip|No|No|ip|
|destport|No|No|int|
|deviation|No|No|double|
|dispatcherip|No|No|ip|
|dispatcherport|No|No|int|
|domain|No|No|string|
|downtime|No|No|int|
|evalrule|No|No|string|
|failureretries|No|No|int|
|filename|No|No|string|
|filter|No|No|string|
|firmwarerevision|No|No|double|
|group|No|No|string|
|hostipaddress|No|No|string|
|hostname|No|No|string|
|httprequest|No|No|string|
|inbandsecurityid|No|No|string|
|interval|No|No|int|
|iptunnel|No|No|string|
|kcdaccount|No|No|string|
|lasversion|No|No|string|
|logonpointname|No|No|string|
|lrtm|No|No|string|
|maxforwards|No|No|double|
|metrictable|No|No|string|
|monitorname|No|No|string|
|mssqlprotocolversion|No|No|string|
|netprofile|No|No|string|
|oraclesid|No|No|string|
|originhost|No|No|string|
|originrealm|No|No|string|
|password|No|No|string|
|productname|No|No|string|
|query|No|No|string|
|querytype|No|No|string|
|radaccountsession|No|No|string|
|radaccounttype|No|No|double|
|radapn|No|No|string|
|radframedip|No|No|ip|
|radkey|No|No|string|
|radmsisdn|No|No|string|
|radnasid|No|No|string|
|radnasip|No|No|ip|
|recv|No|No|string|
|resptimeout|No|No|int|
|resptimeoutthresh|No|No|double|
|retries|No|No|int|
|reverse|No|No|string|
|rtsprequest|No|No|string|
|scriptargs|No|No|string|
|scriptname|No|No|string|
|secondarypassword|No|No|string|
|secure|No|No|string|
|send|No|No|string|
|sipmethod|No|No|string|
|sipreguri|No|No|string|
|sipuri|No|No|string|
|sitepath|No|No|string|
|snmpcommunity|No|No|string|
|snmpoid|No|No|string|
|snmpthreshold|No|No|string|
|snmpversion|No|No|string|
|sqlquery|No|No|string|
|sslprofile|No|No|string|
|state|No|No|ENABLED, DISABLED|
|storedb|No|No|string|
|storefrontacctservice|No|No|string|
|storefrontcheckbackendservices|No|No|string|
|storename|No|No|string|
|successretries|No|No|int|
|tos|No|No|string|
|tosid|No|No|double|
|transparent|No|No|string|
|trofscode|No|No|double|
|trofsstring|No|No|string|
|type|No|No|PING, TCP, HTTP, TCP-ECV, HTTP-ECV, UDP-ECV, DNS, FTP, LDNS-PING, LDNS-TCP, LDNS-DNS, RADIUS, USER, HTTP-INLINE, SIP-UDP, SIP-TCP, LOAD, FTP-EXTENDED, SMTP, SNMP, NNTP, MYSQL, MYSQL-ECV, MSSQL-ECV, ORACLE-ECV, LDAP, POP3, CITRIX-XML-SERVICE, CITRIX-WEB-INTERFACE, DNS-TCP, RTSP, ARP, CITRIX-AG, CITRIX-AAC-LOGINPAGE, CITRIX-AAC-LAS, CITRIX-XD-DDC, ND6, CITRIX-WI-EXTENDED, DIAMETER, RADIUS_ACCOUNTING, STOREFRONT, APPC, SMPP, CITRIX-XNC-ECV, CITRIX-XDM, CITRIX-STA-SERVICE, CITRIX-STA-SERVICE-NHOP|
|units1|No|No|string|
|units2|No|No|string|
|units3|No|No|string|
|units4|No|No|string|
|username|No|No|string|
|validatecred|No|No|string|
|vendorid|No|No|double|
|vendorspecificvendorid|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmonitor/lbmonitor/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbmonitor" "<resource_name>" {

    action = "NONE"
    alertretries = 42
    application = "abc"
    attribute = "abc"
    basedn = "abc"
    binddn = "abc"
    customheaders = "abc"
    database = "abc"
    destip = "1.2.3.4"
    destport = 42
    deviation = 42
    dispatcherip = "1.2.3.4"
    dispatcherport = 42
    domain = "abc"
    downtime = 42
    evalrule = "abc"
    failureretries = 42
    filename = "abc"
    filter = "abc"
    firmwarerevision = 42
    group = "abc"
    hostipaddress = "abc"
    hostname = "abc"
    httprequest = "abc"
    inbandsecurityid = "abc"
    interval = 42
    iptunnel = "abc"
    kcdaccount = "abc"
    lasversion = "abc"
    logonpointname = "abc"
    lrtm = "abc"
    maxforwards = 42
    metrictable = "abc"
    monitorname = "abc"
    mssqlprotocolversion = "abc"
    netprofile = "abc"
    oraclesid = "abc"
    originhost = "abc"
    originrealm = "abc"
    password = "abc"
    productname = "abc"
    query = "abc"
    querytype = "abc"
    radaccountsession = "abc"
    radaccounttype = 42
    radapn = "abc"
    radframedip = "1.2.3.4"
    radkey = "abc"
    radmsisdn = "abc"
    radnasid = "abc"
    radnasip = "1.2.3.4"
    recv = "abc"
    resptimeout = 42
    resptimeoutthresh = 42
    retries = 42
    reverse = "abc"
    rtsprequest = "abc"
    scriptargs = "abc"
    scriptname = "abc"
    secondarypassword = "abc"
    secure = "abc"
    send = "abc"
    sipmethod = "abc"
    sipreguri = "abc"
    sipuri = "abc"
    sitepath = "abc"
    snmpcommunity = "abc"
    snmpoid = "abc"
    snmpthreshold = "abc"
    snmpversion = "abc"
    sqlquery = "abc"
    sslprofile = "abc"
    state = "ENABLED"
    storedb = "abc"
    storefrontacctservice = "abc"
    storefrontcheckbackendservices = "abc"
    storename = "abc"
    successretries = 42
    tos = "abc"
    tosid = 42
    transparent = "abc"
    trofscode = 42
    trofsstring = "abc"
    type = "PING"
    units1 = "abc"
    units2 = "abc"
    units3 = "abc"
    units4 = "abc"
    username = "abc"
    validatecred = "abc"
    vendorid = 42
    vendorspecificvendorid = 42
}
```


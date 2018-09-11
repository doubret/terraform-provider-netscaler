# netscaler_nstcpprofile

Terraform resource name : ```netscaler_nstcpprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ackaggregation|No|No|ENABLED, DISABLED|
|ackonpush|No|No|ENABLED, DISABLED|
|buffersize|No|No|double|
|burstratecontrol|No|No|DISABLED, FIXED, DYNAMIC|
|delayedack|No|No|double|
|dropestconnontimeout|No|No|ENABLED, DISABLED|
|drophalfclosedconnontimeout|No|No|ENABLED, DISABLED|
|dsack|No|No|ENABLED, DISABLED|
|dupackthresh|No|No|double|
|dynamicreceivebuffering|No|No|ENABLED, DISABLED|
|ecn|No|No|ENABLED, DISABLED|
|establishclientconn|No|No|AUTOMATIC, CONN_ESTABLISHED, ON_FIRST_DATA|
|fack|No|No|ENABLED, DISABLED|
|flavor|No|No|Default, Westwood, BIC, CUBIC, Nile|
|frto|No|No|ENABLED, DISABLED|
|hystart|No|No|ENABLED, DISABLED|
|initialcwnd|No|No|double|
|ka|No|No|ENABLED, DISABLED|
|kaconnidletime|No|No|double|
|kamaxprobes|No|No|double|
|kaprobeinterval|No|No|double|
|kaprobeupdatelastactivity|No|No|ENABLED, DISABLED|
|maxburst|No|No|double|
|maxcwnd|No|No|double|
|maxpktpermss|No|No|double|
|minrto|No|No|double|
|mptcp|No|No|ENABLED, DISABLED|
|mptcpdropdataonpreestsf|No|No|ENABLED, DISABLED|
|mptcpfastopen|No|No|ENABLED, DISABLED|
|mptcpsessiontimeout|No|No|double|
|mss|No|No|double|
|nagle|No|No|ENABLED, DISABLED|
|name|No|No|string|
|oooqsize|No|No|double|
|pktperretx|No|No|double|
|rateqmax|No|No|double|
|rstmaxack|No|No|ENABLED, DISABLED|
|rstwindowattenuate|No|No|ENABLED, DISABLED|
|sack|No|No|ENABLED, DISABLED|
|sendbuffsize|No|No|double|
|slowstartincr|No|No|double|
|spoofsyndrop|No|No|ENABLED, DISABLED|
|syncookie|No|No|ENABLED, DISABLED|
|tcpfastopen|No|No|ENABLED, DISABLED|
|tcpmode|No|No|TRANSPARENT, ENDPOINT|
|tcprate|No|No|double|
|tcpsegoffload|No|No|AUTOMATIC, DISABLED|
|timestamp|No|No|ENABLED, DISABLED|
|ws|No|No|ENABLED, DISABLED|
|wsval|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ns/nstcpprofile/nstcpprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_nstcpprofile" "tf_name" {

    ackaggregation = "ENABLED"
    ackonpush = "ENABLED"
    buffersize = 42
    burstratecontrol = "DISABLED"
    delayedack = 42
    dropestconnontimeout = "ENABLED"
    drophalfclosedconnontimeout = "ENABLED"
    dsack = "ENABLED"
    dupackthresh = 42
    dynamicreceivebuffering = "ENABLED"
    ecn = "ENABLED"
    establishclientconn = "AUTOMATIC"
    fack = "ENABLED"
    flavor = "Default"
    frto = "ENABLED"
    hystart = "ENABLED"
    initialcwnd = 42
    ka = "ENABLED"
    kaconnidletime = 42
    kamaxprobes = 42
    kaprobeinterval = 42
    kaprobeupdatelastactivity = "ENABLED"
    maxburst = 42
    maxcwnd = 42
    maxpktpermss = 42
    minrto = 42
    mptcp = "ENABLED"
    mptcpdropdataonpreestsf = "ENABLED"
    mptcpfastopen = "ENABLED"
    mptcpsessiontimeout = 42
    mss = 42
    nagle = "ENABLED"
    name = "abc"
    oooqsize = 42
    pktperretx = 42
    rateqmax = 42
    rstmaxack = "ENABLED"
    rstwindowattenuate = "ENABLED"
    sack = "ENABLED"
    sendbuffsize = 42
    slowstartincr = 42
    spoofsyndrop = "ENABLED"
    syncookie = "ENABLED"
    tcpfastopen = "ENABLED"
    tcpmode = "TRANSPARENT"
    tcprate = 42
    tcpsegoffload = "AUTOMATIC"
    timestamp = "ENABLED"
    ws = "ENABLED"
    wsval = 42
}
```


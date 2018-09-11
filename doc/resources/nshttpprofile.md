# netscaler_nshttpprofile

Terraform resource name : ```netscaler_nshttpprofile```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|adpttimeout|No|No|ENABLED, DISABLED|
|altsvc|No|No|ENABLED, DISABLED|
|apdexcltresptimethreshold|No|No|double|
|clientiphdrexpr|No|No|string|
|cmponpush|No|No|ENABLED, DISABLED|
|conmultiplex|No|No|ENABLED, DISABLED|
|dropextracrlf|No|No|ENABLED, DISABLED|
|dropextradata|No|No|ENABLED, DISABLED|
|dropinvalreqs|No|No|ENABLED, DISABLED|
|http2|No|No|ENABLED, DISABLED|
|http2direct|No|No|ENABLED, DISABLED|
|http2headertablesize|No|No|double|
|http2initialwindowsize|No|No|double|
|http2maxconcurrentstreams|No|No|double|
|http2maxframesize|No|No|double|
|http2maxheaderlistsize|No|No|double|
|http2minseverconn|No|No|double|
|incomphdrdelay|No|No|double|
|markconnreqinval|No|No|ENABLED, DISABLED|
|markhttp09inval|No|No|ENABLED, DISABLED|
|maxheaderlen|No|No|double|
|maxreq|No|No|double|
|maxreusepool|No|No|double|
|minreusepool|No|No|double|
|name|No|No|string|
|persistentetag|No|No|ENABLED, DISABLED|
|reqtimeout|No|No|double|
|reqtimeoutaction|No|No|string|
|reusepooltimeout|No|No|double|
|rtsptunnel|No|No|ENABLED, DISABLED|
|spdy|No|No|DISABLED, ENABLED, V2, V3|
|weblog|No|No|ENABLED, DISABLED|
|websocket|No|No|ENABLED, DISABLED|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ns/nshttpprofile/nshttpprofile/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_nshttpprofile" "<resource_name>" {

    adpttimeout = "ENABLED"
    altsvc = "ENABLED"
    apdexcltresptimethreshold = 42
    clientiphdrexpr = "abc"
    cmponpush = "ENABLED"
    conmultiplex = "ENABLED"
    dropextracrlf = "ENABLED"
    dropextradata = "ENABLED"
    dropinvalreqs = "ENABLED"
    http2 = "ENABLED"
    http2direct = "ENABLED"
    http2headertablesize = 42
    http2initialwindowsize = 42
    http2maxconcurrentstreams = 42
    http2maxframesize = 42
    http2maxheaderlistsize = 42
    http2minseverconn = 42
    incomphdrdelay = 42
    markconnreqinval = "ENABLED"
    markhttp09inval = "ENABLED"
    maxheaderlen = 42
    maxreq = 42
    maxreusepool = 42
    minreusepool = 42
    name = "abc"
    persistentetag = "ENABLED"
    reqtimeout = 42
    reqtimeoutaction = "abc"
    reusepooltimeout = 42
    rtsptunnel = "ENABLED"
    spdy = "DISABLED"
    weblog = "ENABLED"
    websocket = "ENABLED"
}
```


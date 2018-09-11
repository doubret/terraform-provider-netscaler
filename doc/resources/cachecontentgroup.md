# netscaler_cachecontentgroup

Terraform resource name : ```netscaler_cachecontentgroup```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|absexpiry|Yes|No|string|
|absexpirygmt|Yes|No|string|
|alwaysevalpolicies|No|No|YES, NO|
|cachecontrol|No|No|string|
|expireatlastbyte|No|No|YES, NO|
|flashcache|No|No|YES, NO|
|heurexpiryparam|No|No|double|
|hitparams|Yes|No|string|
|hitselector|No|No|string|
|ignoreparamvaluecase|No|No|YES, NO|
|ignorereloadreq|No|No|YES, NO|
|ignorereqcachinghdrs|No|No|YES, NO|
|insertage|No|No|YES, NO|
|insertetag|No|No|YES, NO|
|insertvia|No|No|YES, NO|
|invalparams|Yes|No|string|
|invalrestrictedtohost|No|No|YES, NO|
|invalselector|No|No|string|
|lazydnsresolve|No|No|YES, NO|
|matchcookies|No|No|YES, NO|
|maxressize|No|No|double|
|memlimit|No|No|double|
|minhits|No|No|int|
|minressize|No|No|double|
|name|No|No|string|
|persistha|No|No|YES, NO|
|pinned|No|No|YES, NO|
|polleverytime|No|No|YES, NO|
|prefetch|No|No|YES, NO|
|prefetchmaxpending|No|No|double|
|prefetchperiod|No|No|double|
|prefetchperiodmillisec|No|No|double|
|quickabortsize|No|No|double|
|relexpiry|No|No|double|
|relexpirymillisec|No|No|double|
|removecookies|No|No|YES, NO|
|type|No|No|HTTP, MYSQL, MSSQL|
|weaknegrelexpiry|No|No|double|
|weakposrelexpiry|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/integrated-caching/cachecontentgroup/cachecontentgroup/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cachecontentgroup" "<resource_name>" {

    absexpiry = [ ... ]
    absexpirygmt = [ ... ]
    alwaysevalpolicies = "YES"
    cachecontrol = "abc"
    expireatlastbyte = "YES"
    flashcache = "YES"
    heurexpiryparam = 42
    hitparams = [ ... ]
    hitselector = "abc"
    ignoreparamvaluecase = "YES"
    ignorereloadreq = "YES"
    ignorereqcachinghdrs = "YES"
    insertage = "YES"
    insertetag = "YES"
    insertvia = "YES"
    invalparams = [ ... ]
    invalrestrictedtohost = "YES"
    invalselector = "abc"
    lazydnsresolve = "YES"
    matchcookies = "YES"
    maxressize = 42
    memlimit = 42
    minhits = 42
    minressize = 42
    name = "abc"
    persistha = "YES"
    pinned = "YES"
    polleverytime = "YES"
    prefetch = "YES"
    prefetchmaxpending = 42
    prefetchperiod = 42
    prefetchperiodmillisec = 42
    quickabortsize = 42
    relexpiry = 42
    relexpirymillisec = 42
    removecookies = "YES"
    type = "HTTP"
    weaknegrelexpiry = 42
    weakposrelexpiry = 42
}
```


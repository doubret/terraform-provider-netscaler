# netscaler_feoaction

Terraform resource name : ```netscaler_feoaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|cachemaxage|No|No|double|
|clientsidemeasurements|No|No|bool|
|convertimporttolink|No|No|bool|
|csscombine|No|No|bool|
|cssimginline|No|No|bool|
|cssinline|No|No|bool|
|cssminify|No|No|bool|
|cssmovetohead|No|No|bool|
|dnsshards|Yes|No|string|
|domainsharding|No|No|string|
|htmlminify|No|No|bool|
|imggiftopng|No|No|bool|
|imginline|No|No|bool|
|imglazyload|No|No|bool|
|imgshrinktoattrib|No|No|bool|
|imgtojpegxr|No|No|bool|
|imgtowebp|No|No|bool|
|jpgoptimize|No|No|bool|
|jsinline|No|No|bool|
|jsminify|No|No|bool|
|jsmovetoend|No|No|bool|
|name|No|No|string|
|pageextendcache|No|No|bool|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/front-end-optimization/feoaction/feoaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_feoaction" "tf_name" {

    cachemaxage = 42
    clientsidemeasurements = true
    convertimporttolink = true
    csscombine = true
    cssimginline = true
    cssinline = true
    cssminify = true
    cssmovetohead = true
    dnsshards = [ ... ]
    domainsharding = "abc"
    htmlminify = true
    imggiftopng = true
    imginline = true
    imglazyload = true
    imgshrinktoattrib = true
    imgtojpegxr = true
    imgtowebp = true
    jpgoptimize = true
    jsinline = true
    jsminify = true
    jsmovetoend = true
    name = "abc"
    pageextendcache = true
}
```


# netscaler_cachepolicy

Terraform resource name : ```netscaler_cachepolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|CACHE, NOCACHE, MAY_CACHE, MAY_NOCACHE, INVAL|
|invalgroups|Yes|No|[cachecontentgroup.name](/doc/resources/cachecontentgroup.md)|
|invalobjects|Yes|No|[cachecontentgroup.name](/doc/resources/cachecontentgroup.md)|
|policyname|No|No|string|
|rule|No|No|string|
|storeingroup|No|No|[cachecontentgroup.name](/doc/resources/cachecontentgroup.md)|
|undefaction|No|No|NOCACHE, RESET|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/integrated-caching/cachepolicy/cachepolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_cachepolicy" "<resource_name>" {

    action = "CACHE"
    invalgroups = [ ... ]"${netscaler_cachecontentgroup.<resource_name>.name[]}"
    invalobjects = [ ... ]"${netscaler_cachecontentgroup.<resource_name>.name[]}"
    policyname = "abc"
    rule = "abc"
    storeingroup = "${netscaler_cachecontentgroup.<resource_name>.name}"
    undefaction = "NOCACHE"
}
```


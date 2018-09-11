# netscaler_lbvserver_capolicy_binding

Terraform resource name : ```netscaler_lbvserver_capolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|bindpoint|No|No|REQUEST, RESPONSE|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|labelname|No|No|string|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|name|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|No|[capolicy.name](/doc/resources/capolicy.md)|
|priority|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_capolicy_binding/lbvserver_capolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbvserver_capolicy_binding" "<resource_name>" {

    bindpoint = "REQUEST"
    gotopriorityexpression = "abc"
    invoke = true
    labelname = "abc"
    labeltype = "reqvserver"
    name = "${netscaler_lbvserver.<resource_name>.name}"
    policyname = "${netscaler_capolicy.<resource_name>.name}"
    priority = 42
}
```


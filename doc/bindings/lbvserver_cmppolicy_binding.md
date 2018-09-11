# netscaler_lbvserver_cmppolicy_binding

Terraform resource name : ```netscaler_lbvserver_cmppolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|bindpoint|No|No|REQUEST, RESPONSE|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|labelname|No|No|string|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|name|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|No|[cmppolicy.name](/doc/resources/cmppolicy.md)|
|priority|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_cmppolicy_binding/lbvserver_cmppolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbvserver_cmppolicy_binding" "<resource_name>" {

    bindpoint = "REQUEST"
    gotopriorityexpression = "abc"
    invoke = true
    labelname = "abc"
    labeltype = "reqvserver"
    name = "${netscaler_lbvserver.<resource_name>.name}"
    policyname = "${netscaler_cmppolicy.<resource_name>.name}"
    priority = 42
}
```


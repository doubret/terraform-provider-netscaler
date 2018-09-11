# netscaler_cspolicylabel_cspolicy_binding

Terraform resource name : ```netscaler_cspolicylabel_cspolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|invoke_labelname|No|No|string|
|labelname|No|No|[cspolicylabel.labelname](/doc/resources/cspolicylabel.md)|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|policyname|No|No|[cspolicy.policyname](/doc/resources/cspolicy.md)|
|priority|No|No|double|
|targetvserver|No|No|string|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/cspolicylabel_cspolicy_binding/cspolicylabel_cspolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_cspolicylabel_cspolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    invoke_labelname = "abc"
    labelname = "${netscaler_cspolicylabel.<resource_name>.labelname}"
    labeltype = "reqvserver, resvserver, policylabel"
    policyname = "${netscaler_cspolicy.<resource_name>.policyname}"
    priority = 42
    targetvserver = "abc"
}
```


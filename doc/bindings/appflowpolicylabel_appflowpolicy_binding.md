# netscaler_appflowpolicylabel_appflowpolicy_binding

Terraform resource name : ```netscaler_appflowpolicylabel_appflowpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|invoke_labelname|No|No|string|
|labelname|No|No|[appflowpolicylabel.labelname](/doc/resources/appflowpolicylabel.md)|
|labeltype|No|No|vserver, policylabel|
|policyname|No|No|[appflowpolicy.name](/doc/resources/appflowpolicy.md)|
|priority|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowpolicylabel_appflowpolicy_binding/appflowpolicylabel_appflowpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appflowpolicylabel_appflowpolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    invoke_labelname = "abc"
    labelname = "${netscaler_appflowpolicylabel.<resource_name>.labelname}"
    labeltype = "vserver"
    policyname = "${netscaler_appflowpolicy.<resource_name>.name}"
    priority = 42
}
```


# netscaler_csvserver_appflowpolicy_binding

Terraform resource name : ```netscaler_csvserver_appflowpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|labelname|No|No|string|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|name|No|No|[csvserver.name](/doc/resources/csvserver.md)|
|policyname|No|No|[spilloverpolicy.name](/doc/resources/spilloverpolicy.md)|
|priority|No|No|double|
|targetlbvserver|No|No|[lbvserver.name](/doc/resources/lbvserver.md)|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/csvserver_appflowpolicy_binding/csvserver_appflowpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_csvserver_appflowpolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    labelname = "abc"
    labeltype = "reqvserver"
    name = "${netscaler_csvserver.<resource_name>.name}"
    policyname = "${netscaler_spilloverpolicy.<resource_name>.name}"
    priority = 42
    targetlbvserver = "${netscaler_lbvserver.<resource_name>.name}"
}
```


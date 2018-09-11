# netscaler_authorizationpolicylabel_authorizationpolicy_binding

Terraform resource name : ```netscaler_authorizationpolicylabel_authorizationpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|invoke_labelname|No|No|string|
|labelname|No|No|[authorizationpolicylabel.labelname](/doc/resources/authorizationpolicylabel.md)|
|labeltype|No|No|reqvserver, resvserver, policylabel|
|policyname|No|No|[authorizationpolicy.name](/doc/resources/authorizationpolicy.md)|
|priority|No|No|double|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/authorization/authorizationpolicylabel_authorizationpolicy_binding/authorizationpolicylabel_authorizationpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_authorizationpolicylabel_authorizationpolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    invoke_labelname = "abc"
    labelname = "${netscaler_authorizationpolicylabel.<resource_name>.labelname}"
    labeltype = "reqvserver, resvserver, policylabel"
    policyname = "${netscaler_authorizationpolicy.<resource_name>.name}"
    priority = 42
}
```


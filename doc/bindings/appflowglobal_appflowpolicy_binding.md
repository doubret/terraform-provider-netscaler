# netscaler_appflowglobal_appflowpolicy_binding

Terraform resource name : ```netscaler_appflowglobal_appflowpolicy_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|gotopriorityexpression|No|No|string|
|invoke|No|No|bool|
|labelname|No|No|[appflowpolicylabel.labelname](/doc/resources/appflowpolicylabel.md)|
|labeltype|No|No|vserver, policylabel|
|policyname|No|No|[appflowpolicy.name](/doc/resources/appflowpolicy.md)|
|priority|No|No|double|
|type|No|No|REQ_OVERRIDE, REQ_DEFAULT, OVERRIDE, DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, MSSQL_REQ_OVERRIDE, MSSQL_REQ_DEFAULT, MYSQL_REQ_OVERRIDE, MYSQL_REQ_DEFAULT, ICA_REQ_OVERRIDE, ICA_REQ_DEFAULT, ORACLE_REQ_OVERRIDE, ORACLE_REQ_DEFAULT|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowglobal_appflowpolicy_binding/appflowglobal_appflowpolicy_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_appflowglobal_appflowpolicy_binding" "<resource_name>" {

    gotopriorityexpression = "abc"
    invoke = true
    labelname = "${netscaler_appflowpolicylabel.<resource_name>.labelname}"
    labeltype = "vserver"
    policyname = "${netscaler_appflowpolicy.<resource_name>.name}"
    priority = 42
    type = "REQ_OVERRIDE"
}
```


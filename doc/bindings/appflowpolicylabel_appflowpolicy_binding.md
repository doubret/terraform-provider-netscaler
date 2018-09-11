# Binding appflowpolicylabel_appflowpolicy_binding

Spec for **appflowpolicylabel_appflowpolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowpolicylabel_appflowpolicy_binding/appflowpolicylabel_appflowpolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|invoke_labelname|No|string|
|labelname|No|[appflowpolicylabel.labelname](/doc/resources/appflowpolicylabel.md)|
|labeltype|No|vserver, policylabel|
|policyname|No|[appflowpolicy.name](/doc/resources/appflowpolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| labelname | appflowpolicylabel.labelname |
| policyname | appflowpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/appflowpolicylabel_appflowpolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/appflowpolicylabel_appflowpolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/appflowpolicylabel_appflowpolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/appflowpolicylabel_appflowpolicy_binding` |


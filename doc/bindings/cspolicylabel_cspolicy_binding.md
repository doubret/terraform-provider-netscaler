# Binding cspolicylabel_cspolicy_binding

Spec for **cspolicylabel_cspolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/cspolicylabel_cspolicy_binding/cspolicylabel_cspolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|invoke_labelname|No|string|
|labelname|No|[cspolicylabel.labelname](/doc/resources/cspolicylabel.md)|
|labeltype|No|reqvserver, resvserver, policylabel|
|policyname|No|[cspolicy.policyname](/doc/resources/cspolicy.md)|
|priority|No|double|
|targetvserver|No|string|

## Key

| Name | Type |
|----|----|
| labelname | cspolicylabel.labelname |
| policyname | cspolicy.policyname |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/cspolicylabel_cspolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/cspolicylabel_cspolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/cspolicylabel_cspolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/cspolicylabel_cspolicy_binding` |


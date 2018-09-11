# Binding authorizationpolicylabel_authorizationpolicy_binding

Spec for **authorizationpolicylabel_authorizationpolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/authorization/authorizationpolicylabel_authorizationpolicy_binding/authorizationpolicylabel_authorizationpolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|invoke_labelname|No|string|
|labelname|No|[authorizationpolicylabel.labelname](/doc/resources/authorizationpolicylabel.md)|
|labeltype|No|reqvserver, resvserver, policylabel|
|policyname|No|[authorizationpolicy.name](/doc/resources/authorizationpolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| labelname | authorizationpolicylabel.labelname |
| policyname | authorizationpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/authorizationpolicylabel_authorizationpolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/authorizationpolicylabel_authorizationpolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/authorizationpolicylabel_authorizationpolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/authorizationpolicylabel_authorizationpolicy_binding` |


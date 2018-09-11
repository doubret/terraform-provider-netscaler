# Binding lbvserver_appqoepolicy_binding

Spec for **lbvserver_appqoepolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_appqoepolicy_binding/lbvserver_appqoepolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|labelname|No|string|
|labeltype|No|reqvserver, resvserver, policylabel|
|name|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|[appqoepolicy.name](/doc/resources/appqoepolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| policyname | appqoepolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_appqoepolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_appqoepolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_appqoepolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_appqoepolicy_binding` |


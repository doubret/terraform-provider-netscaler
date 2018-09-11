# Binding csvserver_cachepolicy_binding

Spec for **csvserver_cachepolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/csvserver_cachepolicy_binding/csvserver_cachepolicy_binding/)

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
|name|No|[csvserver.name](/doc/resources/csvserver.md)|
|policyname|No|[spilloverpolicy.name](/doc/resources/spilloverpolicy.md)|
|priority|No|double|
|targetlbvserver|No|[lbvserver.name](/doc/resources/lbvserver.md)|

## Key

| Name | Type |
|----|----|
| name | csvserver.name |
| policyname | spilloverpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/csvserver_cachepolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/csvserver_cachepolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/csvserver_cachepolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/csvserver_cachepolicy_binding` |


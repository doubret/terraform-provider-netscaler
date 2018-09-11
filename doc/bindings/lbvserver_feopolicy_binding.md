# Binding lbvserver_feopolicy_binding

Spec for **lbvserver_feopolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_feopolicy_binding/lbvserver_feopolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|bindpoint|No|REQUEST, RESPONSE|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|labelname|No|string|
|labeltype|No|reqvserver, resvserver, policylabel|
|name|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|[feopolicy.name](/doc/resources/feopolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| policyname | feopolicy.name |
| bindpoint | REQUEST, RESPONSE |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_feopolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_feopolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_feopolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_feopolicy_binding` |


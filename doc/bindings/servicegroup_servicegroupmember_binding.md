# Binding servicegroup_servicegroupmember_binding

Spec for **servicegroup_servicegroupmember_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/servicegroup_servicegroupmember_binding/servicegroup_servicegroupmember_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|port|No|int|
|servername|No|string|
|servicegroupname|No|string|
|weight|No|double|

## Key

| Name | Type |
|----|----|
| servicegroupname | string |
| servername | string |
| port | int |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_servicegroupmember_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_servicegroupmember_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_servicegroupmember_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_servicegroupmember_binding` |


# Binding lbvserver_servicegroup_binding

Spec for **lbvserver_servicegroup_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_servicegroup_binding/lbvserver_servicegroup_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|name|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|servicegroupname|No|[servicegroup.servicegroupname](/doc/resources/servicegroup.md)|
|weight|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| servicegroupname | servicegroup.servicegroupname |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_servicegroup_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_servicegroup_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_servicegroup_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_servicegroup_binding` |


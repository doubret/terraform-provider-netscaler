# Binding servicegroup_lbmonitor_binding

Spec for **servicegroup_lbmonitor_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/servicegroup_lbmonitor_binding/servicegroup_lbmonitor_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|monitor_name|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|servicegroupname|No|[servicegroup.servicegroupname](/doc/resources/servicegroup.md)|
|weight|No|double|

## Key

| Name | Type |
|----|----|
| servicegroupname | servicegroup.servicegroupname |
| monitor_name | lbmonitor.monitorname |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_lbmonitor_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_lbmonitor_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_lbmonitor_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/servicegroup_lbmonitor_binding` |


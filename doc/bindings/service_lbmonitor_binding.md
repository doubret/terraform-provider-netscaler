# Binding service_lbmonitor_binding

Spec for **service_lbmonitor_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service_lbmonitor_binding/service_lbmonitor_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|monitor_name|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|monstate|No|string|
|name|No|[service.name](/doc/resources/service.md)|
|passive|No|bool|
|weight|No|double|

## Key

| Name | Type |
|----|----|
| name | service.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_lbmonitor_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_lbmonitor_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/service_lbmonitor_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/service_lbmonitor_binding` |


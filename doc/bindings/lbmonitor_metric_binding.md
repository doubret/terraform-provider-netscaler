# Binding lbmonitor_metric_binding

Spec for **lbmonitor_metric_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmonitor_metric_binding/lbmonitor_metric_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|metric|No|string|
|metricthreshold|No|double|
|metricweight|No|double|
|monitorname|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|

## Key

| Name | Type |
|----|----|
| monitorname | lbmonitor.monitorname |
| metric | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_metric_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_metric_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_metric_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_metric_binding` |


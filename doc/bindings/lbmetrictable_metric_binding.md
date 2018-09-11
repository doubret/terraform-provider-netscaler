# Binding lbmetrictable_metric_binding

Spec for **lbmetrictable_metric_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmetrictable_metric_binding/lbmetrictable_metric_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|metric|No|string|
|metrictable|No|[lbmetrictable.metrictable](/doc/resources/lbmetrictable.md)|
|snmpoid|No|string|

## Key

| Name | Type |
|----|----|
| metrictable | lbmetrictable.metrictable |
| metric | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmetrictable_metric_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmetrictable_metric_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbmetrictable_metric_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbmetrictable_metric_binding` |


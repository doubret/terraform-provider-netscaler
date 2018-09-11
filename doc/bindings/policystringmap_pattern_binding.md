# Binding policystringmap_pattern_binding

Spec for **policystringmap_pattern_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policystringmap_pattern_binding/policystringmap_pattern_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|key|No|string|
|name|No|[policystringmap.name](/doc/resources/policystringmap.md)|
|value|No|string|

## Key

| Name | Type |
|----|----|
| name | policystringmap.name |
| key | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/policystringmap_pattern_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/policystringmap_pattern_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/policystringmap_pattern_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/policystringmap_pattern_binding` |


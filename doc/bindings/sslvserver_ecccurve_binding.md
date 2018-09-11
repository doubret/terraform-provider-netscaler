# Binding sslvserver_ecccurve_binding

Spec for **sslvserver_ecccurve_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ssl/sslvserver_ecccurve_binding/sslvserver_ecccurve_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|ecccurvename|No|ALL, P_224, P_256, P_384, P_521|
|vservername|No|string|

## Key

| Name | Type |
|----|----|
| vservername | string |
| ecccurvename | ALL, P_224, P_256, P_384, P_521 |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_ecccurve_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_ecccurve_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_ecccurve_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_ecccurve_binding` |


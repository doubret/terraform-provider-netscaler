# Binding sslvserver_sslciphersuite_binding

Spec for **sslvserver_sslciphersuite_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ssl/sslvserver_sslciphersuite_binding/sslvserver_sslciphersuite_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|ciphername|No|string|
|vservername|No|string|

## Key

| Name | Type |
|----|----|
| vservername | string |
| ciphername | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslciphersuite_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslciphersuite_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslciphersuite_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslciphersuite_binding` |


# Binding sslvserver_sslcertkey_binding

Spec for **sslvserver_sslcertkey_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/ssl/sslvserver_sslcertkey_binding/sslvserver_sslcertkey_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|ca|No|bool|
|certkeyname|No|string|
|crlcheck|No|string|
|ocspcheck|No|string|
|skipcaname|No|bool|
|snicert|No|bool|
|vservername|No|string|

## Key

| Name | Type |
|----|----|
| vservername | string |
| certkeyname | string |
| ca | bool |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslcertkey_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslcertkey_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslcertkey_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/sslvserver_sslcertkey_binding` |


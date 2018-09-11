# Binding lbmonitor_sslcertkey_binding

Spec for **lbmonitor_sslcertkey_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmonitor_sslcertkey_binding/lbmonitor_sslcertkey_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|ca|No|bool|
|certkeyname|No|string|
|crlcheck|No|string|
|monitorname|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|ocspcheck|No|string|

## Key

| Name | Type |
|----|----|
| monitorname | lbmonitor.monitorname |
| certkeyname | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_sslcertkey_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_sslcertkey_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_sslcertkey_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbmonitor_sslcertkey_binding` |


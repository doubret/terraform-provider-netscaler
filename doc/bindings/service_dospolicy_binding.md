# Binding service_dospolicy_binding

Spec for **service_dospolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service_dospolicy_binding/service_dospolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|name|No|[service.name](/doc/resources/service.md)|
|policyname|No|[dospolicy.name](/doc/resources/dospolicy.md)|

## Key

| Name | Type |
|----|----|
| name | service.name |
| policyname | dospolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_dospolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_dospolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/service_dospolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/service_dospolicy_binding` |


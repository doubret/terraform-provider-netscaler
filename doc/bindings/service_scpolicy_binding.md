# Binding service_scpolicy_binding

Spec for **service_scpolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service_scpolicy_binding/service_scpolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|name|No|[service.name](/doc/resources/service.md)|
|policyname|No|[scpolicy.name](/doc/resources/scpolicy.md)|

## Key

| Name | Type |
|----|----|
| name | service.name |
| policyname | scpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_scpolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/service_scpolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/service_scpolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/service_scpolicy_binding` |


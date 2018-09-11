# Binding policypatset_pattern_binding

Spec for **policypatset_pattern_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policypatset_pattern_binding/policypatset_pattern_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|charset|No|ASCII, UTF_8|
|index|No|double|
|name|No|[policypatset.name](/doc/resources/policypatset.md)|
|string|No|string|

## Key

| Name | Type |
|----|----|
| name | policypatset.name |
| string | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/policypatset_pattern_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/policypatset_pattern_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/policypatset_pattern_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/policypatset_pattern_binding` |


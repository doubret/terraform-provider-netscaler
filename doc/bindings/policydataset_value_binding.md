# Binding policydataset_value_binding

Spec for **policydataset_value_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policydataset_value_binding/policydataset_value_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|index|No|double|
|name|No|[policydataset.name](/doc/resources/policydataset.md)|
|value|No|string|

## Key

| Name | Type |
|----|----|
| name | policydataset.name |
| value | string |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/policydataset_value_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/policydataset_value_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/policydataset_value_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/policydataset_value_binding` |


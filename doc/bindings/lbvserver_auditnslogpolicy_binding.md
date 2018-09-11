# Binding lbvserver_auditnslogpolicy_binding

Spec for **lbvserver_auditnslogpolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_auditnslogpolicy_binding/lbvserver_auditnslogpolicy_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|gotopriorityexpression|No|string|
|invoke|No|bool|
|labelname|No|string|
|labeltype|No|reqvserver, resvserver, policylabel|
|name|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|policyname|No|[auditnslogpolicy.name](/doc/resources/auditnslogpolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| policyname | auditnslogpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditnslogpolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditnslogpolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditnslogpolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditnslogpolicy_binding` |


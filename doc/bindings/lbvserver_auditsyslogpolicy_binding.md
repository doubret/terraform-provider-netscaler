# Binding lbvserver_auditsyslogpolicy_binding

Spec for **lbvserver_auditsyslogpolicy_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_auditsyslogpolicy_binding/lbvserver_auditsyslogpolicy_binding/)

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
|policyname|No|[auditsyslogpolicy.name](/doc/resources/auditsyslogpolicy.md)|
|priority|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| policyname | auditsyslogpolicy.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditsyslogpolicy_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditsyslogpolicy_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditsyslogpolicy_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_auditsyslogpolicy_binding` |


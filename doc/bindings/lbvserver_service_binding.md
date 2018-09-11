# Binding lbvserver_service_binding

Spec for **lbvserver_service_binding** binding - [citrix documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbvserver_service_binding/lbvserver_service_binding/)

- [Fields](#fields)
- [Key](#key)
- [Operations](#operations)

## Fields

| Name | Array | Type |
|----|----|
|name|No|[lbvserver.name](/doc/resources/lbvserver.md)|
|servicename|No|[service.name](/doc/resources/service.md)|
|weight|No|double|

## Key

| Name | Type |
|----|----|
| name | lbvserver.name |
| servicename | service.name |

## Operations

| Name | Method | Url |
|----|----|----|
| List | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_service_binding` |
| Get | GET | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_service_binding/<name>` |
| Delete | DELETE | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_service_binding/<name>` |
| Add | POST | `http://<netscaler-ip-address>/nitro/v1/config/lbvserver_service_binding` |


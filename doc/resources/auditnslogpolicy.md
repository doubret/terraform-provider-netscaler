# netscaler_auditnslogpolicy

Terraform resource name : ```netscaler_auditnslogpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[auditnslogaction.name](/doc/resources/auditnslogaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/audit/auditnslogpolicy/auditnslogpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_auditnslogpolicy" "tf_name" {

    action = 
    name = "abc"
    rule = "abc"
}
```


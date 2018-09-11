# netscaler_auditsyslogpolicy

Terraform resource name : ```netscaler_auditsyslogpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[auditsyslogaction.name](/doc/resources/auditsyslogaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/audit/auditsyslogpolicy/auditsyslogpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_auditsyslogpolicy" "tf_name" {

    action = "${netscaler_auditsyslogaction.<resource_name>.name}"
    name = "abc"
    rule = "abc"
}
```


# netscaler_tmtrafficpolicy

Terraform resource name : ```netscaler_tmtrafficpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[tmtrafficaction.name](/doc/resources/tmtrafficaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/traffic-management/tmtrafficpolicy/tmtrafficpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_tmtrafficpolicy" "<resource_name>" {

    action = "${netscaler_tmtrafficaction.<resource_name>.name}"
    name = "abc"
    rule = "abc"
}
```


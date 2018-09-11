# netscaler_policystringmap

Terraform resource name : ```netscaler_policystringmap```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|name|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/policy/policystringmap/policystringmap/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_policystringmap" "<resource_name>" {

    comment = "abc"
    name = "abc"
}
```


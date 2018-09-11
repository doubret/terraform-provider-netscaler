# netscaler_authorizationpolicy

Terraform resource name : ```netscaler_authorizationpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|string|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/authorization/authorizationpolicy/authorizationpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_authorizationpolicy" "<resource_name>" {

    action = "abc"
    name = "abc"
    rule = "abc"
}
```


# netscaler_authorizationpolicylabel

Terraform resource name : ```netscaler_authorizationpolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|labelname|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/authorization/authorizationpolicylabel/authorizationpolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_authorizationpolicylabel" "tf_name" {

    labelname = "abc"
}
```


# netscaler_responderaction

Terraform resource name : ```netscaler_responderaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|bypasssafetycheck|No|No|YES, NO|
|comment|No|No|string|
|htmlpage|No|No|string|
|name|No|No|string|
|reasonphrase|No|No|string|
|responsestatuscode|No|No|double|
|target|No|No|string|
|type|No|No|noop, respondwith, redirect, respondwithhtmlpage, sqlresponse_ok, sqlresponse_error|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/responder/responderaction/responderaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_responderaction" "tf_name" {

    bypasssafetycheck = "YES"
    comment = "abc"
    htmlpage = "abc"
    name = "abc"
    reasonphrase = "abc"
    responsestatuscode = 42
    target = "abc"
    type = "noop"
}
```


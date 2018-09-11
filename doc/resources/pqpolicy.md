# netscaler_pqpolicy

Terraform resource name : ```netscaler_pqpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|policyname|No|No|string|
|polqdepth|No|No|double|
|priority|No|No|double|
|qdepth|No|No|double|
|rule|No|No|string|
|weight|No|No|double|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/priority-queuing/pqpolicy/pqpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_pqpolicy" "tf_name" {

    policyname = "abc"
    polqdepth = 42
    priority = 42
    qdepth = 42
    rule = "abc"
    weight = 42
}
```


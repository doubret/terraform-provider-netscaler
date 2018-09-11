# netscaler_scpolicy

Terraform resource name : ```netscaler_scpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|string|
|altcontentpath|No|No|ACS, NS, NOACTION|
|altcontentsvcname|No|No|string|
|delay|No|No|double|
|maxconn|No|No|double|
|name|No|No|string|
|rule|No|No|string|
|url|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/sure-connect/scpolicy/scpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_scpolicy" "tf_name" {

    action = "abc"
    altcontentpath = "ACS"
    altcontentsvcname = "abc"
    delay = 42
    maxconn = 42
    name = "abc"
    rule = "abc"
    url = "abc"
}
```


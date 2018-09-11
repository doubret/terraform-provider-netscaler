# netscaler_appflowcollector

Terraform resource name : ```netscaler_appflowcollector```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ipaddress|No|No|ip|
|name|No|No|string|
|netprofile|No|No|[netprofile.name](/doc/resources/netprofile.md)|
|port|No|No|int|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appflow/appflowcollector/appflowcollector/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_appflowcollector" "tf_name" {

    ipaddress = "1.2.3.4"
    name = "abc"
    netprofile = 
    port = 42
}
```


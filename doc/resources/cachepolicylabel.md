# netscaler_cachepolicylabel

Terraform resource name : ```netscaler_cachepolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|evaluates|No|No|REQ, RES, MSSQL_REQ, MSSQL_RES, MYSQL_REQ, MYSQL_RES|
|labelname|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/integrated-caching/cachepolicylabel/cachepolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cachepolicylabel" "<resource_name>" {

    evaluates = "REQ"
    labelname = "abc"
}
```


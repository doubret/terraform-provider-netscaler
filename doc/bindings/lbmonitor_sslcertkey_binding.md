# netscaler_lbmonitor_sslcertkey_binding

Terraform resource name : ```netscaler_lbmonitor_sslcertkey_binding```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|ca|No|No|bool|
|certkeyname|No|No|string|
|crlcheck|No|No|string|
|monitorname|No|No|[lbmonitor.monitorname](/doc/resources/lbmonitor.md)|
|ocspcheck|No|No|string|


##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/load-balancing/lbmonitor_sslcertkey_binding/lbmonitor_sslcertkey_binding/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_lbmonitor_sslcertkey_binding" "<resource_name>" {

    ca = true
    certkeyname = "abc"
    crlcheck = "abc"
    monitorname = "${netscaler_lbmonitor.<resource_name>.monitorname}"
    ocspcheck = "abc"
}
```


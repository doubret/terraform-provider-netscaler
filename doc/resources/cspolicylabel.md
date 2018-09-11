# netscaler_cspolicylabel

Terraform resource name : ```netscaler_cspolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|cspolicylabeltype|No|No|HTTP, TCP, RTSP, SSL, SSL_TCP, UDP, DNS, SIP_UDP, SIP_TCP, ANY, RADIUS, RDP, MYSQL, MSSQL, ORACLE, DIAMETER, SSL_DIAMETER, FTP, DNS_TCP, SMPP|
|labelname|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/content-switching/cspolicylabel/cspolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_cspolicylabel" "tf_name" {

    cspolicylabeltype = "HTTP"
    labelname = "abc"
}
```


# netscaler_appqoeaction

Terraform resource name : ```netscaler_appqoeaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|altcontentpath|No|No|string|
|altcontentsvcname|No|No|[service.name](/doc/resources/service.md)|
|customfile|No|No|string|
|delay|No|No|double|
|dosaction|No|No|SimpleResponse, HICResponse|
|dostrigexpression|No|No|string|
|maxconn|No|No|double|
|name|No|No|string|
|polqdepth|No|No|double|
|priority|No|No|HIGH, MEDIUM, LOW, LOWEST|
|priqdepth|No|No|double|
|respondwith|No|No|ACS, NS|
|tcpprofile|No|No|[nstcpprofile.name](/doc/resources/nstcpprofile.md)|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/appqoe/appqoeaction/appqoeaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_appqoeaction" "tf_name" {

    altcontentpath = "abc"
    altcontentsvcname = "${netscaler_service.<resource_name>.name}"
    customfile = "abc"
    delay = 42
    dosaction = "SimpleResponse"
    dostrigexpression = "abc"
    maxconn = 42
    name = "abc"
    polqdepth = 42
    priority = "HIGH"
    priqdepth = 42
    respondwith = "ACS"
    tcpprofile = "${netscaler_nstcpprofile.<resource_name>.name}"
}
```


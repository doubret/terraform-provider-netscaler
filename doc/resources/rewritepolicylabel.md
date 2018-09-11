# netscaler_rewritepolicylabel

Terraform resource name : ```netscaler_rewritepolicylabel```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|comment|No|No|string|
|labelname|No|No|string|
|transform|No|No|http_req, http_res, othertcp_req, othertcp_res, url, text, clientless_vpn_req, clientless_vpn_res, sipudp_req, sipudp_res, siptcp_req, siptcp_res, diameter_req, diameter_res, radius_req, radius_res, dns_req, dns_res|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/rewrite/rewritepolicylabel/rewritepolicylabel/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_rewritepolicylabel" "tf_name" {

    comment = "abc"
    labelname = "abc"
    transform = "http_req"
}
```


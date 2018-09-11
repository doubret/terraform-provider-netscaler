# netscaler_rewriteaction

Terraform resource name : ```netscaler_rewriteaction```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|bypasssafetycheck|No|No|YES, NO|
|comment|No|No|string|
|name|No|No|string|
|pattern|No|No|string|
|refinesearch|No|No|string|
|search|No|No|string|
|stringbuilderexpr|No|No|string|
|target|No|No|string|
|type|No|No|noop, delete, insert_http_header, delete_http_header, corrupt_http_header, insert_before, insert_after, replace, replace_http_res, delete_all, replace_all, insert_before_all, insert_after_all, clientless_vpn_encode, clientless_vpn_encode_all, clientless_vpn_decode, clientless_vpn_decode_all, insert_sip_header, delete_sip_header, corrupt_sip_header, replace_sip_res, replace_diameter_header_field, replace_dns_header_field, replace_dns_answer_section|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/rewrite/rewriteaction/rewriteaction/) for possible values for these arguments and for an exhaustive list of arguments.

##### Exemple

```
resource "netscaler_rewriteaction" "tf_name" {

    bypasssafetycheck = "YES"
    comment = "abc"
    name = "abc"
    pattern = "abc"
    refinesearch = "abc"
    search = "abc"
    stringbuilderexpr = "abc"
    target = "abc"
    type = "noop"
}
```


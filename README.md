# terraform-provider-netscaler

[Terraform](https://www.terraform.io) Provider for [Citrix NetScaler](https://www.citrix.com/products/netscaler-adc/)

## Description

This project is a terraform custom provider for Citrix NetScaler. It uses the [Nitro API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to create/configure NetScaler resources and bindings. 
This project is largely inspired from [chiradeep](https://github.com/chiradeep)'s work [terraform-provider-netscaler](https://github.com/citrix/terraform-provider-netscaler).
The implementation is completely different though.

**Important note: The provider will not commit the config changes to NetScaler's persistent store.**

## Requirement

* [hashicorp/terraform](https://github.com/hashicorp/terraform)


## Usage

### Running
1. Copy the binary (either from the [build](#building) or from the
   [releases](https://github.com/citrix/terraform-provider-netscaler/releases) page)
   `terraform-provider-netscaler` to an appropriate location.

   [Configure](https://www.terraform.io/docs/plugins/basics.html) `.terraformrc` to use the
   `netscaler` provider. An example `.terraformrc`:

```
providers {
    netscaler = "<path-to-custom-providers>/terraform-provider-netscaler"
}
```

2. Run `terraform` as usual 

```
terraform plan
terraform apply
```

3. The provider will not commit the config changes to NetScaler's persistent store. To do this, run the shell script `ns_commit.sh`:

```
export NS_URL=http://<host>:<port>/
export NS_USER=nsroot
export NS_PASSWORD=nsroot
./ns_commit.sh
```

To ensure that the config is saved on every run, we can use something like `terraform apply && ns_commit.sh`

### Provider Configuration

```
provider "netscaler" {
    username = "${var.ns_user}"
    password = "${var.ns_password}"
    endpoint = "http://10.71.136.250/"
}
```

We can use a `https` URL and accept the untrusted authority certificate on the NetScaler by specifying `insecure_skip_verify = true`

##### Argument Reference

The following arguments are supported.

* `username` - This is the user name to access to NetScaler. Defaults to `nsroot` unless environment variable `NS_LOGIN` has been set
* `password` - This is the password to access to NetScaler. Defaults to `nsroot` unless environment variable `NS_PASSWORD` has been set
* `endpoint` - (Required) Nitro API endpoint in the form `http://<NS_IP>/` or `http://<NS_IP>:<PORT>/`. Can be specified in environment variable `NS_URL`
* `insecure_skip_verify` - (Optional, true/false) Whether to accept the untrusted certificate on the NetScaler when the NetScaler endpoint is `https`

The username, password and endpoint can be provided in environment variables `NS_LOGIN`, `NS_PASSWORD` and `NS_URL`. 

### Resource Configuration

- [appflowaction](doc/resources/appflowaction.md)
- [appflowcollector](doc/resources/appflowcollector.md)
- [appflowpolicy](doc/resources/appflowpolicy.md)
- [appflowpolicylabel](doc/resources/appflowpolicylabel.md)
- [appfwpolicy](doc/resources/appfwpolicy.md)
- [appqoeaction](doc/resources/appqoeaction.md)
- [appqoepolicy](doc/resources/appqoepolicy.md)
- [auditnslogaction](doc/resources/auditnslogaction.md)
- [auditnslogpolicy](doc/resources/auditnslogpolicy.md)
- [auditsyslogaction](doc/resources/auditsyslogaction.md)
- [auditsyslogpolicy](doc/resources/auditsyslogpolicy.md)
- [authorizationpolicy](doc/resources/authorizationpolicy.md)
- [authorizationpolicylabel](doc/resources/authorizationpolicylabel.md)
- [caaction](doc/resources/caaction.md)
- [cachecontentgroup](doc/resources/cachecontentgroup.md)
- [cachepolicy](doc/resources/cachepolicy.md)
- [cachepolicylabel](doc/resources/cachepolicylabel.md)
- [capolicy](doc/resources/capolicy.md)
- [cmpaction](doc/resources/cmpaction.md)
- [cmppolicy](doc/resources/cmppolicy.md)
- [cmppolicylabel](doc/resources/cmppolicylabel.md)
- [csaction](doc/resources/csaction.md)
- [cspolicy](doc/resources/cspolicy.md)
- [cspolicylabel](doc/resources/cspolicylabel.md)
- [csvserver](doc/resources/csvserver.md)
- [dbdbprofile](doc/resources/dbdbprofile.md)
- [dnsaction64](doc/resources/dnsaction64.md)
- [dnspolicy64](doc/resources/dnspolicy64.md)
- [dnsprofile](doc/resources/dnsprofile.md)
- [dospolicy](doc/resources/dospolicy.md)
- [feoaction](doc/resources/feoaction.md)
- [feopolicy](doc/resources/feopolicy.md)
- [filteraction](doc/resources/filteraction.md)
- [filterpolicy](doc/resources/filterpolicy.md)
- [lbgroup](doc/resources/lbgroup.md)
- [lbmetrictable](doc/resources/lbmetrictable.md)
- [lbmonitor](doc/resources/lbmonitor.md)
- [lbprofile](doc/resources/lbprofile.md)
- [lbvserver](doc/resources/lbvserver.md)
- [lbwlm](doc/resources/lbwlm.md)
- [netprofile](doc/resources/netprofile.md)
- [nshttpprofile](doc/resources/nshttpprofile.md)
- [nstcpprofile](doc/resources/nstcpprofile.md)
- [policydataset](doc/resources/policydataset.md)
- [policyexpression](doc/resources/policyexpression.md)
- [policypatset](doc/resources/policypatset.md)
- [policystringmap](doc/resources/policystringmap.md)
- [pqpolicy](doc/resources/pqpolicy.md)
- [responderaction](doc/resources/responderaction.md)
- [responderpolicy](doc/resources/responderpolicy.md)
- [responderpolicylabel](doc/resources/responderpolicylabel.md)
- [rewriteaction](doc/resources/rewriteaction.md)
- [rewritepolicy](doc/resources/rewritepolicy.md)
- [rewritepolicylabel](doc/resources/rewritepolicylabel.md)
- [scpolicy](doc/resources/scpolicy.md)
- [server](doc/resources/server.md)
- [service](doc/resources/service.md)
- [servicegroup](doc/resources/servicegroup.md)
- [spilloveraction](doc/resources/spilloveraction.md)
- [spilloverpolicy](doc/resources/spilloverpolicy.md)
- [tmsessionaction](doc/resources/tmsessionaction.md)
- [tmsessionpolicy](doc/resources/tmsessionpolicy.md)
- [tmtrafficaction](doc/resources/tmtrafficaction.md)
- [tmtrafficpolicy](doc/resources/tmtrafficpolicy.md)
- [transformaction](doc/resources/transformaction.md)
- [transformpolicy](doc/resources/transformpolicy.md)
- [transformpolicylabel](doc/resources/transformpolicylabel.md)
- [transformprofile](doc/resources/transformprofile.md)
- [videooptimizationaction](doc/resources/videooptimizationaction.md)
- [videooptimizationpolicy](doc/resources/videooptimizationpolicy.md)
- [videooptimizationpolicylabel](doc/resources/videooptimizationpolicylabel.md)

### Binding Configuration

- [appflowglobal_appflowpolicy_binding](doc/bindings/appflowglobal_appflowpolicy_binding.md)
- [appflowpolicylabel_appflowpolicy_binding](doc/bindings/appflowpolicylabel_appflowpolicy_binding.md)
- [authorizationpolicylabel_authorizationpolicy_binding](doc/bindings/authorizationpolicylabel_authorizationpolicy_binding.md)
- [cspolicylabel_cspolicy_binding](doc/bindings/cspolicylabel_cspolicy_binding.md)
- [csvserver_appflowpolicy_binding](doc/bindings/csvserver_appflowpolicy_binding.md)
- [csvserver_appfwpolicy_binding](doc/bindings/csvserver_appfwpolicy_binding.md)
- [csvserver_appqoepolicy_binding](doc/bindings/csvserver_appqoepolicy_binding.md)
- [csvserver_auditnslogpolicy_binding](doc/bindings/csvserver_auditnslogpolicy_binding.md)
- [csvserver_auditsyslogpolicy_binding](doc/bindings/csvserver_auditsyslogpolicy_binding.md)
- [csvserver_authorizationpolicy_binding](doc/bindings/csvserver_authorizationpolicy_binding.md)
- [csvserver_cachepolicy_binding](doc/bindings/csvserver_cachepolicy_binding.md)
- [csvserver_cmppolicy_binding](doc/bindings/csvserver_cmppolicy_binding.md)
- [csvserver_cspolicy_binding](doc/bindings/csvserver_cspolicy_binding.md)
- [csvserver_feopolicy_binding](doc/bindings/csvserver_feopolicy_binding.md)
- [csvserver_filterpolicy_binding](doc/bindings/csvserver_filterpolicy_binding.md)
- [csvserver_responderpolicy_binding](doc/bindings/csvserver_responderpolicy_binding.md)
- [csvserver_rewritepolicy_binding](doc/bindings/csvserver_rewritepolicy_binding.md)
- [csvserver_spilloverpolicy_binding](doc/bindings/csvserver_spilloverpolicy_binding.md)
- [csvserver_tmtrafficpolicy_binding](doc/bindings/csvserver_tmtrafficpolicy_binding.md)
- [csvserver_transformpolicy_binding](doc/bindings/csvserver_transformpolicy_binding.md)
- [lbmetrictable_metric_binding](doc/bindings/lbmetrictable_metric_binding.md)
- [lbmonitor_metric_binding](doc/bindings/lbmonitor_metric_binding.md)
- [lbmonitor_sslcertkey_binding](doc/bindings/lbmonitor_sslcertkey_binding.md)
- [lbvserver_appflowpolicy_binding](doc/bindings/lbvserver_appflowpolicy_binding.md)
- [lbvserver_appfwpolicy_binding](doc/bindings/lbvserver_appfwpolicy_binding.md)
- [lbvserver_appqoepolicy_binding](doc/bindings/lbvserver_appqoepolicy_binding.md)
- [lbvserver_auditnslogpolicy_binding](doc/bindings/lbvserver_auditnslogpolicy_binding.md)
- [lbvserver_auditsyslogpolicy_binding](doc/bindings/lbvserver_auditsyslogpolicy_binding.md)
- [lbvserver_authorizationpolicy_binding](doc/bindings/lbvserver_authorizationpolicy_binding.md)
- [lbvserver_cachepolicy_binding](doc/bindings/lbvserver_cachepolicy_binding.md)
- [lbvserver_capolicy_binding](doc/bindings/lbvserver_capolicy_binding.md)
- [lbvserver_cmppolicy_binding](doc/bindings/lbvserver_cmppolicy_binding.md)
- [lbvserver_dnspolicy64_binding](doc/bindings/lbvserver_dnspolicy64_binding.md)
- [lbvserver_feopolicy_binding](doc/bindings/lbvserver_feopolicy_binding.md)
- [lbvserver_filterpolicy_binding](doc/bindings/lbvserver_filterpolicy_binding.md)
- [lbvserver_pqpolicy_binding](doc/bindings/lbvserver_pqpolicy_binding.md)
- [lbvserver_responderpolicy_binding](doc/bindings/lbvserver_responderpolicy_binding.md)
- [lbvserver_rewritepolicy_binding](doc/bindings/lbvserver_rewritepolicy_binding.md)
- [lbvserver_scpolicy_binding](doc/bindings/lbvserver_scpolicy_binding.md)
- [lbvserver_service_binding](doc/bindings/lbvserver_service_binding.md)
- [lbvserver_servicegroup_binding](doc/bindings/lbvserver_servicegroup_binding.md)
- [lbvserver_spilloverpolicy_binding](doc/bindings/lbvserver_spilloverpolicy_binding.md)
- [lbvserver_tmtrafficpolicy_binding](doc/bindings/lbvserver_tmtrafficpolicy_binding.md)
- [lbvserver_transformpolicy_binding](doc/bindings/lbvserver_transformpolicy_binding.md)
- [lbvserver_videooptimizationpolicy_binding](doc/bindings/lbvserver_videooptimizationpolicy_binding.md)
- [policydataset_value_binding](doc/bindings/policydataset_value_binding.md)
- [policypatset_pattern_binding](doc/bindings/policypatset_pattern_binding.md)
- [policystringmap_pattern_binding](doc/bindings/policystringmap_pattern_binding.md)
- [service_dospolicy_binding](doc/bindings/service_dospolicy_binding.md)
- [service_lbmonitor_binding](doc/bindings/service_lbmonitor_binding.md)
- [service_scpolicy_binding](doc/bindings/service_scpolicy_binding.md)
- [servicegroup_lbmonitor_binding](doc/bindings/servicegroup_lbmonitor_binding.md)
- [servicegroup_servicegroupmember_binding](doc/bindings/servicegroup_servicegroupmember_binding.md)
- [sslvserver_ecccurve_binding](doc/bindings/sslvserver_ecccurve_binding.md)
- [sslvserver_sslcertkey_binding](doc/bindings/sslvserver_sslcertkey_binding.md)
- [sslvserver_sslciphersuite_binding](doc/bindings/sslvserver_sslciphersuite_binding.md)

## Using `remote-exec` for one-time tasks
Terraform is useful for maintaining desired state for a set of resources. It is less useful for tasks such as network configuration which don't change. Network configuration is like using a provisioner inside Terraform. The directory `examples/remote-exec` show examples of how Terraform can use ssh to accomplish these one-time tasks.

## Building

### Assumption

* You have (some) experience with Terraform, the different provisioners and providers that come out of the box,
its configuration files, tfstate files, etc.
* You are comfortable with the Go language and its code organization.

1. Install `terraform` from <https://www.terraform.io/downloads.html>
2. Install `dep` (<https://github.com/golang/dep>)
3. Check out this code: `git clone https://<>`
4. Build this code using `make build`

## Samples
See the `examples` directory for various LB topologies that can be driven from this terraform provider.


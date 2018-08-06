generator.exe -i %GOPATH%/src/github.com/doubret/citrix-netscaler-nitro-11-yaml-specs/yml

go fmt github.com/doubret/terraform-provider-netscaler/netscaler/utils
go fmt github.com/doubret/terraform-provider-netscaler/netscaler/resources
go fmt github.com/doubret/terraform-provider-netscaler/netscaler/bindings
go fmt github.com/doubret/terraform-provider-netscaler/netscaler

go fmt

go fmt github.com/doubret/citrix-netscaler-terraform-provider/generator
go build github.com/doubret/citrix-netscaler-terraform-provider/generator

generator.exe -i %GOPATH%/src/github.com/doubret/citrix-netscaler-nitro-11-yaml-specs/yml

go fmt github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils
go fmt github.com/doubret/citrix-netscaler-terraform-provider/netscaler/resources
go fmt github.com/doubret/citrix-netscaler-terraform-provider/netscaler/bindings
go fmt github.com/doubret/citrix-netscaler-terraform-provider/netscaler

go fmt

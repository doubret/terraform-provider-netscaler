go fmt      github.com/doubret/citrix-netscaler-terraform-provider/generator
go build    github.com/doubret/citrix-netscaler-terraform-provider/generator

generator.exe

go fmt      github.com/doubret/citrix-netscaler-terraform-provider/netscaler/resources
go fmt      github.com/doubret/citrix-netscaler-terraform-provider/netscaler/bindings
go fmt      github.com/doubret/citrix-netscaler-terraform-provider/netscaler

go fmt
go build

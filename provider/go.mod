module github.com/pulumi/pulumi-grafana/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/grafana/grafana-api-golang-client v0.0.0-20210720012848-3049c6914b86 // indirect
	github.com/grafana/terraform-provider-grafana v1.11.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.4.1-0.20210714215802-5020116ac4e6
	github.com/pulumi/pulumi/pkg/v3 v3.7.1-0.20210714212650-083fc64ff547 // indirect
	github.com/pulumi/pulumi/sdk/v3 v3.7.0
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f

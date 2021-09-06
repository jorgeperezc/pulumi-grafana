[![Actions Status](https://github.com/pulumi/pulumi-grafana/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-grafana/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fgrafana.svg)](https://npmjs.com/package/@pulumi/grafana)
[![NuGet version](https://badge.fury.io/nu/pulumi.grafana.svg)](https://badge.fury.io/nu/pulumi.grafana)
[![Python version](https://badge.fury.io/py/pulumi-grafana.svg)](https://pypi.org/project/pulumi-grafana)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-grafana/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-grafana/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-grafana/blob/master/LICENSE)

# Grafana Resource Provider

The Grafana resource provider for Pulumi lets you manage grafana resources in your cloud programs. To use this package, please install the Pulumi CLI first.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/grafana

or `yarn`:

    $ yarn add @pulumi/grafana

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_grafana

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-grafana/sdk/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Grafana

## Configuration

The following configuration points are available for the `grafana` provider:

- `grafana:auth` (Required, Sensitive) - API token or basic auth username:password. May alternatively be set via the `GRAFANA_AUTH` environment variable.
- `grafana:url` (Required) - The root URL of a Grafana server. May alternatively be set via the `GRAFANA_URL` environment variable.
- `grafana:ca_cert` (Optional) Certificate CA bundle to use to verify the Grafana server's certificate. May alternatively be set via the `GRAFANA_CA_CERT` environment variable.
- `grafana:insecure_skip_verify` (Optional) Skip TLS certificate verification. May alternatively be set via the `GRAFANA_INSECURE_SKIP_VERIFY` environment variable.
- `grafana:org_id` (Optional) The organization id to operate on within grafana. May alternatively be set via the `GRAFANA_ORG_ID` environment variable.
- `grafana:sm_access_token` (Optional, Sensitive) A Synthetic Monitoring access token. May alternatively be set via the `GRAFANA_SM_ACCESS_TOKEN` environment variable.
- `grafana:sm_url` (Optional) Synthetic monitoring backend address. May alternatively be set via the `GRAFANA_SM_URL` environment variable.
- `grafana:tls_cert` (Optional) Client TLS certificate file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_CERT` environment variable.
- `grafana:tls_key` (Optional) Client TLS key file to use to authenticate to the Grafana server. May alternatively be set via the `GRAFANA_TLS_KEY` environment variable.

## Reference

For detailed reference documentation, please visit [the API docs][1].


[1]: https://www.pulumi.com/docs/reference/pkg/grafana/

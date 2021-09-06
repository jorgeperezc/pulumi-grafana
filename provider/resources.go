// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grafana

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi-grafana/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/grafana/terraform-provider-grafana/grafana"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "grafana"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(grafana.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "grafana",
		Description: "A Pulumi package for creating and managing grafana cloud resources.",
		Keywords:    []string{"pulumi", "grafana"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-grafana",
		Config: map[string]*tfbridge.SchemaInfo{
		},
		Resources: map[string]*tfbridge.ResourceInfo{
            "grafana_alert_notification": {Tok: makeResource(mainMod, "AlertNotification")},
            "grafana_builtin_role_assignment": {Tok: makeResource(mainMod, "BuiltinRoleAssignment")},
            "grafana_dashboard": {Tok: makeResource(mainMod, "Dashboard")},
            "grafana_dashboard_permission": {Tok: makeResource(mainMod, "DashboardPermission")},
            "grafana_data_source": {Tok: makeResource(mainMod, "DataSource")},
            "grafana_folder": {Tok: makeResource(mainMod, "Folder")},
            "grafana_folder_permission": {Tok: makeResource(mainMod, "FolderPermission")},
            "grafana_organization": {Tok: makeResource(mainMod, "Organization")},
            "grafana_role": {Tok: makeResource(mainMod, "Role")},
            "grafana_synthetic_monitoring_check": {Tok: makeResource(mainMod, "SyntheticMonitoringCheck")},
            "grafana_synthetic_monitoring_probe": {Tok: makeResource(mainMod, "SyntheticMonitoringProbe")},
            "grafana_team": {Tok: makeResource(mainMod, "Team")},
            "grafana_team_external_group": {Tok: makeResource(mainMod, "TeamExternalGroup")},
            "grafana_team_preferences": {Tok: makeResource(mainMod, "TeamPreferences")},
            "grafana_user": {Tok: makeResource(mainMod, "User")},
        },
        DataSources: map[string]*tfbridge.DataSourceInfo{
            "grafana_synthetic_monitoring_probe": {Tok: makeDataSource(mainMod, "getSyntheticMonitoringProbe")},
            "grafana_synthetic_monitoring_probes": {Tok: makeDataSource(mainMod, "getSyntheticMonitoringProbes")},
        },
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	"nexus_repository_apt_hosted": config.NameAsIdentifier,
	"nexus_repository_apt_proxy": config.NameAsIdentifier,
	"nexus_repository_cargo_group": config.NameAsIdentifier,
	"nexus_repository_cargo_hosted": config.NameAsIdentifier,
	"nexus_repository_cargo_proxy": config.NameAsIdentifier,
	"nexus_repository_cocoapods_proxy": config.NameAsIdentifier,
	"nexus_repository_conan_proxy": config.NameAsIdentifier,
	"nexus_repository_conda_proxy": config.NameAsIdentifier,
	"nexus_repository_docker_group": config.NameAsIdentifier,
	"nexus_repository_docker_hosted": config.NameAsIdentifier,
	"nexus_repository_docker_proxy": config.NameAsIdentifier,
	"nexus_repository_gitlfs_hosted": config.NameAsIdentifier,
	"nexus_repository_go_group": config.NameAsIdentifier,
	"nexus_repository_go_proxy": config.NameAsIdentifier,
	"nexus_repository_helm_hosted": config.NameAsIdentifier,
	"nexus_repository_helm_proxy": config.NameAsIdentifier,
	"nexus_repository_maven_group": config.NameAsIdentifier,
	"nexus_repository_maven_hosted": config.NameAsIdentifier,
	"nexus_repository_maven_proxy": config.NameAsIdentifier,
	"nexus_repository_npm_group": config.NameAsIdentifier,
	"nexus_repository_npm_hosted": config.NameAsIdentifier,
	"nexus_repository_npm_proxy": config.NameAsIdentifier,
	"nexus_repository_nuget_group": config.NameAsIdentifier,
	"nexus_repository_nuget_hosted": config.NameAsIdentifier,
	"nexus_repository_nuget_proxy": config.NameAsIdentifier,
	"nexus_repository_p2_proxy": config.NameAsIdentifier,
	"nexus_repository_pypi_group": config.NameAsIdentifier,
	"nexus_repository_pypi_hosted": config.NameAsIdentifier,
	"nexus_repository_pypi_proxy": config.NameAsIdentifier,
	"nexus_repository_r_group": config.NameAsIdentifier,
	"nexus_repository_r_hosted": config.NameAsIdentifier,
	"nexus_repository_r_proxy": config.NameAsIdentifier,
	"nexus_repository_raw_group": config.NameAsIdentifier,
	"nexus_repository_raw_hosted": config.NameAsIdentifier,
	"nexus_repository_raw_proxy": config.NameAsIdentifier,
	"nexus_repository_rubygems_group": config.NameAsIdentifier,
	"nexus_repository_rubygems_hosted": config.NameAsIdentifier,
	"nexus_repository_rubygems_proxy": config.NameAsIdentifier,
	"nexus_repository_yum_group": config.NameAsIdentifier,
	"nexus_repository_yum_hosted": config.NameAsIdentifier,
	"nexus_repository_yum_proxy": config.NameAsIdentifier,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	repositoryMavenGroupCluster  "github.com/seelmn/provider-nexus/config/cluster/repositorymavengroup"
	repositoryMavenHostedCluster "github.com/seelmn/provider-nexus/config/cluster/repositorymavenhosted"
	repositoryMavenProxyCluster  "github.com/seelmn/provider-nexus/config/cluster/repositorymavenproxy"
	repositoryNpmGroupCluster    "github.com/seelmn/provider-nexus/config/cluster/repositorynpmgroup"
	repositoryNpmHostedCluster   "github.com/seelmn/provider-nexus/config/cluster/repositorynpmhosted"
	repositoryNpmProxyCluster    "github.com/seelmn/provider-nexus/config/cluster/repositorynpmproxy"
	repositoryPypiGroupCluster   "github.com/seelmn/provider-nexus/config/cluster/repositorypypigroup"
	repositoryPypiHostedCluster  "github.com/seelmn/provider-nexus/config/cluster/repositorypypihosted"
	repositoryPypiProxyCluster   "github.com/seelmn/provider-nexus/config/cluster/repositorypypiproxy"
	repositoryRawGroupCluster    "github.com/seelmn/provider-nexus/config/cluster/repositoryrawgroup"
	repositoryRawHostedCluster   "github.com/seelmn/provider-nexus/config/cluster/repositoryrawhosted"
	repositoryRawProxyCluster    "github.com/seelmn/provider-nexus/config/cluster/repositoryrawproxy"
	repositoryYumGroupCluster    "github.com/seelmn/provider-nexus/config/cluster/repositoryyumgroup"
	repositoryYumHostedCluster   "github.com/seelmn/provider-nexus/config/cluster/repositoryyumhosted"
	repositoryYumProxyCluster    "github.com/seelmn/provider-nexus/config/cluster/repositoryyumproxy"

	repositoryMavenGroupNamespaced  "github.com/seelmn/provider-nexus/config/namespaced/repositorymavengroup"
	repositoryMavenHostedNamespaced "github.com/seelmn/provider-nexus/config/namespaced/repositorymavenhosted"
	repositoryMavenProxyNamespaced  "github.com/seelmn/provider-nexus/config/namespaced/repositorymavenproxy"
	repositoryNpmGroupNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositorynpmgroup"
	repositoryNpmHostedNamespaced   "github.com/seelmn/provider-nexus/config/namespaced/repositorynpmhosted"
	repositoryNpmProxyNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositorynpmproxy"
	repositoryPypiGroupNamespaced   "github.com/seelmn/provider-nexus/config/namespaced/repositorypypigroup"
	repositoryPypiHostedNamespaced  "github.com/seelmn/provider-nexus/config/namespaced/repositorypypihosted"
	repositoryPypiProxyNamespaced   "github.com/seelmn/provider-nexus/config/namespaced/repositorypypiproxy"
	repositoryRawGroupNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositoryrawgroup"
	repositoryRawHostedNamespaced   "github.com/seelmn/provider-nexus/config/namespaced/repositoryrawhosted"
	repositoryRawProxyNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositoryrawproxy"
	repositoryYumGroupNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositoryyumgroup"
	repositoryYumHostedNamespaced   "github.com/seelmn/provider-nexus/config/namespaced/repositoryyumhosted"
	repositoryYumProxyNamespaced    "github.com/seelmn/provider-nexus/config/namespaced/repositoryyumproxy"
)

const (
	resourcePrefix = "nexus"
	modulePath     = "github.com/seelmn/provider-nexus"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("nexus.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		repositoryMavenGroupCluster.Configure,
		repositoryMavenHostedCluster.Configure,
		repositoryMavenProxyCluster.Configure,
		repositoryNpmGroupCluster.Configure,
		repositoryNpmHostedCluster.Configure,
		repositoryNpmProxyCluster.Configure,
		repositoryPypiGroupCluster.Configure,
		repositoryPypiHostedCluster.Configure,
		repositoryPypiProxyCluster.Configure,
		repositoryRawGroupCluster.Configure,
		repositoryRawHostedCluster.Configure,
		repositoryRawProxyCluster.Configure,
		repositoryYumGroupCluster.Configure,
		repositoryYumHostedCluster.Configure,
		repositoryYumProxyCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("nexus.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		repositoryMavenGroupNamespaced.Configure,
		repositoryMavenHostedNamespaced.Configure,
		repositoryMavenProxyNamespaced.Configure,
		repositoryNpmGroupNamespaced.Configure,
		repositoryNpmHostedNamespaced.Configure,
		repositoryNpmProxyNamespaced.Configure,
		repositoryPypiGroupNamespaced.Configure,
		repositoryPypiHostedNamespaced.Configure,
		repositoryPypiProxyNamespaced.Configure,
		repositoryRawGroupNamespaced.Configure,
		repositoryRawHostedNamespaced.Configure,
		repositoryRawProxyNamespaced.Configure,
		repositoryYumGroupNamespaced.Configure,
		repositoryYumHostedNamespaced.Configure,
		repositoryYumProxyNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	repositoryMavenHostedCluster    "github.com/seelmn/provider-nexus/config/cluster/repositorymavenhosted"
	repositoryPypiHostedCluster     "github.com/seelmn/provider-nexus/config/cluster/repositorypypihosted"

	repositoryMavenHostedNamespaced "github.com/seelmn/provider-nexus/config/namespaced/repositorymavenhosted"
	repositoryPypiHostedNamespaced  "github.com/seelmn/provider-nexus/config/namespaced/repositorypypihosted"
)

const (
	resourcePrefix = "template"
	modulePath     = "github.com/crossplane/upjet-provider-template"
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
		repositoryMavenHostedCluster.Configure,
		repositoryPypiHostedCluster.Configure, 
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
		repositoryMavenHostedCluster.Configure,
		repositoryPypiHostedCluster.Configure, 
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

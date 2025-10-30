package repositorynpmgroup

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("nexus_repository_npm_group", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "harbor"
        r.ShortGroup = "repositorynpmgroup"
        r.Kind       = "RepositoryNpmGroup"
    })
}

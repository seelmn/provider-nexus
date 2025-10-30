package repositorynpmproxy

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("nexus_repository_npm_proxy", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "harbor"
        r.ShortGroup = "repositorynpmproxy"
        r.Kind       = "RepositoryNpmProxy"
    })
}

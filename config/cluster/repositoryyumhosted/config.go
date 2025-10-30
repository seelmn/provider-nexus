package repositoryyumhosted

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("nexus_repository_yum_hosted", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "nexus"
        r.ShortGroup = "repository_yum_hosted"
        r.Kind       = "repository_yum_hosted"
    })
}

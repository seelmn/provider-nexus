package repositorymavenhosted

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("sonatyperepo_repository_maven_hosted", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "sonatyperepo"
        r.ShortGroup = "repository"
		r.Kind       = "MavenHosted"
    })
}
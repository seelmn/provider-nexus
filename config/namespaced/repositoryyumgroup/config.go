package repositoryyumgroup

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("nexus_repository_yum_group", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "nexus"
        r.ShortGroup = "repositoryyumgroup"
        r.Kind       = "RepositoryYumGroup"

        r.References["group.member_names"] = config.Reference{
           TerraformName: "nexus_repository_yum_hosted",
           TerraformName: "nexus_repository_yum_proxy",
        }
    })
}

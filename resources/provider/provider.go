package provider

import (
	"embed"
	// CHANGEME: change the following to your own package
	"github.com/ovotech/cq-provider-cloudflare/client"
	"github.com/ovotech/cq-provider-cloudflare/resources"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*/*.sql
	providerMigrations embed.FS
	Version            = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version: Version,
		// CHANGEME: Change to your provider name
		Name:      "cloudflare",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			// CHANGEME: place here all supported resources
			"zone": resources.CloudflareZoneResource(),
			"waf":  resources.CloudflareWafPackageResource(),
		},
		Migrations: providerMigrations,
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}

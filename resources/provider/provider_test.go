package provider_test

import (
	"testing"

	"github.com/ovotech/cq-provider-cloudflare/resources/provider"

	"github.com/cloudquery/cq-provider-sdk/migration"
)

func TestMigrations(t *testing.T) {
	migration.RunMigrationsTest(t, provider.Provider(), []string{"latest"})
}

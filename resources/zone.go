package resources

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/ovotech/cq-provider-cloudflare/client"
)

func Zones() *schema.Table {
	return &schema.Table{
		Name:     "zones",
		Resolver: nil,
		// Those are optional
		// DeleteFilter: nil,
		// Multiplex:    nil,
		// IgnoreError:  nil,
		//PostResourceResolver: nil,
		Options: schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("node_id"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)

	// api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	zones, err := c.ThirdPartyClient.ListZones(ctx)

	// Fetch zone details on the account
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Print zone details
	fmt.Println(zones)

	// opts := github.RepositoryListOptions{}
	// // repositories, response, err := c.GithubClient.Repositories.List(ctx, "cloudquery", &opts)
	// if err != nil {
	// 	return err
	// }
	// _ = response
	res <- zones
	return nil
}

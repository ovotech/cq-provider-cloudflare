package resources

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	// CHANGEME change this to your package name
	"github.com/ovotech/cq-provider-cloudflare/client"
)

func CloudflareZoneResource() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_zones",
		Resolver: fetchZoneResources,
		// Those are optional
		// DeleteFilter: nil,
		// Multiplex:    nil,
		// IgnoreError:  nil,
		//PostResourceResolver: nil,

		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Description: "Description of the column to appear in the generated documentation",
				//Resolver: provider.ResolveAWSAccount,
			},
			{
				Name: "region",
				Type: schema.TypeString,
				//Resolver: fetchS3BucketLocation,
			},
			{
				Name: "creation_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("other_name_in_struct"),
			},
		},
		// A table can have relations
		//Relations: []*schema.Table{
		//	{
		//		Name:     "provider_demo_resource_children",
		//		Resolver: fetchDemoResourceChildren,
		//		Columns: []schema.Column{
		//			{
		//				Name:     "bucket_id",
		//				Type:     schema.TypeUUID,
		//				Resolver: schema.ParentIdResolver,
		//			},
		//			{
		//				Name:     "resource_id",
		//				Type:     schema.TypeString,
		//				Resolver: schema.PathResolver("Grantee.ID"),
		//			},
		//			{
		//				Name:     "type",
		//				Type:     schema.TypeString,
		//				Resolver: schema.PathResolver("Grantee.Type"),
		//			},
		//		},
		//	},
		//},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchZoneResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	_ = c
	// Fetch using the third party client and put the result in res
	// res <- c.ThirdPartyClient.getDat()

	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	// Fetch zone details on the account
	zones, err := api.ListZones(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zones)

	res <- zones
	return nil
}

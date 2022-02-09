package resources

import (
	"context"
	"log"

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
				Name:        "ID",
				Type:        schema.TypeString,
				Description: "Description of the column to appear in the generated documentation",
				// Resolver:    schema.PathResolver("id"),
			},
			{
				Name: "Name",
				Type: schema.TypeString,
				// Resolver: schema.PathResolver("name"),
			},
			{
				Name:     "Account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Name"),
			},
			{
				Name:     "NameServers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NameServers"),
			},
			{
				Name:     "Status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "CreatedOn",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "ModifiedOn",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ModifiedOn"),
			},
			{
				Name:     "ActivatedOn",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ActivatedOn"),
			},
			{
				Name:     "OwnerEmail",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Email"),
			},
			{
				Name:     "OwnerName",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Name"),
			},
			{
				Name:     "HostName",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Host.Name"),
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

	// Most API calls require a Context
	// Fetch zone details on the account
	zones, err := c.ThirdPartyClient.ListZones(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	// for _, v := range zones {
	// 	// c.Logger().Debug(v.Name)
	// }

	res <- zones
	return nil
}

package resources

import (
	"context"
	"log"

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
				Name:        "ID",
				Type:        schema.TypeString,
				Description: "ZoneID",
				Resolver:    schema.PathResolver("ID"),
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
		Relations: []*schema.Table{
			{
				Name:     "cloudflare_zone_waf_packages",
				Resolver: fetchWafChildResources,
				Columns: []schema.Column{
					{
						Name: "ZoneID",
						Type: schema.TypeString,
					},
					{
						Name: "Name",
						Type: schema.TypeString,
					},
					{
						Name:     "ID",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "DetectionMode",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "cloudflare_zone_dns_records",
				Resolver: fetchZoneDNSResources,
				Columns: []schema.Column{
					{
						Name: "ZoneID",
						Type: schema.TypeString,
					},
					{
						Name: "Name",
						Type: schema.TypeString,
					},
					{
						Name: "Content",
						Type: schema.TypeString,
					},
					{
						Name: "ZoneName",
						Type: schema.TypeString,
					},
					{
						Name: "Priority",
						Type: schema.TypeBigInt,
					},
					{
						Name:     "TTL",
						Type:     schema.TypeSmallInt,
						Resolver: schema.IntResolver("TTL"),
					},
					{
						Name: "Proxied",
						Type: schema.TypeBool,
					},
					{
						Name: "Locked",
						Type: schema.TypeBool,
					},
					{
						Name:     "ID",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "DetectionMode",
						Type: schema.TypeString,
					},
				},
			},
		},
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

func fetchWafChildResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	_ = c
	p := parent.Item.(cloudflare.Zone)

	// Most API calls require a Context
	// Fetch zone details on the account
	zone_waf_packages, err := c.ThirdPartyClient.ListWAFPackages(ctx, p.ID)
	if err != nil {
		log.Fatal(err)
	}

	res <- zone_waf_packages
	return nil
}

func fetchZoneDNSResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	_ = c
	p := parent.Item.(cloudflare.Zone)

	// Most API calls require a Context
	// Fetch zone details on the account
	dns_records, err := c.ThirdPartyClient.DNSRecords(ctx, p.ID, cloudflare.DNSRecord{})
	if err != nil {
		log.Fatal(err)
	}

	res <- dns_records
	return nil
}

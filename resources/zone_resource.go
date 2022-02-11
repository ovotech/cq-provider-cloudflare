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
				Name:        "id",
				Type:        schema.TypeString,
				Description: "Cloudflare ZoneID",
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Account.Name"),
			},
			{
				Name:     "name_servers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NameServers"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "created_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedOn"),
			},
			{
				Name:     "modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ModifiedOn"),
			},
			{
				Name:     "activated_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ActivatedOn"),
			},
			{
				Name:     "owner_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Email"),
			},
			{
				Name:     "owner_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner.Name"),
			},
			{
				Name:     "host_name",
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
						Name:     "zone_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ZoneID"),
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Name"),
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "detection_mode",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DetectionMode"),
					},
				},
			},
			{
				Name:     "cloudflare_zone_dns_records",
				Resolver: fetchZoneDNSResources,
				Columns: []schema.Column{
					{
						Name:     "zone_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ZoneID"),
					},
					{
						Name:     "name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Name"),
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name:     "content",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Content"),
					},
					{
						Name:     "zone_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ZoneName"),
					},
					{
						Name:     "priority",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("Priority"),
					},
					// {
					// 	Name: "TTL",
					// 	Type: schema.TypeString,
					// },
					{
						Name:     "proxied",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Proxied"),
					},
					{
						Name:     "locked",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Locked"),
					},
					{
						Name:     "detection_mode",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DetectionMode"),
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

package resources

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func ZoneResource() *schema.Table {
	return &schema.Table{
		Name:     "demo_table",
		Resolver: fetchDemoResources,
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

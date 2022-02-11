-- Autogenerated by migration tool on 2022-02-11 15:22:35

-- Resource: waf
CREATE TABLE IF NOT EXISTS "cloudflare_waf_package" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"ID" text,
	"Description" text,
	"Action" text,
	"Filter" jsonb,
	"CreatedOn" timestamp without time zone,
	"ModifiedOn" timestamp without time zone,
	"Paused" boolean,
	CONSTRAINT cloudflare_waf_package_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);

-- Resource: zone
CREATE TABLE IF NOT EXISTS "cloudflare_zones" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"ID" text,
	"Name" text,
	"Account" text,
	"NameServers" text[],
	"Status" text,
	"CreatedOn" timestamp without time zone,
	"ModifiedOn" timestamp without time zone,
	"ActivatedOn" timestamp without time zone,
	"OwnerEmail" text,
	"OwnerName" text,
	"HostName" text,
	CONSTRAINT cloudflare_zones_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "cloudflare_zone_waf_packages" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"ZoneID" text,
	"Name" text,
	"ID" text,
	"DetectionMode" text,
	CONSTRAINT cloudflare_zone_waf_packages_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "cloudflare_zone_dns_records" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"ZoneID" text,
	"Name" text,
	"Content" text,
	"ZoneName" text,
	"Priority" bigint,
	"TTL" smallint,
	"Proxied" boolean,
	"Locked" boolean,
	"ID" text,
	"DetectionMode" text,
	CONSTRAINT cloudflare_zone_dns_records_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);

-- Autogenerated by migration tool on 2022-02-11 15:37:16

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
	"id" text,
	"name" text,
	"account" text,
	"name_servers" text[],
	"status" text,
	"created_on" timestamp without time zone,
	"modified_on" timestamp without time zone,
	"activated_on" timestamp without time zone,
	"owner_email" text,
	"owner_name" text,
	"host_name" text,
	CONSTRAINT cloudflare_zones_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "cloudflare_zone_waf_packages" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"zone_id" text,
	"name" text,
	"id" text,
	"detection_mode" text,
	CONSTRAINT cloudflare_zone_waf_packages_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "cloudflare_zone_dns_records" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"zone_id" text,
	"name" text,
	"id" text,
	"content" text,
	"zone_name" text,
	"priority" bigint,
	"proxied" boolean,
	"locked" boolean,
	"detection_mode" text,
	CONSTRAINT cloudflare_zone_dns_records_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id)
);

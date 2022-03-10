package client

import (
	"errors"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	// CHANGEME:  Usually you store here your 3rd party clients and use them in the fetcher
	ThirdPartyClient cloudflare.API
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	_ = providerConfig
	// Init client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")

	if apiToken == "" {
		apiToken = providerConfig.CloudflareToken
		if apiToken == "" {
			log.Fatal("Missing Cloudflare API Token")
			return nil, errors.New("Missing Cloudflare API Token")
		}
	}

	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Most API calls require a Context
	// Fetch zone details on the account

	client := Client{
		logger:           logger,
		ThirdPartyClient: *api,
	}

	// Return the initialized client and it will be passed to your resources
	return &client, nil
}

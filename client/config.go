package client

// Provider Configuration

type Config struct {
	// ADD THIS LINE:
	GitHubToken string `hcl:"cloudflare_token,optional"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []struct {
		Name  string
		Other map[string]interface{} `hcl:",inline"`
	}
}

func (c Config) Example() string {
	return `configuration {
    // Add this line    
    // api_key = ${your_env_variable}
    // api_key = static_api_key
}
`
}

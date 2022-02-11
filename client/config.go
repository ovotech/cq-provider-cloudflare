package client

// Provider Configuration

type Config struct {
	// ADD THIS LINE:
	CloudflareToken string    `hcl:"cloudflare_token,optional"`
	Account         []Account `hcl:"account,block"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []struct {
		Name  string
		Other map[string]interface{} `hcl:",inline"`
	}
}

type Account struct {
	Name string `hcl:"name,optional"`
	ID   string `hcl:"ID,optional"`
}

// type Config struct {
// 	User  string `hcl:"user,optional"`
// 	Debug bool   `hcl:"debug,optional"`
// }

func (c Config) Example() string {
	return `configuration {
    // Add this line    
    // api_key = ${your_env_variable}
    // api_key = static_api_key
}
`
}

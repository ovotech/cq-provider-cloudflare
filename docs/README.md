# Configuration


Add the following to your `config.hcl` file

```provider "cloudflare" {
  configuration {
    // Add this line    
    cloudflare_token = "your_cloudflare_api_token"
  }
   
  accounts {
    name = "Foo"
    ID = "1111222233334444"

  }
  // list of resources to fetch
  resources = [
    "zone",
  ]
  // enables partial fetching, allowing for any failures to not stop full resource pull
  enable_partial_fetch = true
}

Alternatively, if you dont want to include your cloudflare_token in the config, you can set the `CLOUDFLARE_API_TOKEN` environment variable instead.
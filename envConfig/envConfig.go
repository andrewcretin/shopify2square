package envConfig

import "os"

type Environment struct {
	Stage             string
	Region            string
	Service           string
	Alias             string
	ShopifyKey        string
	ShopifyPassword   string
	SquareAccessToken string
}

//noinspection ALL
func CurrentEnvironment() Environment {
	env := Environment{}
	env.Stage = os.Getenv("STAGE")
	if len(env.Stage) == 0 {
		env.Stage = "staging"
	}
	env.Region = os.Getenv("REGION")
	if len(env.Region) == 0 {
		env.Region = "us-west-2"
	}
	env.Service = os.Getenv("SERVICE")
	if len(env.Service) == 0 {
		env.Service = "shopify2square"
	}
	env.Alias = os.Getenv("ALIAS")
	if len(env.Alias) == 0 {
		env.Alias = env.Stage
	}

	env.ShopifyKey = os.Getenv("SHOPIFY_KEY")
	env.ShopifyPassword = os.Getenv("SHOPIFY_PASSWORD")
	env.SquareAccessToken = os.Getenv("SQUARE_ACCESS_TOKEN")

	return env
}

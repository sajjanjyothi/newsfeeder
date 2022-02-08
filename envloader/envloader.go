package envloader

import "github.com/kelseyhightower/envconfig"

type ENV struct {
	RedisURL        string `envconfig:"REDIS_URL"`
	ServicePort     string `envconfig:"SERVICE_PORT"`
	RedisPassword   string `envconfig:"REDIS_PASSWORD"`
	NewsFedderToken string `envconfig:"NEWS_TOKEN"`
}

func LoadConfig(environment *ENV) error {

	err := envconfig.Process("newsfeeder", environment)
	return err
}

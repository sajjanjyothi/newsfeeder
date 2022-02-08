package envloader

import "github.com/kelseyhightower/envconfig"

type ENV struct {
	RedisURL    string `envconfig:"REDIS_URL"`
	ServicePort string `envconfig:"SERVICE_PORT"`
}

func LoadConfig(environment *ENV) error {

	err := envconfig.Process("newsfeeder", environment)
	return err
}

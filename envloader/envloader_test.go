package envloader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Env(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{
			key:   "REDIS_URL",
			value: "Test1",
		},
		{
			key:   "SERVICE_PORT",
			value: "Test_Port",
		},
	}

	for _, test := range tests {
		os.Setenv(test.key, test.value)
	}
	env := &ENV{}
	LoadConfig(env)

	assert.Equal(t, env.RedisURL, "Test1", "REDIS URL not matching")
	assert.Equal(t, env.ServicePort, "Test_Port", "SERVICE_PORT not matching")
}

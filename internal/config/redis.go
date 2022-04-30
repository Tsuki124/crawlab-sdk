package config

import (
	"crawlab-sdk/internal/constants"
	"crawlab-sdk/internal/interfaces"
	"os"
)

type WithRedisConfig struct {
	interfaces.WithConfig
}

func (my *WithRedisConfig) GetConfigMap() map[string]string {
	keys := []string{
		constants.ENV_REDIS_HOST,
		constants.ENV_REDIS_PORT,
		constants.ENV_REDIS_USERNAME,
		constants.ENV_REDIS_PASSWORD,
	}

	configMap := make(map[string]string)
	for _, key := range keys {
		configMap[key] = os.Getenv(key)
	}

	return configMap
}

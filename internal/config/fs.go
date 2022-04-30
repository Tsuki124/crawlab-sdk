package config

import (
	"crawlab-sdk/internal/constants"
	"crawlab-sdk/internal/interfaces"
	"os"
)

type WithFsConfig struct {
	interfaces.WithConfig
}

func (my *WithFsConfig) GetConfigMap() map[string]string {
	keys := []string{
		constants.ENV_SEAWEED_FS_FILER_URL,
	}

	configMap := make(map[string]string)
	for _, key := range keys {
		configMap[key] = os.Getenv(key)
	}

	return configMap
}


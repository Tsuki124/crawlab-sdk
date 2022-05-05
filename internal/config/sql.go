package config

import (
	"github.com/Tsuki124/crawlab-sdk/internal/constants"
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"os"
)

type WithSQLConfig struct {
	interfaces.WithConfig
}

func (my *WithSQLConfig) GetConfigMap() map[string]string {
	keys := []string{
		constants.ENV_SQL_DRIVER,
		constants.ENV_SQL_HOST,
		constants.ENV_SQL_PORT,
		constants.ENV_SQL_USERNAME,
		constants.ENV_SQL_PASSWORD,
	}

	configMap := make(map[string]string)
	for _, key := range keys {
		configMap[key] = os.Getenv(key)
	}

	return configMap
}

package driver

import (
	"crawlab-sdk/internal/config"
	"crawlab-sdk/internal/constants"
	"fmt"
	"github.com/crawlab-team/go-trace"
	"github.com/go-redis/redis"
)

var Redis = redisDriver{}

type redisDriver struct {
	_conf config.WithRedisConfig
}

func (my *redisDriver) New(db int) (*redis.Client, error) {
	//获取配置参数
	configMap := my._conf.GetConfigMap()
	addr := fmt.Sprintf("%s:%s", configMap[constants.ENV_REDIS_HOST], configMap[constants.ENV_REDIS_PORT])
	password := configMap[constants.ENV_REDIS_PASSWORD]
	//创建数据库连接
	engine := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	//连通性检查
	_, err := engine.Ping().Result()
	if err != nil {
		return nil, trace.Error(err)
	}

	return engine, nil
}

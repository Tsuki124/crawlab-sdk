package driver

import (
	"context"
	"crawlab-sdk/internal/config"
	"crawlab-sdk/internal/constants"
	"fmt"
	"github.com/crawlab-team/go-trace"
	"github.com/qiniu/qmgo"
)

var Mongo = mongoDriver{}

type mongoDriver struct {
	_conf config.WithMongoConfig
}

func (my *mongoDriver) New(name string) (*qmgo.Database, error) {
	//获取配置参数
	configMap := my._conf.GetConfigMap()

	//创建数据库连接
	ctx := context.Background()
	uri := fmt.Sprintf(
		constants.KEY_DATA_SOURCE_NAME_MONGO,
		configMap[constants.ENV_MONGO_HOST],
		configMap[constants.ENV_MONGO_PORT])
	auth := &qmgo.Credential{
		AuthMechanism: configMap[constants.ENV_MONGO_AUTH_MECHANISM],
		AuthSource:    configMap[constants.ENV_MONGO_AUTH_SOURCE],
		Username:      configMap[constants.ENV_MONGO_USERNAME],
		Password:      configMap[constants.ENV_MONGO_PASSWORD],
		PasswordSet:   false,
	}

	conf := &qmgo.Config{
		Uri:      uri,
		Database: name,
		Auth:     auth,
	}
	engine, err := qmgo.Open(ctx, conf)
	if err != nil {
		return nil, trace.Error(err)
	}

	//连通性检查
	if err = engine.Ping(5); err != nil {
		return nil, trace.Error(err)
	}

	return engine.Database, nil
}

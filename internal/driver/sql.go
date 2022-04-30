package driver

import (
	"crawlab-sdk/internal/config"
	"crawlab-sdk/internal/constants"
	"errors"
	"fmt"
	"github.com/crawlab-team/go-trace"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var SQL = sqlDriver{}

type sqlDriver struct {
	_conf config.WithSQLConfig
}

func (my *sqlDriver) New(name string) (*gorm.DB, error) {
	//获取配置参数
	configMap := my._conf.GetConfigMap()

	//创建数据库连接
	var keyDataSourceName string
	var openDialectorHander func(dsn string) gorm.Dialector
	switch configMap[constants.ENV_SQL_DRIVER] {
	case constants.KEY_SQL_DRIVER_MYSQL:
		keyDataSourceName = constants.KEY_DATA_SOURCE_NAME_MYSQL
		openDialectorHander = mysql.Open
	case constants.KEY_SQL_DRIVER_POSTGRES:
		keyDataSourceName = constants.KEY_DATA_SOURCE_NAME_POSTGRES
		openDialectorHander = postgres.Open
	default:
		return nil, trace.Error(errors.New("not found the sql driver"))
	}

	dsn := fmt.Sprintf(
		keyDataSourceName,
		configMap[constants.ENV_SQL_USERNAME],
		configMap[constants.ENV_SQL_PASSWORD],
		configMap[constants.ENV_SQL_HOST],
		configMap[constants.ENV_SQL_PORT],
		name)

	dialector := openDialectorHander(dsn)
	engine, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, trace.Error(err)
	}

	return engine, nil
}

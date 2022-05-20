package driver

import (
	"errors"
	"fmt"
	"github.com/Tsuki124/crawlab-sdk/internal/config"
	"github.com/Tsuki124/crawlab-sdk/internal/constants"
	"github.com/crawlab-team/go-trace"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             10 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Silent,    // 日志级别
			IgnoreRecordNotFoundError: true,             // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,            // 禁用彩色打印
		},
	)
	dialector := openDialectorHander(dsn)
	engine, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, trace.Error(err)
	}

	return engine, nil
}

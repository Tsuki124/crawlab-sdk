package crawlab_sdk

import (
	"github.com/Tsuki124/crawlab-sdk/internal/engines"
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"sync"
)

var SQL = sqlService{_DBs: make(map[string]interfaces.SQLDb)}

type sqlService struct {
	_Mtx sync.RWMutex
	_DBs map[string]interfaces.SQLDb
}

func (my *sqlService) Db(name string) interfaces.SQLDb {
	my._Mtx.RLock()
	db, ok := my._DBs[name]
	my._Mtx.RUnlock()
	if ok {
		return db
	}

	//根据配置创建数据库连接句柄
	my._Mtx.Lock()
	db = engines.NewSQLDb(name)
	my._DBs[name] = db
	my._Mtx.Unlock()

	return db
}

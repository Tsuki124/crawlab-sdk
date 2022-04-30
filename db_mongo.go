package crawlab_sdk

import (
	"crawlab-sdk/internal/engines"
	"crawlab-sdk/internal/interfaces"
	"sync"
)

var Mongo = mongoService{_DBs: make(map[string]interfaces.MongoDb)}

type mongoService struct {
	_Mtx sync.RWMutex
	_DBs map[string]interfaces.MongoDb
}

func (my *mongoService) Db(name string) interfaces.MongoDb {
	my._Mtx.RLock()
	db, ok := my._DBs[name]
	my._Mtx.RUnlock()
	if ok {
		return db
	}

	//根据配置创建数据库连接句柄
	my._Mtx.Lock()
	db = engines.NewMongoDb(name)
	my._DBs[name] = db
	my._Mtx.Unlock()

	return db
}

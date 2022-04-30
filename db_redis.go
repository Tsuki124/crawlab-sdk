package crawlab_sdk

import (
	"crawlab-sdk/internal/engines"
	"crawlab-sdk/internal/interfaces"
	"sync"
)

var Redis = redisService{_DBs: make(map[int]interfaces.RedisDb)}

type redisService struct {
	_Mtx sync.RWMutex
	_DBs map[int]interfaces.RedisDb
}

func (my *redisService) Db(names ...int) interfaces.RedisDb {
	//默认选择0
	name := 0
	if len(names)>0 {
		name = names[0]
	}

	my._Mtx.RLock()
	db, ok := my._DBs[name]
	my._Mtx.RUnlock()
	if ok {
		return db
	}

	//根据配置创建数据库连接句柄
	my._Mtx.Lock()
	db = engines.NewRedis(name)
	my._DBs[name] = db
	my._Mtx.Unlock()

	return db
}


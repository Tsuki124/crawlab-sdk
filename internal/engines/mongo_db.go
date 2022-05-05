package engines

import (
	"github.com/crawlab-team/go-trace"
	"github.com/qiniu/qmgo"
	"sync"
)

type MongoDb struct {
	interfaces.MongoDb
	_DB *qmgo.Database

	_Mtx sync.RWMutex
	_TBs map[string]interfaces.MongoTb
}

func NewMongoDb(name string) interfaces.MongoDb {
	db, err := driver.Mongo.New(name)
	if err != nil {
		panic(trace.TraceError(err))
	}

	return &MongoDb{_DB: db, _TBs: make(map[string]interfaces.MongoTb)}
}

func (my *MongoDb) TB(name string) interfaces.MongoTb {
	my._Mtx.RLock()
	tb, ok := my._TBs[name]
	my._Mtx.RUnlock()
	if ok {
		return tb
	}

	my._Mtx.Lock()
	tb = &MongoTb{_TB: my._DB.Collection(name)}
	my._TBs[name] = tb
	my._Mtx.Unlock()

	return tb
}

//func (my *MongoDb) _table(name string) (table interfaces.MongoTb,has bool) {
//	has = true
//	defer func() {
//		if err:=recover();err!=nil {
//			has = false
//		}
//	}()
//
//	table = my._DB.Collection(name)
//	if table==nil {
//		has = false
//	}
//
//	return table,has
//}

//func (my *MongoDb) _createTB(names ...string) error {
//	ctx := context.Background()
//	for _, name := range names {
//		err := my._DB.CreateCollection(ctx, name)
//		if err != nil {
//			return trace.Error(err)
//		}
//	}
//	return nil
//}

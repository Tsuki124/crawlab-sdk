package engines

import (
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"github.com/crawlab-team/go-trace"
	"gorm.io/gorm"
	"strings"
	"xorm.io/xorm"
)

type SQLTb struct {
	interfaces.SQLTb
	_InstanceFn func() *gorm.DB
}

func trimSQLReturnID(sql string) string {
	return strings.TrimRight(sql,`RETURNING "id"`)
}

func (my *SQLTb) _toSQLAndExecute(queryFn func(tx *gorm.DB) *gorm.DB) (string,error) {
	sql := trimSQLReturnID(my._InstanceFn().ToSQL(queryFn))
	err := queryFn(my._InstanceFn()).Error

	return sql,trace.Error(err)
}

func (my *SQLTb) _executeAndToSQL(queryFn func(tx *gorm.DB) *gorm.DB) (string,error) {
	err := queryFn(my._InstanceFn()).Error
	sql := trimSQLReturnID(my._InstanceFn().ToSQL(queryFn))

	return sql,trace.Error(err)
}

func (my *SQLTb) Insert(data interface{}) (string,error) {
	//先执行INSERT,执行成功后会将ID加入到data中
	//最后使用新data生成带有ID的SQL语句
	return my._executeAndToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(data)
	})
}

func (my *SQLTb) Delete(query interface{}, args ...interface{}) (string,error) {
	return my._toSQLAndExecute(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query,args).Delete(nil)
	})
}

func (my *SQLTb) Update(replacement interface{},query interface{}, args ...interface{}) (string,error) {
	return my._toSQLAndExecute(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query,args).Updates(replacement)
	})
}

func (my *SQLTb) Upsert(replacement interface{},query interface{}, args ...interface{}) (string,error) {
	has,err := my.Exist(query,args)
	if err!=nil {
		return "", trace.Error(err)
	}

	//数据不存在,插入数据
	if !has {
		return my._executeAndToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Create(replacement)
		})
	}

	//数据存在，更新数据
	return my._toSQLAndExecute(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query,args).Updates(replacement)
	})
}

func (my *SQLTb) FindOne(result interface{},query interface{}, args ...interface{}) (string,error) {
	return my._toSQLAndExecute(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query,args).First(result)
	})
}

func (my *SQLTb) FindALL(result interface{},query interface{}, args ...interface{}) (string,error) {
	return my._toSQLAndExecute(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query,args).Find(result)
	})
}

func (my *SQLTb) Count(query interface{}, args ...interface{}) (int64,error) {
	var count int64
	err := my._InstanceFn().Where(query,args).Count(&count).Error
	return count,trace.Error(err)
}

func (my *SQLTb) Exist(query interface{}, args ...interface{}) (bool,error) {
	err := my._InstanceFn().First(query,args).Error
	if err==nil {
		return true,nil
	}

	if err!=nil && err.Error()==gorm.ErrRecordNotFound.Error() {
		return false,nil
	}

	return false,err
}

func (my *SQLTb) UseGorm(queryFn func(tx *gorm.DB) error) error {
	return trace.Error(queryFn(my._InstanceFn()))
}

func (my *SQLTb) UseXorm(queryFn func(tx *xorm.Engine) error) error {
	//todo add the xorm
	panic("Not impl the method")
	return nil
}


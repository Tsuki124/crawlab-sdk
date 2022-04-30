package interfaces

import (
	"gorm.io/gorm"
	"xorm.io/xorm"
)

type SQLTb interface {
	Insert(data interface{}) (sql string,err error)
	Delete(query interface{}, args ...interface{}) (sql string,err error)
	Update(replacement interface{},query interface{}, args ...interface{}) (sql string,err error)
	Upsert(replacement interface{},query interface{}, args ...interface{}) (sql string,err error)
	FindOne(result interface{}, query interface{}, args ...interface{}) (sql string,err error)
	FindALL(result interface{}, query interface{}, args ...interface{}) (sql string,err error)

	Count(query interface{}, args ...interface{}) (count int64,err error)
	Exist(query interface{}, args ...interface{}) (has bool,err error)
	UseGorm(queryFn func(tx *gorm.DB) error) error
	UseXorm(queryFn func(tx *xorm.Engine) error) error
}


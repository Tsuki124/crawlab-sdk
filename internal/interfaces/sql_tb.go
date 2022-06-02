package interfaces

import "gorm.io/gorm"

type SQLTb interface {
	Exec(queryFn func(tx *gorm.DB) *gorm.DB) error //执行
	ExecToSQL(queryFn func(tx *gorm.DB) *gorm.DB) (sql string,err error) //先执行后生成SQL
	ToSQLExec(queryFn func(tx *gorm.DB) *gorm.DB) (sql string,err error) //先产生SQL后执行
}


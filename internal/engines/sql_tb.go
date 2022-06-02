package engines

import (
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"github.com/crawlab-team/go-trace"
	"gorm.io/gorm"
	"strings"
)

type SQLTb struct {
	interfaces.SQLTb
	_InstanceFn func() *gorm.DB
}

func trimSQLReturnID(sql string) string {
	return strings.TrimRight(sql,`RETURNING "id"`)
}

func (my *SQLTb) Exec(queryFn func(tx *gorm.DB) *gorm.DB) error  {
	return trace.Error(queryFn(my._InstanceFn()).Error)
}

func (my *SQLTb) ExecToSQL(queryFn func(tx *gorm.DB) *gorm.DB) (sql string,err error)   {
	err = queryFn(my._InstanceFn()).Error
	if err!=nil {
		return "",trace.Error(err)
	}

	sql = trimSQLReturnID(my._InstanceFn().ToSQL(queryFn))
	return sql,nil
}

func (my *SQLTb) ToSQLExec(queryFn func(tx *gorm.DB) *gorm.DB) (sql string,err error)   {
	sql = trimSQLReturnID(my._InstanceFn().ToSQL(queryFn))
	err = queryFn(my._InstanceFn()).Error
	return sql,trace.Error(err)
}


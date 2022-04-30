package interfaces

type SQLDb interface {
	CreateTB(models ...interface{}) error
	TB(name string) SQLTb
}

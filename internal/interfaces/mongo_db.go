package interfaces

type MongoDb interface {
	TB(name string) MongoTb
}

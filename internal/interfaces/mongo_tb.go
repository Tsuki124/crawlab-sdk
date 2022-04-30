package interfaces

type MongoTb interface {
	Insert(data interface{}) error
	Delete(condi interface{}) error
	Update(replacement, condi interface{}) error
	Upsert(replacement, condi interface{}) error
	FindOne(result interface{},filter interface{}) error
	FindALL(result interface{},filter interface{}) error
}

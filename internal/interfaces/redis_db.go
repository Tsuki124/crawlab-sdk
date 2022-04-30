package interfaces

type RedisDb interface {
	Subscribe(topic string,msgFn func(msg RedisMsg))
	Publish(topic,msg string) (count int64,err error)
}

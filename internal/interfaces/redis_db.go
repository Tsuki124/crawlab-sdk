package interfaces

type RedisDb interface {
	Subscribe(topic string,msgFn func(channel,pattern,payload string))
	Publish(topic,msg string) (count int64,err error)
}

package engines

import (
	"crawlab-sdk/internal/driver"
	"crawlab-sdk/internal/interfaces"
	"github.com/crawlab-team/go-trace"
	"github.com/go-redis/redis"
)

type RedisDb struct {
	interfaces.RedisDb
	_DB *redis.Client
}

func NewRedis(name int) interfaces.RedisDb {
	db,err := driver.Redis.New(name)
	if err!=nil {
		panic(trace.Error(err))
	}

	return &RedisDb{_DB: db}
}

func (my *RedisDb) Subscribe(topic string,msgFn func(pubsub interfaces.RedisMsg))  {
	pubsub := my._DB.Subscribe(topic)
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			trace.PrintError(err)
		}
	}(pubsub)

	for pubsubMsg := range pubsub.Channel() {
		msg := &RedisMsg{
			Channel: pubsubMsg.Channel,
			Pattern: pubsubMsg.Pattern,
			Payload: pubsubMsg.Payload,
		}
		msgFn(msg)
	}

}

func (my *RedisDb) Publish(topic,msg string) (int64,error) {
	count,err := my._DB.Publish(topic,msg).Result()
	return count,trace.Error(err)
}
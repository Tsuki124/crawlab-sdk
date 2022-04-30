package engines

type RedisMsg struct {
	Channel string
	Pattern string
	Payload string
}

func (my *RedisMsg) GetChannel() string {
	return my.Channel
}

func (my *RedisMsg) GetPattern() string {
	return my.Pattern
}

func (my *RedisMsg) GetPayload() string {
	return my.Payload
}

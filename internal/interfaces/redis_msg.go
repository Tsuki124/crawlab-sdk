package interfaces

type RedisMsg interface {
	GetChannel() string
	GetPattern() string
	GetPayload() string
}

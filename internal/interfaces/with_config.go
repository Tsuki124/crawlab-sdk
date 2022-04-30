package interfaces

type WithConfig interface {
	GetConfigMap() map[string]string
}

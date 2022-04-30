package interfaces

type SeaweedFile interface {
	Path() string
	Name() string
	IsDir() bool
}

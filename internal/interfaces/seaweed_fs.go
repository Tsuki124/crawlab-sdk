package interfaces

type SeaweedFS interface {
	List(dirpath string) ([]SeaweedFile,error)
	Download(path string) ([]byte,error)
	Upload(path string,content []byte) error
	Delete(path string) error
	Info(path string) (SeaweedFile,error)
}

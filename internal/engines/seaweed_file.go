package engines

type SeaweedFile struct {
	name string
	path string
	isDir bool
}

func (my *SeaweedFile) Path() string {
	return my.path
}

func (my *SeaweedFile) Name() string {
	return my.name
}

func (my *SeaweedFile) IsDir() bool {
	return my.isDir
}


package file

type File struct {
	Name                    string
	AccessLevelBiba         int
	AccessLevelBellLaPadula int
}

//var Files = make([]file,5)

func NewFile(name string, accessLevelBiba int, accessLevelBellLapadula int) *File {
	var newFile = &File{
		Name:                    name,
		AccessLevelBiba:         accessLevelBiba,
		AccessLevelBellLaPadula: accessLevelBellLapadula,
	}
	return newFile
}

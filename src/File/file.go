package file

type FileExtension string

const (
	Txt     FileExtension = ".txt"
	Csv     FileExtension = ".csv"
	Excel   FileExtension = ".xlsx"
	Json    FileExtension = ".json"
	Graphql FileExtension = ".graphql"
	Sql     FileExtension = ".sql"
	Xml     FileExtension = ".xml"
	Zip     FileExtension = ".zip"
)

type File struct {
	name      string
	path      string
	extension FileExtension
}

func NewFile(name string, path string, extension FileExtension) (*File, error) {
	f := new(File)
	f.SetName(name)
	f.SetPath(path)
	f.SetExtension(extension)
	return f, nil
}

func (this File) SetName(name string) {
	this.name = name
}

func (this File) SetPath(path string) {
	this.path = path
}

func (this File) SetExtension(extension FileExtension) {
	this.extension = extension
}

func (this File) GetName() string {
	return this.name
}

func (this File) GetPath() string {
	return this.path
}

func (this File) GetExtension() string {
	return string(this.extension)
}

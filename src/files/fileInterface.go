package files

// FileInterface is a File interface for writing and reading Files
type FileInterface interface {
	New(string, string, FileExtension) *interface{}
	Create() error
	Read() error
	WriteReplace(bool) error
	WriteReplaceTo(string, bool) error
	WriteAppend(bool) error
	WriteAppendTo(string, bool) error
	Delete() error
	Copy(string) error
	Move(FileInterface) error
	Print()
	Clone() *FileInterface
	GetName() string
	GetPath() string
	GetExtension() FileExtension
	GetData() interface{}
	GetFullPath() string
	SetName(string)
	SetPath(string)
	SetExtension(FileExtension)
	SetData(interface{})
}

// File abstract struct for FileInterface
type File struct {
	name string
	path string
	FileExtension
}

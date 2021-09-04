package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Concrete File struct for txt files.
type TxtFile struct {
	file File
	data []byte
}

// New creates a new File with the given name, path, and extension.
func New(name string, path string, extension FileExtension) *TxtFile {
	f := new(TxtFile)
	f.SetName(name)
	f.SetPath(path)
	f.SetExtension(extension)
	f.SetData(nil)
	return f
}

func (txtFile *TxtFile) Create() error { // TODO
	if _, err := os.Stat(txtFile.GetPath()); err == nil {
		if _, err := os.Stat(txtFile.GetFullPath()); err == nil {
			// file exists

		} else if os.IsNotExist(err) {
			// file don't exist

		} else {
			// file may or may not exist. See err for details.

		}
	} else if os.IsNotExist(err) {
		// parent path doesn't exist

	} else {
		// parent path may or may not exist. See err for details.

	}
	return nil
}

// Read reads a file at the File's path and stores the data.
func (txtFile *TxtFile) Read() error {
	data, err := ioutil.ReadFile(txtFile.GetPath())
	if err != nil {
		return err
	}
	txtFile.data = data
	return nil
}

func (txtFile *TxtFile) WriteReplace() error // TODO

func (txtFile *TxtFile) WriteReplaceTo(path string) error // TODO

func (txtFile *TxtFile) WriteAppend() error // TODO

func (txtFile *TxtFile) WriteAppendTo(path string) error // TODO

// Delete deletes the file at the File's path.
func (txtFile *TxtFile) Delete() error

// Copy copies the file at the File's path to the given path.
func (txtFile *TxtFile) Copy(string) error

// Move moves the file at the File's path to the given path.
func (txtFile *TxtFile) Move(string) error

// Print prints the File's data to the standard output.
func (txtFile *TxtFile) Print() {
	fmt.Println(txtFile.GetData())
}

// Clone clones the file at the File's path to the given path. // TODO
func (txtFile *TxtFile) Clone() (TxtFile, error)

// GetName returns the name of this file.
func (txtFile *TxtFile) GetName() string {
	return txtFile.file.name
}

// GetPath returns the path of this File.
func (txtFile *TxtFile) GetPath() string {
	return txtFile.file.path
}

// GetExtension gets the FileExtension of this File.
func (txtFile *TxtFile) GetExtension() FileExtension {
	return txtFile.file.extension
}

// GetData returns the File's data.
func (txtFile *TxtFile) GetData() []byte {
	return txtFile.data
}

// GetFullPath gets the File's path + name + extension.
func (txtFile *TxtFile) GetFullPath() string {
	return txtFile.GetPath() + txtFile.GetName() + string(txtFile.GetExtension())
}

// SetName sets this File's name with a given string.
func (txtFile *TxtFile) SetName(name string) {
	txtFile.file.name = name
}

// SetPath sets this File's path with a given string.
func (txtFile *TxtFile) SetPath(path string) {
	txtFile.file.path = path
}

// SetExtension sets this File's extension with a given FileExtension.
func (txtFile *TxtFile) SetExtension(extension FileExtension) {
	txtFile.file.extension = extension
}

// SetData replaces the File's data with the given data.
func (txtFile *TxtFile) SetData(data []byte) error {
	txtFile.data = data
	return nil
}

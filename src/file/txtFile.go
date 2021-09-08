package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Concrete File struct for txt files.
type TxtFile struct {
	file File
	data []byte
	FileInterface
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

// Create creates a new File at the File's path if it doesn't already exist.
func (txtFile *TxtFile) Create() error {
	if _, err := os.Stat(txtFile.GetPath()); err == nil {
		if _, err := os.Stat(txtFile.GetFullPath()); err == nil { // file exists
			return nil
		} else if os.IsNotExist(err) { // file don't exist
			_, err := os.Create(txtFile.GetFullPath())
			return err
		} else { // file may or may not exist. See err for details.
			return err
		}
	} else if os.IsNotExist(err) { // parent path doesn't exist
		err := os.Mkdir(txtFile.GetPath(), 0644)
		if err != nil {
			return err
		}
		_, err1 := os.Create(txtFile.GetFullPath())
		return err1
	} else { // parent path may or may not exist. See err for details.
		return err
	}
}

// Read reads a file at the File's path and stores the data.
func (txtFile *TxtFile) Read() error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	data, err1 := ioutil.ReadFile(txtFile.GetPath())
	if err1 != nil {
		return err1
	}
	txtFile.data = data
	return nil
}

// WriteReplace writes the File's data to a file at the File's path.
func (txtFile *TxtFile) WriteReplace() error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	err1 := txtFile.WriteReplaceTo(txtFile.GetFullPath())
	return err1
}

// WriteReplaceTo writes the File's data to a file at the given path.
func (txtFile *TxtFile) WriteReplaceTo(path string) error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	err1 := os.WriteFile(path, txtFile.data, 0644)
	return err1
}

// WriteAppend appends the File's data to a file at the File's path.
func (txtFile *TxtFile) WriteAppend() error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(txtFile.GetFullPath(), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err1 := f.Write(txtFile.GetData())
	return err1
}

// WriteAppendTo appends the File's data to a file at the given path.
func (txtFile *TxtFile) WriteAppendTo(path string) error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	f, err1 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		return err1
	}
	defer f.Close()
	_, err2 := f.Write(txtFile.GetData())
	return err2
}

// Delete deletes a file at the File's path.
func (txtFile *TxtFile) Delete() error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	err1 := os.Remove(txtFile.GetFullPath())
	return err1
}

// Copy copies the file at the File's path to the given path.
func (txtFile *TxtFile) Copy(path string) error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	fSource, err1 := os.Open(txtFile.GetFullPath())
	if err1 != nil {
		return err1
	}
	defer fSource.Close()
	fDestination, err2 := os.Create(path)
	if err2 != nil {
		return err2
	}
	defer fDestination.Close()
	_, err3 := io.Copy(fDestination, fSource)
	return err3
}

// Move moves the file at the File's path to the given path.
func (txtFile *TxtFile) Move(path string) error {
	err := txtFile.checkFile()
	if err != nil {
		return err
	}
	err1 := os.Rename(txtFile.GetFullPath(), path)
	return err1
}

// Print prints the File's data to the standard output.
func (txtFile *TxtFile) Print() {
	fmt.Println(txtFile.GetData())
}

// Clone returns an identical copy of the File.
func (txtFile *TxtFile) Clone() *TxtFile {
	clone := New(txtFile.GetName(), txtFile.GetPath(), txtFile.GetExtension())
	clone.SetData(txtFile.GetData())
	return clone
}

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

// checkFile returns an error if the file is not regular or doesn't exist.
func (txtFile *TxtFile) checkFile() error {
	fileStat, err := os.Stat(txtFile.GetFullPath())
	if err != nil {
		return err
	}
	if !fileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", txtFile.GetFullPath())
	}
	return nil
}

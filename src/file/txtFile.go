package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// TxtFile is a File that contains txt data.
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
			return os.ErrExist
		} else if os.IsNotExist(err) { // file doesn't exist
			if _, err := os.Create(txtFile.GetFullPath()); err == nil {
				return txtFile.WriteReplace()
			}
			return err
		} else { // file may or may not exist. See err for details.
			return err
		}
	} else if os.IsNotExist(err) { // parent path doesn't exist
		if err := os.Mkdir(txtFile.GetPath(), 0644); err != nil {
			return err
		}
		if _, err := os.Create(txtFile.GetFullPath()); err != nil {
			return err
		}
		if txtFile.GetData() != nil {
			err := txtFile.WriteReplace()
			return err
		}
		return nil
	} else { // parent path may or may not exist. See err for details.
		return err
	}
}

// Read reads a file at the File's path and stores the data.
func (txtFile *TxtFile) Read() error {
	err := txtFile.CheckFile()
	if err != nil {
		return err
	}
	data, err1 := ioutil.ReadFile(txtFile.GetFullPath())
	if err1 != nil {
		return err1
	}
	txtFile.data = data
	return nil
}

// WriteReplace writes the File's data to a file at the File's path.
//This method is a wrapper of WriteReplaceTo(txtFile.GetFullPath).
func (txtFile *TxtFile) WriteReplace() error {
	err := txtFile.WriteReplaceTo(txtFile.GetFullPath())
	return err
}

// WriteReplaceTo writes the File's data to a file at the given path.
//This method doesn't append the File's extension at the end of the specified
//path.
func (txtFile *TxtFile) WriteReplaceTo(path string) error {
	if err := txtFile.CheckFile(); os.IsNotExist(err) {
		if err := txtFile.Create(); err != nil {
			return err
		}
		if err := txtFile.WriteAppendTo(path); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	err := os.WriteFile(path, txtFile.data, 0644)
	return err
}

// WriteAppend appends the File's data in a new line to a file at the File's path.
// This method is a wrapper of WriteAppendTo(txtFile.GetFullPath).
func (txtFile *TxtFile) WriteAppend() error {
	err := txtFile.WriteAppendTo(txtFile.GetFullPath())
	return err
}

// WriteAppendTo appends the File's data in a new line to a file at the given path.
func (txtFile *TxtFile) WriteAppendTo(path string) error {
	if err := txtFile.CheckFile(); err != nil {
		return err
	}
	f, err1 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		return err1
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	aux, err1 := ioutil.ReadFile(txtFile.GetFullPath())
	if err1 != nil {
		return err1
	}
	if aux != nil {

		txtFile.SetData(append([]byte("\n"), txtFile.GetData()...))
	}
	_, err2 := f.Write(txtFile.GetData())
	return err2
}

// Delete deletes a file at the File's path.
func (txtFile *TxtFile) Delete() error {
	err := txtFile.CheckFile()
	if err != nil {
		return err
	}
	err1 := os.Remove(txtFile.GetFullPath())
	return err1
}

// Copy copies the file at the File's path to the given path.
func (txtFile *TxtFile) Copy(path string) error {
	err := txtFile.CheckFile()
	if err != nil {
		return err
	}
	fSource, err1 := os.Open(txtFile.GetFullPath())
	if err1 != nil {
		return err1
	}
	defer func(fSource *os.File) {
		_ = fSource.Close()
	}(fSource)
	fDestination, err2 := os.Create(path)
	if err2 != nil {
		return err2
	}
	defer func(fDestination *os.File) {
		_ = fDestination.Close()
	}(fDestination)
	_, err3 := io.Copy(fDestination, fSource)
	return err3
}

// Move moves the file at the File's path to the given path.
func (txtFile *TxtFile) Move(path string) error {
	err := txtFile.CheckFile()
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
func (txtFile *TxtFile) Clone() TxtFile {
	clone := New(txtFile.GetName(), txtFile.GetPath(), txtFile.GetExtension())
	clone.SetData(txtFile.GetData())
	return *clone
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
func (txtFile *TxtFile) SetData(data []byte) {
	txtFile.data = data
}

// CheckFile returns an error if the file is not regular or doesn't exist.
func (txtFile *TxtFile) CheckFile() error {
	fileStat, err := os.Stat(txtFile.GetFullPath())
	if err != nil {
		return err
	}
	if !fileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", txtFile.GetFullPath())
	}
	return nil
}

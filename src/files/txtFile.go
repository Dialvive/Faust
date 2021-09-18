package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// TxtFile is a File that contains txt data.
type TxtFile struct {
	File
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

// Create creates a new File at the File's path if it doesn't already exist.
func (file *TxtFile) Create() error {
	if _, err := os.Stat(file.GetPath()); err == nil {
		if _, err := os.Stat(file.GetFullPath()); err == nil { // file exists
			return os.ErrExist
		} else if os.IsNotExist(err) { // file doesn't exist
			if _, err := os.Create(file.GetFullPath()); err == nil {
				return file.WriteReplace()
			}
			return err
		} else { // file may or may not exist. See err for details.
			return err
		}
	} else if os.IsNotExist(err) { // parent path doesn't exist
		if err := os.Mkdir(file.GetPath(), 0644); err != nil {
			return err
		}
		if _, err := os.Create(file.GetFullPath()); err != nil {
			return err
		}
		if file.GetData() != nil {
			err := file.WriteReplace()
			return err
		}
		return nil
	} else { // parent path may or may not exist. See err for details.
		return err
	}
}

// Read reads a file at the File's path and stores the data.
func (file *TxtFile) Read() error {
	err := file.CheckFile()
	if err != nil {
		return err
	}
	data, err1 := ioutil.ReadFile(file.GetFullPath())
	if err1 != nil {
		return err1
	}
	file.data = data
	return nil
}

// WriteReplace writes the File's data to a file at the File's path.
//This method is a wrapper of WriteReplaceTo(file.GetFullPath).
func (file *TxtFile) WriteReplace() error {
	err := file.WriteReplaceTo(file.GetFullPath())
	return err
}

// WriteReplaceTo writes the File's data to a file at the given path.
//This method doesn't append the File's extension at the end of the specified
//path.
func (file *TxtFile) WriteReplaceTo(path string) error {
	if err := file.CheckFile(); os.IsNotExist(err) {
		if err := file.Create(); err != nil {
			return err
		}
		if err := file.WriteAppendTo(path); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	err := os.WriteFile(path, file.data, 0644)
	return err
}

// WriteAppend appends the File's data and a new line to a file at the File's path.
// This method is a wrapper of WriteAppendTo(file.GetFullPath).
func (file *TxtFile) WriteAppend() error {
	err := file.WriteAppendTo(file.GetFullPath())
	return err
}

// WriteAppendTo appends the File's data and a new line to a file at the given path.
func (file *TxtFile) WriteAppendTo(path string) error {
	if err := file.CheckFile(); err != nil {
		return err
	}
	f, err1 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err1 != nil {
		return err1
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	_, err2 := f.Write(append(file.GetData(), []byte("\n")...))
	return err2
}

// Delete deletes a file at the File's path.
func (file *TxtFile) Delete() error {
	err := file.CheckFile()
	if err != nil {
		return err
	}
	err1 := os.Remove(file.GetFullPath())
	return err1
}

// Copy copies the file at the File's path to the given path.
func (file *TxtFile) Copy(path string) error {
	err := file.CheckFile()
	if err != nil {
		return err
	}
	fSource, err1 := os.Open(file.GetFullPath())
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

// Move moves the file at this File's path to the given File's path and name.
// Mind that this method keeps the same FileExtension.
func (file *TxtFile) Move(newFile TxtFile) error {
	err := file.CheckFile()
	if err != nil {
		return err
	}
	err1 := os.Rename(file.GetFullPath(), newFile.GetFullPath())
	return err1
}

// Print prints the File's data to the standard output.
func (file *TxtFile) Print() {
	fmt.Println(file.GetData())
}

// Clone returns an identical copy of the File.
func (file *TxtFile) Clone() *TxtFile {
	clone := New(file.GetName(), file.GetPath(), file.GetExtension())
	clone.SetData(file.GetData())
	return clone
}

// GetName returns the name of this file.
func (file *TxtFile) GetName() string {
	return file.name
}

// GetPath returns the path of this File.
func (file *TxtFile) GetPath() string {
	return file.path
}

// GetExtension gets the FileExtension of this File.
func (file *TxtFile) GetExtension() FileExtension {
	return file.FileExtension
}

// GetData returns the File's data.
func (file *TxtFile) GetData() []byte {
	return file.data
}

// GetFullPath gets the File's path + name + extension.
func (file *TxtFile) GetFullPath() string {
	return file.GetPath() + file.GetName() + string(file.GetExtension())
}

// SetName sets this File's name with a given string.
func (file *TxtFile) SetName(name string) {
	file.name = name
}

// SetPath sets this File's path with a given string.
func (file *TxtFile) SetPath(path string) {
	file.path = path
}

// SetExtension sets this File's extension with a given FileExtension.
func (file *TxtFile) SetExtension(extension FileExtension) {
	file.FileExtension = extension
}

// SetData replaces the File's data with the given data.
func (file *TxtFile) SetData(data []byte) {
	file.data = data
}

// CheckFile returns an error if the file is not regular or doesn't exist.
func (file *TxtFile) CheckFile() error {
	fileStat, err := os.Stat(file.GetFullPath())
	if err != nil {
		return err
	}
	if !fileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", file.GetFullPath())
	}
	return nil
}


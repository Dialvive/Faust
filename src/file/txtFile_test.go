package file

import (
	"os"
	"reflect"
	"testing"
)

func TestTxtFile(t *testing.T) {

	// Create test directory
	var Dir string = "./../../test/"
	err := os.Mkdir(Dir, 0755)
	if err != nil {
		t.Log(err)
	}

	t.Run("TxtGetName", func(t *testing.T) {
		file := New("TxtGetName", Dir, Txt)
		if file.GetName() != "TxtGetName" {
			t.Log("TxtGetName FAILED")
			t.Fail()
		}
	})

	t.Run("TxtGetPath", func(t *testing.T) {
		file := New("TxtGetPath", Dir, Txt)
		if file.GetPath() != Dir {
			t.Log("TxtGetPath FAILED")
			t.Fail()
		}
	})

	t.Run("TxtGetExtension", func(t *testing.T) {
		file := New("TxtGetExtension", Dir, Txt)
		if file.GetExtension() != Txt {
			t.Log("TxtGetExtension FAILED")
			t.Fail()
		}
	})

	t.Run("TxtGetFullPath", func(t *testing.T) {
		file := New("TxtGetFullPath", Dir, Txt)
		if file.GetFullPath() != Dir+"TxtGetFullPath"+string(file.GetExtension()) {
			t.Log("TxtGetFullPath FAILED: " + file.GetFullPath())
			t.Fail()
		}
	})

	t.Run("TxtGetData", func(t *testing.T) {
		file := New("TxtGetData", Dir, Txt)
		file.data = []byte("TEST")
		if !reflect.DeepEqual(file.GetData(), []byte("TEST")) {
			t.Log("TxtGetData FAILED: " + string(file.GetData()))
			t.Fail()
		}
	})

	t.Run("TxtSetName", func(t *testing.T) {
		file := New("Wrong", Dir, Txt)
		file.SetName("TxtGetName")
		if file.GetName() != "TxtGetName" {
			t.Log("TxtGetName FAILED")
			t.Fail()
		}
	})

	t.Run("TxtSetPath", func(t *testing.T) {
		file := New("TxtSetPath", "TxtSetPath", Txt)
		file.SetPath(Dir)
		if file.GetPath() != Dir {
			t.Log("TxtSetPath FAILED")
			t.Fail()
		}
	})

	t.Run("TxtSetExtension", func(t *testing.T) {
		file := New("TxtSetExtension", "TxtSetPath", Xml)
		file.SetExtension(Txt)
		if file.GetExtension() != Txt {
			t.Log("TxtSetPath FAILED")
			t.Fail()
		}
	})

	t.Run("TxtSetData", func(t *testing.T) {
		file := New("TxtGetData", Dir, Txt)
		file.SetData([]byte("TEST"))
		if !reflect.DeepEqual(file.GetData(), []byte("TEST")) {
			t.Log("TxtGetFullPath FAILED: " + file.GetFullPath())
			t.Fail()
		}
	})

	t.Run("TxtCreate", func(t *testing.T) {
		var file TxtFile = *New("TxtCreate", Dir, Txt)
		file.SetData([]byte("TEST"))
		if err := file.Create(); err == nil {
			file.SetData([]byte("ERROR"))
			if err := file.Create(); err == nil { // Try rewritting file
				t.Log("Rewritted file")
				t.Fail()
			} else if os.IsExist(err) {
				return
			}
		} else {
			t.Log("Error while creating file: " + err.Error())
			t.Fail()
		}
	})

	t.Run("TxtRead", func(t *testing.T) {
		file := New("TxtCreate", Dir, Txt)
		if err := file.Read(); err != nil {
			t.Log("TxtRead FAILED: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("TEST")) {
			t.Log("TxtRead FAILED: " + err.Error())
			t.Fail()
		}
	})

	t.Run("TxtWriteReplace", func(t *testing.T) {
		var file TxtFile = *New("TxtCreate", Dir, Txt)
		file.SetData([]byte("REWRITTEN"))
		if err := file.WriteReplace(); err == nil {
			file.SetData([]byte("ERROR"))
			if err := file.Read(); err != nil {
				if !reflect.DeepEqual(file.GetData(), []byte("REWRITTEN")) {
					t.Log("TxtWriteReplace FAILED: " + err.Error())
					t.Fail()
				}
			}
		} else {
			t.Log("TxtWriteReplace FAILED: " + err.Error())
			t.Fail()
		}
	})

	t.Run("TxtWriteReplaceTo", func(t *testing.T) {

	})

	t.Run("TxtWriteAppend", func(t *testing.T) {

	})

	t.Run("TxtWriteAppendTo", func(t *testing.T) {

	})

	t.Run("TxtDelete", func(t *testing.T) {

	})

	t.Run("TxtCopy", func(t *testing.T) {

	})

	t.Run("TxtMove", func(t *testing.T) {

	})

	t.Run("TxtClone", func(t *testing.T) {

	})
}

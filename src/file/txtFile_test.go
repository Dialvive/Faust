package file

import (
	"os"
	"reflect"
	"testing"
)

func TestTxtFile(t *testing.T) {

	// Create test directory
	var Dir string = "./../../test/txt/"
	if err := os.Mkdir(Dir, 0755); err != nil {
		t.Log("Error initializing Test 1: " + err.Error())
		t.Fail()
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

	// Expected file: ./../../test/txt/TxtCreate.txt
	// With content: "SUCCESS"
	t.Run("TxtCreate", func(t *testing.T) {
		var file TxtFile = *New("TxtCreate", Dir, Txt)
		file.SetData([]byte("SUCCESS"))
		if err := file.Create(); err != nil {
			t.Log("TxtCreate FAILED 0: " + err.Error())
			t.Fail()
		}
		file.SetData([]byte("ERROR 0"))
		if err := file.Create(); err == nil { // Try rewriting file
			t.Log("TxtCreate FAILED 2: " + err.Error())
			t.Fail()
		}
	})

	t.Run("TxtRead", func(t *testing.T) {
		file := New("TxtCreate", Dir, Txt)
		if err := file.Read(); err != nil {
			t.Log("TxtRead FAILED: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("SUCCESS")) {
			t.Log("TxtRead FAILED: not equal data")
			t.Fail()
		}
	})

	// Expected file: ./../../test/txt/TxtWriteReplaceTo.txt
	// With content: "SUCCESS"
	t.Run("TxtWriteReplaceTo", func(t *testing.T) {
		var file TxtFile = *New("TxtWriteReplaceTo", Dir, Txt)
		file.SetData([]byte("ERROR"))
		if err := file.WriteReplaceTo(file.GetFullPath()); err != nil {
			t.Log("TxtWriteReplace FAILED 0: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteReplace FAILED 1: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("ERROR")) {
			t.Log("TxtWriteReplace FAILED, different data 0: " + string(file.GetData()))
			t.Fail()
		}
		file.SetData([]byte("SUCCESS"))
		if err := file.WriteReplaceTo(Dir + "TxtWriteReplaceTo.txt"); err != nil {
			t.Log("TxtWriteReplace FAILED 2: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("SUCCESS")) {
			t.Log("TxtWriteReplace FAILED, different data 1: " + string(file.GetData()))
			t.Fail()
		}
	})

	// Expected file: ./../../test/txt/TxtWriteReplace.txt
	// With content: "SUCCESS"
	t.Run("TxtWriteReplace", func(t *testing.T) {
		var file TxtFile = *New("TxtWriteReplace", Dir, Txt)
		file.SetData([]byte("ERROR"))
		if err := file.WriteReplace(); err != nil {
			t.Log("TxtWriteReplace FAILED 0: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteReplace FAILED 1: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("ERROR")) {
			t.Log("TxtWriteReplace FAILED, different data 0: " + string(file.GetData()))
			t.Fail()
		}
		file.SetData([]byte("SUCCESS"))
		if err := file.WriteReplace(); err != nil {
			t.Log("TxtWriteReplace FAILED 2: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("SUCCESS")) {
			t.Log("TxtWriteReplace FAILED, different data 1: " + string(file.GetData()))
			t.Fail()
		}
	})

	// Expected file: ./../../test/txt/TxtWriteAppendTo.txt
	// With content: "SUCCESS"
	t.Run("TxtWriteAppendTo", func(t *testing.T) {
		var file TxtFile = *New("TxtWriteAppendTo", Dir, Txt)
		if err := file.Create(); err != nil {
			t.Log("TxtWriteAppendTo FAILED 0: " + err.Error())
			t.Fail()
		}
		file.SetData([]byte("TxtWriteAppendTo"))
		if err := file.WriteAppendTo(file.GetFullPath()); err != nil {
			t.Log("TxtWriteAppendTo FAILED 1: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteAppendTo FAILED 2: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("TxtWriteAppendTo")) {
			t.Log("TxtWriteAppendTo FAILED, different data 0: " + string(file.GetData()))
			t.Fail()
		}
		file.SetData([]byte("SUCCESS"))
		if err := file.WriteAppendTo(file.GetFullPath()); err != nil {
			t.Log("TxtWriteAppendTo FAILED 3: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteAppend FAILED 4: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("TxtWriteAppendTo\nSUCCESS")) {
			t.Log("TxtWriteAppendTo FAILED, different data 1: " + string(file.GetData()))
			t.Fail()
		}
	})

	// Expected file: ./../../test/txt/TxtWriteAppend.txt
	// With content: "SUCCESS"
	t.Run("TxtWriteAppend", func(t *testing.T) {
		var file TxtFile = *New("TxtWriteAppend", Dir, Txt)
		if err := file.Create(); err != nil {
			t.Log("TxtWriteAppend FAILED 0: " + err.Error())
			t.Fail()
		}
		file.SetData([]byte("TxtWriteAppend"))
		if err := file.WriteAppend(); err != nil {
			t.Log("TxtWriteAppend FAILED 1: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteAppend FAILED 2: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("TxtWriteAppend")) {
			t.Log("TxtWriteAppend FAILED, different data 0: " + string(file.GetData()))
			t.Fail()
		}
		file.SetData([]byte("SUCCESS"))
		if err := file.WriteAppend(); err != nil {
			t.Log("TxtWriteAppend FAILED 3: " + err.Error())
			t.Fail()
		}
		if err := file.Read(); err != nil {
			t.Log("TxtWriteAppend FAILED 4: " + err.Error())
			t.Fail()
		}
		if !reflect.DeepEqual(file.GetData(), []byte("TxtWriteAppend\nSUCCESS")) {
			t.Log("TxtWriteAppend FAILED, different data 1: " + string(file.GetData()))
			t.Fail()
		}
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

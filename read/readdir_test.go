package read

import (
	"FileCleaner/model"
	"testing"
)

func TestReadDir(t *testing.T) {
	/*	path,err:=os.Getwd()
		if err != nil {
			t.Log(err)
		}
		t.Log("path:",path)*/
	path := "E:\\Golang\\src\\FileCleaner\\testDir"
	ReadDir(path)
	t.Log("files:", model.Files)
}

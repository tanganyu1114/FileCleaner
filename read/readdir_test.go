package read

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	/*	path,err:=os.Getwd()
		if err != nil {
			t.Log(err)
		}
		t.Log("path:",path)*/
	dir := "E:\\Golang\\src\\FileCleaner\\testDir"
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Log(err)
	}
	for _, file := range fileinfo {
		fmt.Println("name:", file.Name(), "mode:", file.Mode(), "sys:", file.Sys())
		if !file.IsDir() {
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				t.Log(err)
			}
			t.Log(file.Name(), " hash:", fileHash(data))
		}
	}
}

func TestEchoFloat(t *testing.T) {
	fmt.Printf("aaaaaaaaaaaaa\taaaaaa\taaaa\t\n")
	fmt.Printf("bbbb\tbbbb\tbbbbbbbbbbbbbbbbbbb\t\n")
	fmt.Printf("c\tcc\tccc\t\n")
}

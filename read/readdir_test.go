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
	fmt.Printf("%.2f ", 0.15)
	var aa int = 1024
	var bb int = 1010
	fmt.Println(float64(aa) / 1024 / 1024)
	fmt.Println(float64(aa / bb))
	numPct := fmt.Sprintf("%.2f %%", float64(aa)/float64(bb))
	fmt.Println(numPct)
}

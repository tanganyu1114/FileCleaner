package read

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"syscall"
	"testing"
	"time"
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
		fmt.Println("name:", file.Name(), "mode:", file.Mode(), "sys:", file.Sys(), "modetime", file.ModTime())
		stat_t := file.Sys().(*syscall.Win32FileAttributeData)
		time.Now().Unix()
		fmt.Println("ctime", time.Unix(stat_t.CreationTime.Nanoseconds()/1e9, 0))
		fmt.Println(stat_t.CreationTime.Nanoseconds() / 1e9)
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
	ss := "abcd"
	fmt.Println(ss[0:1])
	day := "1"
	iday, err := strconv.Atoi(day)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(iday)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().AddDate(0, 0, -1).Unix())
	fmt.Println(time.Now().AddDate(0, 0, -2).Unix())
	fmt.Println(time.Now().AddDate(0, 0, -1).Unix() - time.Now().AddDate(0, 0, -2).Unix())
	fmt.Println(time.Now().Unix() - time.Now().Unix()%86400)
}

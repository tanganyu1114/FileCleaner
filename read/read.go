package read

import (
	"FileCleaner/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

// 纯文件信息
var files []string

// 读进程获取文件信息
func Read(basePath string, casCade bool) {
	// 单线程处理读目录信息
	if casCade {
		readDirMulti(basePath)
	} else {
		readDirSingle(basePath)
	}
	// goroutine 并发处理文件
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		model.ControlCH <- 1
		go func(file string) {
			defer wg.Done()
			defer func() {
				<-model.ControlCH
			}()
			var (
				stat   = true
				md5str string
			)
			data, err := ioutil.ReadFile(file)
			if err != nil {
				stat = false
				fmt.Printf("Read file err, file name :%s, err: %s", file, err.Error())
			} else {
				md5str = fileHash(data)
			}
			rd := model.NewRead(stat, md5str, file, len(data))
			model.RecordCH <- rd
		}(file)
	}
	wg.Wait()
	// 阻塞，等待record处理完成数据
	fmt.Printf("Wait the record to statistics the infomation .")
	for {
		if len(model.RecordCH) == 0 {
			fmt.Printf("\nRecord Read file info Complete\n")
			return
		}
		fmt.Printf(".")
		time.Sleep(time.Second)
	}
}

// 级联读取目录文件信息
func readDirMulti(dir string) {
	var stat = true
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		stat = false
		fmt.Printf("Read file path err, file path %s, err : %s", dir, err.Error())
	}
	rd := model.NewPath(stat, dir)
	model.RecordCH <- rd
	for _, file := range fileinfo {
		if file.IsDir() {
			readDirMulti(filepath.Join(dir, file.Name()))
		} else {
			files = append(files, filepath.Join(dir, file.Name()))
		}
	}
}

// 只读单个目录，非级联读取
func readDirSingle(dir string) {
	var stat = true
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		stat = false
		fmt.Printf("Read file path err, file path %s, err : %s", dir, err.Error())
	}
	rd := model.NewPath(stat, dir)
	model.RecordCH <- rd
	for _, file := range fileinfo {
		if file.IsDir() {
			continue
		}
		files = append(files, filepath.Join(dir, file.Name()))
	}
}

func fileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

package read

import (
	"FileCleaner/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sync"
)

// 读进程获取文件信息
func Read(basePath string, casCade bool) {
	if casCade {
		ReadDir(basePath)
	} else {
		fileinfo, err := ioutil.ReadDir(basePath)
		if err != nil {
			fmt.Printf("read file path err, file path %s, err : %s", basePath, err.Error())
		}
		for _, file := range fileinfo {
			if file.IsDir() {
				continue
			}
			model.Files = append(model.Files, basePath+"/"+file.Name())
		}
	}
	wg := sync.WaitGroup{}
	for _, file := range model.Files {
		wg.Add(1)
		model.ReadCH <- 1
		go func(file string) {
			defer wg.Done()
			data, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("read file err, file name :%s, err: %s", file, err.Error())
			}
			md5str := FileHash(data)
			model.FileMap[md5str] = append(model.FileMap[md5str], file)
			<-model.ReadCH
		}(file)
	}
	wg.Wait()

	// 所有文件读取完毕后，发送消息给写进程
	model.SignalCH <- true
}

func ReadDir(dir string) {
	fileinfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("read file path err, file path %s, err : %s", dir, err.Error())
	}
	for _, file := range fileinfo {
		if file.IsDir() {
			ReadDir(dir + "/" + file.Name())
		} else {
			model.Files = append(model.Files, dir+"/"+file.Name())
		}
	}
}

func FileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

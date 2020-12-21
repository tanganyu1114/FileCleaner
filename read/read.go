package read

import (
	"FileCleaner/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sync"
)

// 后台记录记录文件信息
func Record() {
	for {
		select {
		case record := <-model.RecordCH:
			for k, v := range record {
				if CheckFileMap(k) {
					model.FileMap[k] = append(model.FileMap[k], v)
				} else {
					model.FileMap[k] = []string{v}
				}
			}
		// 接收退出信息
		case <-model.SignalCH:
			return
		}

	}
}

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
			//if CheckFileMap(md5str) {
			//	model.FileMap[md5str] = append(model.FileMap[md5str], file)
			//} else {
			//	model.FileMap[md5str] = []string{file}
			//}
			record := map[string]string{
				md5str: file,
			}
			model.RecordCH <- record
			<-model.ReadCH
		}(file)
	}
	wg.Wait()
	// 关闭读文件限制通道
	close(model.ReadCH)
	// 阻塞，等待record处理完成数据，发送信息退出record goroutine
	for {
		if len(model.RecordCH) == 0 {
			close(model.RecordCH)
			model.SignalCH <- true
			return
		}
	}
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

func CheckFileMap(md5 string) bool {
	_, ok := model.FileMap[md5]
	return ok
}

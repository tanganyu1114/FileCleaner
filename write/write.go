package write

import (
	"FileCleaner/model"
	"fmt"
	"os"
	"sync"
)

// 在这个包主要操作文件，删除文件创建ln

func Write(dm string) {
	// 写阻塞，等待读取完毕
	<-model.SignalCH
	// 删除文件
	RemoveFile()
	if dm == "ln" {
		CreateLink()
	}
}

func RemoveFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.WriteCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			if len(files) > 1 {
				for _, file := range files[1:] {
					fmt.Printf("Remove the File hash: %s", hash)
					fmt.Printf("List the Remove file :")
					err := os.Remove(file)
					if err != nil {
						fmt.Printf("remove file: %s ERROR : %s", file, err.Error())
					} else {
						fmt.Printf("RM: %s", file)
					}
				}
			}
			<-model.WriteCH
		}(hash, files)
	}
	wg.Wait()
}

func CreateLink() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.WriteCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			if len(files) > 1 {
				for _, file := range files[1:] {
					fmt.Printf("Remove the File hash: %s", hash)
					fmt.Printf("List the Remove file :")
					err := os.Link(files[0], file)
					if err != nil {
						fmt.Printf("remove file: %s ERROR : %s", file, err.Error())
					} else {
						fmt.Printf("RM: %s", file)
					}
				}
			}
			<-model.WriteCH
		}(hash, files)
	}
	wg.Wait()
}

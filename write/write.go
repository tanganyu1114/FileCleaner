package write

import (
	"FileCleaner/model"
	"fmt"
	"os"
	"sync"
)

// 在这个包主要操作文件，删除文件创建ln

func Write(dm string) {
	// 删除文件
	RemoveFile()
	if dm == "ln" {
		CreateLink()
	}
	// 退出协程 关闭通道
	model.SignalCH <- true
	close(model.SignalCH)
	close(model.RecordCH)
	close(model.ControlCH)
}

func RemoveFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.ControlCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() { <-model.ControlCH }()
			if len(files) > 1 {
				fmt.Printf("Remove the File hash: %s\n", hash)
				fmt.Printf("List the Remove file:\n")
				for _, file := range files[1:] {
					var stat = true
					err := os.Remove(file)
					if err != nil {
						fmt.Printf("ERROR: Remove File: %s  : %s\n", file, err.Error())
						stat = false
					} else {
						fmt.Printf("ReMove: %s\n", file)
					}
					rd := model.NewWrite(stat, hash, file, model.FileSize[hash])
					model.RecordCH <- rd
				}
			}

		}(hash, files)
	}
	wg.Wait()
}

func CreateLink() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.ControlCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() { <-model.ControlCH }()
			if len(files) > 1 {
				fmt.Printf("Create the File Link hash: %s\n", hash)
				fmt.Printf("Source File: %s\n", files[0])
				fmt.Printf("List the Create file:\n")
				for _, file := range files[1:] {
					var stat = true
					err := os.Link(files[0], file)
					if err != nil {
						stat = false
						fmt.Printf("ERROR: Create File Link: %s  : %s\n", file, err.Error())
					} else {
						fmt.Printf("Link: %s\n", file)
					}
					rd := model.NewLink(stat, files[0], file)
					model.RecordCH <- rd
				}
			}
		}(hash, files)
	}
	wg.Wait()
}

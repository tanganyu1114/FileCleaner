package write

import (
	"FileCleaner/model"
	"fmt"
	"os"
	"sync"
	"time"
)

// 在这个包主要操作文件，删除文件创建ln

func Write(dm string) {
	// 判断文件处理模式
	switch dm {
	case "dry":
		dryRunFile()
	case "rm":
		removeFile()
	case "ln":
		removeFile()
		linkFile()
	}
	// 等待record记录完成 退出协程 关闭通道
	fmt.Printf("[INFO]: Wait to record the write infomation .")
	for {
		if len(model.RecordCH) == 0 {
			fmt.Printf("\n[INFO]: Record write file info Complete\n")
			model.SignalCH <- true
			break
		}
		fmt.Printf(".")
		time.Sleep(time.Second)
	}
	close(model.SignalCH)
	close(model.RecordCH)
	close(model.ControlCH)
}

func dryRunFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.ControlCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() { <-model.ControlCH }()
			if len(files) > 1 {
				for _, file := range files[1:] {
					fmt.Printf("[INFO]: Duplicate File: %s\n", file)
					rd := model.NewWrite(true, hash, file, model.FileSize[hash])
					model.RecordCH <- rd
				}
			}
		}(hash, files)
	}
	wg.Wait()
}

func removeFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.ControlCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() { <-model.ControlCH }()
			if len(files) > 1 {
				for _, file := range files[1:] {
					var stat = true
					err := os.Remove(file)
					if err != nil {
						fmt.Printf("[ERROR]: Remove File: %s  : %s\n", file, err.Error())
						stat = false
					} else {
						fmt.Printf("[INFO]: Remove File: %s\n", file)
					}
					rd := model.NewWrite(stat, hash, file, model.FileSize[hash])
					model.RecordCH <- rd
				}
			}

		}(hash, files)
	}
	wg.Wait()
}

func linkFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.ControlCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() { <-model.ControlCH }()
			if len(files) > 1 {
				for _, file := range files[1:] {
					var stat = true
					err := os.Link(files[0], file)
					if err != nil {
						stat = false
						fmt.Printf("[ERROR]: Create File Link: %s  : %s, Source File: %s\n", file, err.Error(), files[0])
					} else {
						fmt.Printf("[INFO]: Create File Link %s\n", file)
					}
					rd := model.NewLink(stat, files[0], file)
					model.RecordCH <- rd
				}
			}
		}(hash, files)
	}
	wg.Wait()
}

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
	} else {
		close(model.LinkCH)
	}
}

func RemoveFile() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.WriteCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			defer func() {
				<-model.WriteCH
			}()
			if len(files) > 1 {
				fmt.Printf("Remove the File hash: %s\n", hash)
				fmt.Printf("List the Remove file:\n")
				for _, file := range files[1:] {
					err := os.Remove(file)
					if err != nil {
						fmt.Printf("ERROR: Remove File: %s  : %s\n", file, err.Error())
					} else {
						fmt.Printf("ReMove: %s\n", file)
					}
				}
			}

		}(hash, files)
	}
	wg.Wait()
	// 关闭写限制channel
	close(model.WriteCH)
}

func CreateLink() {
	wg := sync.WaitGroup{}
	for hash, files := range model.FileMap {
		wg.Add(1)
		model.LinkCH <- 1
		go func(hash string, files []string) {
			defer wg.Done()
			if len(files) > 1 {
				fmt.Printf("Create the File Link hash: %s\n", hash)
				fmt.Printf("Source File: %s\n", files[0])
				fmt.Printf("List the Create file:\n")
				for _, file := range files[1:] {
					err := os.Link(files[0], file)
					if err != nil {
						fmt.Printf("ERROR: Create File Link: %s  : %s\n", file, err.Error())
					} else {
						fmt.Printf("Link: %s\n", file)
					}
				}
			}
			<-model.LinkCH
		}(hash, files)
	}
	wg.Wait()
	close(model.LinkCH)
}

package main

import (
	"FileCleaner/control"
	"FileCleaner/read"
	"FileCleaner/record"
	"FileCleaner/write"
	"flag"
	"time"
)

var basePath string
var delMethod string
var casCade bool
var parallel int

func Init() {
	flag.StringVar(&basePath, "p", "/home/fastdfs/storage/data", "Plz enter an absolute path")
	flag.BoolVar(&casCade, "c", false, "Cascade control the dir, default false")
	flag.IntVar(&parallel, "n", 4, "The number of parallel control, default 4")
	flag.StringVar(&delMethod, "dm", "ln", "Plz Usage 'rm' or 'ln'; rm: remove the file ,ln: remove and link the file")
	flag.Parse()
}

func main() {
	startTime := time.Now()
	// 初始化输入参数
	Init()
	// 启用goroutine parallel限制
	control.Control(parallel)
	// goroutine 后台记录操作信息 并输出结果
	go record.Record()
	// 读文件
	read.Read(basePath, casCade)
	// 写文件
	write.Write(delMethod)
	// 输出统计结果信息
	record.Report(startTime)

}

package main

import (
	"FileCleaner/control"
	"FileCleaner/read"
	"FileCleaner/record"
	"FileCleaner/version"
	"FileCleaner/write"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	ver       bool
	basePath  string
	delMethod string
	casCade   bool
	parallel  int
	day       string
	ctime     bool
	mtime     bool
	atime     bool
)

// 初始化输入参数
func init() {
	flag.BoolVar(&ver, "v", false, "show version and exit")
	flag.StringVar(&basePath, "p", "/home/fastdfs/storage/data", "Plz enter an absolute path")
	flag.BoolVar(&casCade, "c", false, "Cascade control the dir (default false)")
	flag.IntVar(&parallel, "n", 4, "The number of parallel control")
	flag.StringVar(&delMethod, "dm", "ln", "Plz Usage 'rm','ln' or 'dry'; rm: remove the file, ln: remove and link the file, dry: dryRun mode")
	flag.StringVar(&day, "t", "", "Enter the day, like -1,0,+1")
	flag.BoolVar(&ctime, "ctime", false, "Enable clean file by the created time (default false) ")
	flag.BoolVar(&mtime, "mtime", false, "Enable clean file by the modified time (default false)")
	flag.BoolVar(&atime, "atime", false, "Enable clean file by the accessed time (default false)")
	//flag.Usage = usage()
	flag.Parse()
	if ver {
		fmt.Printf("FileCleaner Version: %s\n", version.Version)
		os.Exit(0)
	}
}

func main() {
	startTime := time.Now()
	// 启用goroutine parallel限制
	control.CtlParallel(parallel)
	// goroutine 后台记录操作信息 并输出结果
	go record.Record()
	// 读文件时间控制器
	read.ReadByTime(ctime, mtime, atime, day)
	// 读文件
	read.Read(basePath, casCade)
	// 写文件
	write.Write(delMethod)
	// 输出统计结果信息
	record.Report(startTime)

}

package main

import (
	"FileCleaner/read"
	"FileCleaner/write"
	"flag"
)

var basePath string
var delMethod string
var casCade bool

func Init() {
	flag.StringVar(&basePath, "p", "/home/fastdfs/storage/data", "Plz enter an absolute path")
	flag.BoolVar(&casCade, "c", false, "Is  the cascade dir file, default false")
	flag.StringVar(&delMethod, "dm", "ln", "Plz Usage 'rm' or 'ln'; rm: remove the file ,ln: remove and link the file")
	flag.Parse()
}

func main() {
	Init()
	go read.Record()
	read.Read(basePath, casCade)
	write.Write(delMethod)
}

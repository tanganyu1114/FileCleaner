package model

type fileMap map[string][]string

var (
	SignalCH = make(chan bool, 0)
	ReadCH   = make(chan int, 10)
	WriteCH  = make(chan int, 10)
	FileMap  = make(fileMap)
	Files    []string
)

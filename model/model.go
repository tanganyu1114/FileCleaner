package model

var (
	SignalCH = make(chan bool, 0)
	ReadCH   = make(chan int, 10)
	WriteCH  = make(chan int, 10)
	RecordCH = make(chan map[string]string, 10)
	FileMap  = make(map[string][]string)
	Files    []string
)

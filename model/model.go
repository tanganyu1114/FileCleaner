package model

type Result struct {
	TotalSize   int64
	ReadPathOK  uint
	ReadPathERR uint
	ReadFileOK  uint
	ReadFileERR uint
	RemoveOK    uint
	RemoveERR   uint
	LinkOK      uint
	LinkERR     uint
}

var (
	// 关闭记录goroutine的信号
	SignalCH = make(chan bool, 0)
	// 控制并发读的数量
	ReadCH = make(chan int, 10)
	// 控制并发写的数量
	WriteCH = make(chan int, 10)
	// 控制并发创建link的数量
	LinkCH = make(chan int, 10)
	// 记录通道的数量应该与读的数量一致
	RecordCH = make(chan map[string]string, 10)
	// hash与文件关系对应表
	FileMap = make(map[string][]string)
	// 文件组
	Files []string
	// 结果记录
	Res Result
)

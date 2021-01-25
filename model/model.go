package model

import "sync"

var (
	// 读取文件完成的信号
	SignalCH = make(chan bool, 0)
	// 控制并发数
	ControlCH = make(chan int, 10)
	// 通道传递信息到record协程
	RecordCH = make(chan Record, 10)
	// 记录文件hash值的map
	FileMap = make(map[string][]string)
	// 记录文件大小
	FileSize = make(map[string]int)
)

type Result struct {
	Path  pathRes
	Read  readRes
	Write writeRes
	Link  linkRes
}

type pathRes struct {
	TotalNum int // 总目录数
	ErrNum   int
	Errs     []string
}

type readRes struct {
	TotalNum  int // 读取数量
	TotalSize int // 读取文件大小
	ErrNum    int
	Errs      []string
}

type writeRes struct {
	TotalNum  int // 删除文件数量
	TotalSize int // 删除文件大小
	ErrNum    int
	Errs      []string
}

type linkRes struct {
	TotalNum int
	ErrNum   int
	ErrMap   map[string][]string // 源文件  目标文件
}

type Record interface {
	Record(res *Result)
}

func (p path) Record(res *Result) {
	res.Path.TotalNum += 1
	if !p.stat {
		res.Path.ErrNum += 1
		res.Path.Errs = append(res.Path.Errs, p.dir)
	}
}

type path struct {
	stat bool
	dir  string
}

func NewPath(stat bool, dir string) *path {
	return &path{
		stat: stat,
		dir:  dir,
	}
}

type read struct {
	wg   *sync.WaitGroup
	stat bool
	hash string
	file string
	size int
}

func NewRead(wg *sync.WaitGroup, stat bool, hash string, file string, size int) *read {
	return &read{
		wg:   wg,
		stat: stat,
		hash: hash,
		file: file,
		size: size,
	}
}

func (r read) Record(res *Result) {
	defer r.wg.Done()
	res.Read.TotalNum += 1
	res.Read.TotalSize += r.size
	if r.stat {
		FileMap[r.hash] = append(FileMap[r.hash], r.file)
		if _, ok := FileSize[r.hash]; !ok {
			FileSize[r.hash] = r.size
		}
	} else {
		res.Read.ErrNum += 1
		res.Read.Errs = append(res.Read.Errs, r.file)
	}
}

type write struct {
	wg   *sync.WaitGroup
	stat bool
	hash string
	file string
	size int
}

func NewWrite(wg *sync.WaitGroup, stat bool, hash string, file string, size int) *write {
	return &write{
		wg:   wg,
		stat: stat,
		hash: hash,
		file: file,
		size: size,
	}
}

func (w write) Record(res *Result) {
	defer w.wg.Done()
	res.Write.TotalNum += 1
	if w.stat {
		res.Write.TotalSize += w.size
	} else {
		res.Write.ErrNum += 1
		res.Write.Errs = append(res.Write.Errs, w.file)
	}

}

type link struct {
	wg   *sync.WaitGroup
	stat bool
	src  string
	dist string
}

func (l link) Record(res *Result) {
	defer l.wg.Done()
	res.Link.TotalNum += 1
	if !l.stat {
		res.Link.ErrNum += 1
		res.Link.ErrMap[l.src] = append(res.Link.ErrMap[l.src], l.dist)
	}
}

func NewLink(wg *sync.WaitGroup, stat bool, src, dist string) *link {
	return &link{
		wg:   wg,
		stat: stat,
		src:  src,
		dist: dist,
	}
}

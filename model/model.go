package model

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
	stat bool
	hash string
	file string
	size int
}

func NewRead(stat bool, hash string, file string, size int) *read {
	return &read{
		stat: stat,
		hash: hash,
		file: file,
		size: size,
	}
}

func (r read) Record(res *Result) {
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
	stat bool
	hash string
	file string
	size int
}

func NewWrite(stat bool, hash string, file string, size int) *write {
	return &write{
		stat: stat,
		hash: hash,
		file: file,
		size: size,
	}
}

func (w write) Record(res *Result) {
	res.Write.TotalNum += 1
	if w.stat {
		res.Write.TotalSize += w.size
	} else {
		res.Write.ErrNum += 1
		res.Write.Errs = append(res.Write.Errs, w.file)
	}

}

type link struct {
	stat bool
	src  string
	dist string
}

func (l link) Record(res *Result) {
	res.Link.TotalNum += 1
	if !l.stat {
		res.Link.ErrNum += 1
		res.Link.ErrMap[l.src] = append(res.Link.ErrMap[l.src], l.dist)
	}
}

func NewLink(stat bool, src, dist string) *link {
	return &link{
		stat: stat,
		src:  src,
		dist: dist,
	}
}

/*type fileMap map[string]struct {
	files []string
	size  int64
}

type FileMaps interface {
	Select(hash string) bool
	Insert(hash, file string, size int)
}

func (f fileMap) Select(hash string) bool {
	_, ok := f[hash]
	return ok
}

func (f fileMap) Insert(record Record) {
	if record.Select().Stats {
		f[record.Select().Hash] = struct {
			files []string
			size  int64
		}{
			files: append(f[record.Select().Hash].files, record.Select().File),
			size: int64(record.Select().Size)}
	}
}

type PathResult interface {
	AddTotalNum()
	AddTotalSize(size int64)
	AddErrInfo(path string)
}

type pathResult struct {
	totalNum  int
	totalSize int64
	errNum    int
	errPath   []string
}

func NewPathResult() *pathResult {
	return &pathResult{
		totalNum:  0,
		totalSize: 0,
		errNum:    0,
		errPath:   nil,
	}
}

func (p pathResult) AddTotalSize(size int64) {
	p.totalSize += size
}

func (p pathResult) AddTotalNum() {
	p.totalNum += 1
}

func (p pathResult) AddErrInfo(path string) {
	p.errNum += 1
	p.errPath = append(p.errPath, path)
}

type Record interface {
	Select() *record
}

type record struct {
	Stats bool
	Hash  string
	File  string
	Size  int
}

func NewRecord(stats bool, hash, file string, size int) *record {
	return &record{
		Stats: stats,
		Hash:  hash,
		File:  file,
		Size:  size,
	}
}

func (r record) Select() *record {
	return &r
}

type Result interface {
	Insert(record Record)
	Select() (totalSize int64, totalNum int, errNum int, errFile []string)
}

type path struct {
	totalNum int
	errNum   int
	errPath  []string
}

type read struct {
	totalSize int
	totalNum  int
	errNum    int
	errFile   []string
}

func NewRead() *read {
	return &read{
		totalSize: 0,
		totalNum:  0,
		errNum:    0,
		errFile:   make([]string, 0),
	}
}

type write struct {
	totalSize int
	totalNum  int
	errNum    int
	errFile   []string
}

func NewWrite() *write {
	return &write{
		totalSize: 0,
		totalNum:  0,
		errNum:    0,
		errFile:   make([]string, 0),
	}
}

type link struct {
	totalSize int
	totalNum  int
	errNum    int
	errFile   []string
}

func NewLink() *link {
	return &link{
		totalSize: 0,
		totalNum:  0,
		errNum:    0,
		errFile:   make([]string, 0),
	}
}

func (r read) Select() (totalSize int, totalNum int, errNum int, errFile []string) {
	return r.totalSize, r.totalNum, r.errNum, r.errFile
}

func (r write) Select() (totalSize int, totalNum int, errNum int, errFile []string) {
	return r.totalSize, r.totalNum, r.errNum, r.errFile
}

func (r link) Select() (totalSize int, totalNum int, errNum int, errFile []string) {
	return r.totalSize, r.totalNum, r.errNum, r.errFile
}

func (r read) Insert(record Record) {
	r.totalNum += 1
	r.totalSize += record.Select().Size
	if !record.Select().Stats {
		r.errNum += 1
		r.errFile = append(r.errFile, record.Select().File)
	}
}

func (r write) Insert(record Record) {
	r.totalNum += 1
	r.totalSize += record.Select().Size
	if !record.Select().Stats {
		r.errNum += 1
		r.errFile = append(r.errFile, record.Select().File)
	}
}

func (r link) Insert(record Record) {
	r.totalNum += 1
	r.totalSize += record.Select().Size
	if !record.Select().Stats {
		r.errNum += 1
		r.errFile = append(r.errFile, record.Select().File)
	}
}
*/

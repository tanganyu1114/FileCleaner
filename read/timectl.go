package read

import (
	"os"
	"strconv"
	"time"
)

var timeFilter fileFilterByTime

// 3个时间点情况，时间节点n
// 时间节点前 +n  n天前,不包括第n天
// 时间节点后 -n  n天前至今
// 时间节点时 n   第n天
func ReadByTime(ctime, mtime, atime bool, day string) {
	var (
		ctl = 0
		num = 0
		t   = time.Now()
	)
	if ctime {
		num += 1
		ctl = 1
	}
	if mtime {
		num += 1
		ctl = 2
	}
	if atime {
		num += 1
		ctl = 3
	}
	switch {
	case num > 1:
		panic("[ERROR]: Can not usage ctime/mtime/atime at the same time !")
	case num == 1:
		iday, err := strconv.Atoi(day)
		if err != nil {
			panic("[ERROR]: Error day format input !")
		}
		nodeTime := absTimeToDay(t.AddDate(0, 0, -abs(iday)).Unix())
		switch ctl {
		case 1: // ctime
			switch day[0:1] {
			case "+":
				timeFilter = NewCtimeBefore(nodeTime)
			case "-":
				timeFilter = NewCtimeAfter(nodeTime)
			default:
				timeFilter = NewCtimeEqual(nodeTime)
			}
		case 2: // mtime
			switch day[0:1] {
			case "+":
				timeFilter = NewMtimeBefore(nodeTime)
			case "-":
				timeFilter = NewMtimeAfter(nodeTime)
			default:
				timeFilter = NewMtimeEqual(nodeTime)
			}
		case 3: // atime
			switch day[0:1] {
			case "+":
				timeFilter = NewAtimeBefore(nodeTime)
			case "-":
				timeFilter = NewAtimeAfter(nodeTime)
			default:
				timeFilter = NewAtimeEqual(nodeTime)
			}
		default:
			timeFilter = nil
		}
	default:
		timeFilter = nil
	}

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func absTimeToDay(t int64) int64 {
	return t - t%86400
}

type fileFilterByTime interface {
	isOk(file os.FileInfo) bool
}

type ctimeBefore struct {
	time int64
}

type ctimeEqual struct {
	time int64
}

type ctimeAfter struct {
	time int64
}

func NewCtimeBefore(t int64) *ctimeBefore {
	return &ctimeBefore{time: t}
}

func NewCtimeEqual(t int64) *ctimeEqual {
	return &ctimeEqual{time: t}
}

func NewCtimeAfter(t int64) *ctimeAfter {
	return &ctimeAfter{time: t}
}

type mtimeBefore struct {
	time int64
}

type mtimeEqual struct {
	time int64
}

type mtimeAfter struct {
	time int64
}

func NewMtimeBefore(t int64) *mtimeBefore {
	return &mtimeBefore{time: t}
}

func NewMtimeEqual(t int64) *mtimeEqual {
	return &mtimeEqual{time: t}
}

func NewMtimeAfter(t int64) *mtimeAfter {
	return &mtimeAfter{time: t}
}

type atimeBefore struct {
	time int64
}

type atimeEqual struct {
	time int64
}

type atimeAfter struct {
	time int64
}

func NewAtimeBefore(t int64) *atimeBefore {
	return &atimeBefore{time: t}
}

func NewAtimeEqual(t int64) *atimeEqual {
	return &atimeEqual{time: t}
}

func NewAtimeAfter(t int64) *atimeAfter {
	return &atimeAfter{time: t}
}

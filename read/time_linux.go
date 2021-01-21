package read

import (
	"os"
	"syscall"
)

// linux系统的接口方法实现

// CTIME
// n天前
func (c ctimeBefore) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Ctim.Sec
	return fileTime < c.time
}

// 第n天
func (c ctimeEqual) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Ctim.Sec
	return absTimeToDay(fileTime) == c.time
}

// n天前至今
func (c ctimeAfter) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Ctim.Sec
	return fileTime > c.time
}

// MTIME
// n天前
func (c mtimeBefore) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Mtim.Sec
	return fileTime < c.time
}

// 第n天
func (c mtimeEqual) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Mtim.Sec
	return absTimeToDay(fileTime) == c.time
}

// n天前至今
func (c mtimeAfter) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Mtim.Sec
	return fileTime > c.time
}

// ATIME
// n天前
func (c atimeBefore) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Atim.Sec
	return fileTime < c.time
}

// 第n天
func (c atimeEqual) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Atim.Sec
	return absTimeToDay(fileTime) == c.time
}

// n天前至今
func (c atimeAfter) isOk(file os.FileInfo) bool {
	stat := file.Sys().(*syscall.Stat_t)
	fileTime := stat.Atim.Sec
	return fileTime > c.time
}

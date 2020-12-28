package record

import (
	"FileCleaner/model"
	"fmt"
)

// 后台记录记录文件信息
func Record() {
	// 操作结果集
	var result model.Result
LOOP:
	for {
		select {
		// 记录读文件信息
		case record := <-model.RecordCH:
			record.Record(&result)
		case <-model.SignalCH:
			break LOOP
		}
	}
	// 输出结果
	fmt.Printf(`
结果概览：
	总目录数量: %d    总文件数量: %d    重复文件数量: %d
	    	 
	总文件大小: %s    去重后文件大小: %s    重复文件大小: %s    去重空间比: %d

	读取文件目录成功: %d   失败: %d  
	读取文件数量成功: %d   失败: %d  
	删除文件数量成功: %d   失败: %d  
	创建文件链接成功: %d   失败: %d  

`)
}

/*
输出结果信息
总计：
	总文件大小:		去重后文件大小:		去重大小:	去重空间占比:
	读取文件目录:     成功:    失败:
	读取文件数量:		成功:	失败:
	删除文件数量:		成功:	失败:
	创建文件链接:		成功:	失败:
if 失败> 0 {
失败详情:
	* * *
}
*/

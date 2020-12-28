package record

import (
	"FileCleaner/model"
	"fmt"
	"time"
)

// 操作结果集
var result model.Result

// 后台记录记录文件信息
func Record() {
	for {
		select {
		// 记录读文件信息
		case record := <-model.RecordCH:
			record.Record(&result)
		case <-model.SignalCH:
			return
		}
	}
}

func Report(t time.Time) {
	totalSize := fmt.Sprintf("%.2f GB", float64(result.Read.TotalSize)/1024/1024/1024)
	removeSize := fmt.Sprintf("%.2f GB", float64(result.Write.TotalSize)/1024/1024/1024)
	saveSize := fmt.Sprintf("%.2f GB", float64(result.Read.TotalSize-result.Write.TotalSize)/1024/1024/1024)
	numPct := fmt.Sprintf("%.2f %%", float64(result.Write.TotalNum)/float64(result.Read.TotalNum))
	sizePct := fmt.Sprintf("%.2f %%", float64(result.Write.TotalSize)/float64(result.Read.TotalSize))
	// 输出结果
	fmt.Printf(`
本次操作总计耗时: %s
结果概览:
	总目录数量: %d    总文件数量: %d    重复文件数量: %d
	    	 
	总文件大小: %s    重复文件大小: %s    剩余文件大小: %s    

	文件去重比: %s	空间优化比: %s

	读取文件目录 成功: %d   失败: %d  
	读取文件数量 成功: %d   失败: %d  
	删除文件数量 成功: %d   失败: %d  
	创建文件链接 成功: %d   失败: %d  

`, time.Since(t).String(),
		result.Path.TotalNum, result.Read.TotalNum, result.Write.TotalNum,
		totalSize, removeSize, saveSize,
		numPct, sizePct,
		result.Path.TotalNum-result.Path.ErrNum, result.Path.ErrNum,
		result.Read.TotalNum-result.Read.ErrNum, result.Read.ErrNum,
		result.Write.TotalNum-result.Write.ErrNum, result.Write.ErrNum,
		result.Link.TotalNum-result.Link.ErrNum, result.Link.ErrNum)

	if result.Path.ErrNum != 0 {
		fmt.Printf("读取目录失败列表:\n")
		for _, path := range result.Path.Errs {
			fmt.Println(path)
		}
	}
	if result.Read.ErrNum != 0 {
		fmt.Printf("读取文件失败列表:\n")
		for _, path := range result.Read.Errs {
			fmt.Println(path)
		}
	}
	if result.Write.ErrNum != 0 {
		fmt.Printf("删除文件失败列表:\n")
		for _, path := range result.Write.Errs {
			fmt.Println(path)
		}
	}
	if result.Link.ErrNum != 0 {
		fmt.Printf("创建文件硬链接失败列表:\n")
		for src, dist := range result.Link.ErrMap {
			fmt.Printf("源文件:%s  -->  目标文件:%s\n", src, dist)
		}
	}
}

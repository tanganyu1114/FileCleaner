# FILECLEANER
# 命令行文件去重工具

**编写这个工具的初衷在于给fastdfs小文件进行去重
并发的读写操作文件，目前默认goroutine数量为10,如需修改，只需要修改model包对应的channel数量**
* * * 
* 命令行参数
```
> -c bool  是否启用级联去重操作,默认false  
> -p string  去重操作的目录,默认路径/home/fastdfs/storage/data  
> -dm string  删除的方式，目前支持2种参数：rm 直接删除重复文件; ln 删除重复文件，并且以第一个文件为源目标，创建其他文件的硬链接
> -n int   并发执行的线程数，默认4，最低1，最高10
```
* 示例  
```
> filecleaner -c=true -p="/home/data"
```


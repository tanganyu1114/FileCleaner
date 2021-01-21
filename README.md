# FILECLEANER
# 命令行文件去重工具

**编写这个工具的初衷在于给fastdfs小文件进行去重
并发的读写操作文件，目前默认goroutine最大数量为10,如需修改，只需要修改model包对应的channel数量**
* * * 
* 命令行参数
```
> -c bool  是否启用级联去重操作(默认 false)
> -p string  去重操作的目录 (默认路径 /home/fastdfs/storage/data)  
> -dm string  删除的方式，目前支持3种参数(默认 ln):
              -dm="dry" 只模拟查看重复文件和执行效果，不真实执行删除操作
              -dm="rm" 直接删除重复文件; 
              -dm="ln" 删除重复文件，并且以第一个文件为源目标，创建其他文件的硬链接
> -n int   并发执行的线程数，最低1，最高10(默认 4)
> -ctime bool 是否按文件创建时间去重
> -mtime bool 是否按文件修改时间去重
> -atime bool 是否按文件访问时间去重
> -t string  时间点，例如:
        +n  n+1天前
         n  n+1天前到n天前的时间段
        -n  n天前至今           
```
* 示例  
```
# 级联清理文件
> filecleaner -c=true -p="/home/data"

# 按文件创建时间清理，清理10天前的
> filecleaner -c=true -p="/home/data" -ctime -t=+10

# 只模拟查看清理效果，不实际删除文件
> filecleaner -c=true -p="/home/data" -dm="dry"
```
* 执行结果展示
![执行结果](https://gitee.com/cosNeaby/FileCleaner/raw/master/picture/result.jpg)


* LICENSE 
```
MIT License

Copyright (c) 2021 沧海cos泪

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

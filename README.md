# filetools
文件处理工具

### 用法
```
-m  模块名  content 和 name   content 是修改内容， name 是修改文件名. delete  删除文件
-f  文件名， 只有修改内容才有
-d  目录
-o  旧字符串
-n  替换成的新字符串
-H   是否忽略隐藏文件
-i  指定包含字符串的文件名，逗号分隔多个, 修改内容才有效
-e   跳过指定包含字符串的文件名，逗号分隔多个， 修改内容才有效
// 将文件夹 /data/ 下面所有文件内容里面的111替换成222
go run ft.go -m content -d /data/ -o 111 -n 222
// 将文件夹 /data/ 下面所有文件名111替换成222
go run ft.go -m name -d /data/ -o 111 -n 222
// 删除超过1天的txt文件
go run ft.go -m delete -d /data/ -i txt -t 1
```

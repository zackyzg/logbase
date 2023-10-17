# logbase
golang 项目日志库
1. go run mylog.go 运行程序
2. 修改mylog.go 中 log = myselflog.NewTypeLog(2) 的参数</br>实现日志输出位置改变 1. 日志输出到控制台 2. 日志输出到文件
3. myselflog下面的console.go是日志输出到控制的的逻辑</br>file是日志输出到文件的逻辑
4. 日志写入文件有两种分割方式日志分割类型1.时间分割2.大小分割;</br>修改myselflog.file种的logsplittype来实现


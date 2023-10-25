/**
 *
 * @package       common
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package myselflog

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

type LogInformation struct {
	Level     LogLevel
	Callname  string
	FormatMsg string
	FuncName  string
	FileName  string
	LineNo    int
	Timestamp string
}

type LogLevel uint16

var logsplittype uint8 = 1           // 日志分割类型1.时间分割2.大小分割;可以使用配置文件
var loglevellimit LogLevel = WARNING // 日志记录级别，级别高于当前设置的日志将被记录
var logbuffersize int = 50000        // 日志缓冲通道容量
var logoutputsize int = 10           // 日志取出channel(消费者)goroutine开启数量

const (
	UNKONW LogLevel = iota
	DEBUG
	TRACE
	WARNING
	INFO
	ERROR
	FATAL
)

func SetLogLv(logstr string) {
	Loglv, err := Check(logstr)
	if err != nil {
		fmt.Println(err)
	}
	loglevellimit = Loglv
}

func GetLogLv() LogLevel {
	return loglevellimit
}

func Check(logstr string) (LogLevel, error) {
	str := strings.ToUpper(logstr)

	switch str {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "WARNING":
		return WARNING, nil
	case "INFO":
		return INFO, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		return UNKONW, errors.New("Undefined log level")
	}
}

func unusualPosInfo(n int) (funcName string, filename string, lineNo int) {
	pc, filename, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]

	//path.Base Base返回path的最后一个元素
	//fmt.Println(funcName, path.Base(file), lineNo)
	//filename = path.Base(file)

	return
}

func NewTypeLog(logtype int) Logger {

	if logtype == 1 {
		return NewConsoleLog()
	} else if logtype == 2 {
		return NewFileLog()
	} else {
		panic(errors.New("parameter error"))
	}

}

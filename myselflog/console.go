/**
 *
 * @package       consolelog
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

// 日志写入控制台
package myselflog

import (
	"fmt"
	"time"
)

// ConsoleLogger  控制台输出日志结构体...
type ConsoleLogger struct {
	Level   LogLevel
	logchan chan *LogInformation
}

// NewConsoleLog  控制台输出日志初始化...
func NewConsoleLog() *ConsoleLogger {
	initLogger := &ConsoleLogger{
		Level:   GetLogLv(),
		logchan: make(chan *LogInformation, logbuffersize),
	}

	for i := 0; i < logoutputsize; i++ {
		go initLogger.LogConsumer()
	}

	return initLogger
}

func (c *ConsoleLogger) LogConsumer() {
	for {
		select {
		case logdata := <-c.logchan:
			fmt.Printf("[%s][%s][%s:%s:%d] %s\n", logdata.Timestamp, logdata.Callname, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.FormatMsg)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func (c *ConsoleLogger) log(callname string, format string, a ...interface{}) {
	Loglv, err := Check(callname)
	if err != nil {
		fmt.Println(err)
	}

	funcName, filename, lineNo := unusualPosInfo(3)
	msg := fmt.Sprintf(format, a...)

	logift := &LogInformation{
		Level:     Loglv,
		Callname:  callname,
		FormatMsg: msg,
		FuncName:  funcName,
		FileName:  filename,
		LineNo:    lineNo,
		Timestamp: time.Now().Format(time.DateTime),
	}

	if c.Level <= Loglv {

		select {
		case c.logchan <- logift:
		default:
			return
		}
	}
}

func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log("DEBUG", format, a...)
}

func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log("Trace", format, a...)
}

func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log("Warning", format, a...)
}

func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.log("Info", format, a...)
}

func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.log("Error", format, a...)
}

func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log("Fatal", format, a...)
}

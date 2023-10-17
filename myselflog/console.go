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

type ConsoleLogger struct {
	Level LogLevel
}

func NewConsoleLog() *ConsoleLogger {

	return &ConsoleLogger{
		Level: GetLogLv(),
	}
}

func (c *ConsoleLogger) log(callname string, format string, a ...interface{}) {
	Loglv, err := Check(callname)
	if err != nil {
		fmt.Println(err)
	}

	funcName, filename, lineNo := unusualPosInfo(4)

	if c.Level <= Loglv {
		msg := fmt.Sprintf(format, a)
		fmt.Printf("[%s][%s][%s:%s:%d] %s\n", time.Now().Format(time.DateTime), callname, filename, funcName, lineNo, msg)
	}
}

func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log("DEBUG", format, a)
}

func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log("Trace", format, a)
}

func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log("Warning", format, a)
}

func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.log("Info", format, a)
}

func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.log("Error", format, a)
}

func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log("Fatal", format, a)
}

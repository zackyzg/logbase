/**
 *
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"logbase/myselflog"
	"time"
)

// [优化建议]为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。
// 1.将日志信息通过一个或者多个后台goroutine写入通道(生产者)
// 2.通过多个gouroutine从通道中读取日志信息，写入文件(消费者)

var log myselflog.Logger

func init() {
	myselflog.SetLogLv("WARNING")
}

func main() {
	log = myselflog.NewTypeLog(2) // 1. 日志输出到控制台 2. 日志输出到文件

	n := 0
	for {
		name := "zack"
		id := 35
		log.Debug("第一--->log,id:%d,name:%s", id, name)
		log.Trace("第二--->log")
		log.Warning("第三--->log")
		log.Info("第四--->log")
		log.Error("第五--->log")
		log.Fatal("第六--->log,id:%d,name:%s", id, name)

		n++

		time.Sleep(time.Millisecond * 100)

		//if n > 120 {
		//	break
		//}
	}

}

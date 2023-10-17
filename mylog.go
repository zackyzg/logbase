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

var log myselflog.Logger

func init() {
	myselflog.SetLogLv("debug")
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

		time.Sleep(time.Second * 1)

		//if n > 120 {
		//	break
		//}
	}

}

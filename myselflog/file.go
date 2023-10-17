/**
 *
 * @package       filelog
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

// 日志写入文件
package myselflog

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const dateformat = "200601021504" // 可以使用配置文件,这决定了按时间分割的时候是按照分钟、小时、天分割

var logsplittype uint8 = 1 // 日志分割类型1.时间分割2.大小分割;可以使用配置文件

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志路径,可以使用配置文件
	fileName    string // 日志文件名 sys.+日期
	errFileName string // 错误日志 err.+日期
	maxFileSize int64  // 文件分割大小限制
	slicingType uint8  // 切割文件方式 1.文件大小切割 2.文件日期切割
	fileobj     *os.File
	errfileobj  *os.File
	// 日志分割方式:大小分割，时间分割
}

func NewFileLog() *FileLogger {
	//year := time.Now().Year()
	//month := int(time.Now().Month()) // 将英文月份转成数字月份
	//day := time.Now().Day()

	fPath := "./" // 可以使用配置文件
	fName := "sys." + time.Now().Format(dateformat) + ".log"
	efName := "err." + time.Now().Format(dateformat) + ".log"
	mfSize := 1 * 1024 * 1024

	fl := &FileLogger{
		Level:       GetLogLv(),
		filePath:    fPath,
		fileName:    fName,
		errFileName: efName,
		maxFileSize: int64(mfSize),
	}

	err := fl.initfile()
	if err != nil {
		panic(err)
	}

	return fl
}

func (f *FileLogger) initfile() error {
	allFileName := path.Join(f.filePath, f.fileName)
	allErrFileName := path.Join(f.filePath, f.errFileName)

	fobj, err := os.OpenFile(allFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return err
	}

	efobj, err := os.OpenFile(allErrFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file failed,err:", err)
		return err
	}
	f.fileobj = fobj
	f.errfileobj = efobj
	return nil
}

func (f *FileLogger) close() {
	f.fileobj.Close()
	f.errfileobj.Close()
}

func (f *FileLogger) log(callname string, format string, a ...interface{}) {

	switch logsplittype {
	case 1:
		f.logTimeSplit(callname, format, a...)
		break
	case 2:
		f.logSizeSplit(callname, format, a...)
		break
	default:
		panic(errors.New("split type parameter error"))

	}

}

func (f *FileLogger) logSizeSplit(callname string, format string, a ...interface{}) {
	Loglv, err := Check(callname)
	if err != nil {
		fmt.Println(err)
	}

	funcName, filename, lineNo := unusualPosInfo(4)
	msg := fmt.Sprintf(format, a...)
	if f.Level <= Loglv {
		newFile, err := f.fileSizeSplit(f.fileobj)
		if err != nil {
			return
		}
		f.fileobj = newFile
		fmt.Fprintf(f.fileobj, "[%s][%s][%s:%s:%d] %s\n", time.Now().Format(time.DateTime), callname, filename, funcName, lineNo, msg)

	}

	if Loglv >= ERROR {
		newErrFile, err := f.fileSizeSplit(f.errfileobj)
		if err != nil {
			return
		}
		f.errfileobj = newErrFile
		fmt.Fprintf(f.errfileobj, "[%s][%s][%s:%s:%d] %s\n", time.Now().Format(time.DateTime), callname, filename, funcName, lineNo, msg)
	}

}

func (f *FileLogger) logTimeSplit(callname string, format string, a ...interface{}) {
	Loglv, err := Check(callname)
	if err != nil {
		fmt.Println(err)
	}

	funcName, filename, lineNo := unusualPosInfo(4)

	f.fileTimeSplit()
	msg := fmt.Sprintf(format, a...)
	if f.Level <= Loglv {
		fmt.Fprintf(f.fileobj, "[%s][%s][%s:%s:%d] %s\n", time.Now().Format(time.DateTime), callname, filename, funcName, lineNo, msg)

	}

	if Loglv >= ERROR {
		fmt.Fprintf(f.errfileobj, "[%s][%s][%s:%s:%d] %s\n", time.Now().Format(time.DateTime), callname, filename, funcName, lineNo, msg)
	}

}

func (f *FileLogger) fileSizeSplit(file *os.File) (*os.File, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if fileInfo.Size() >= f.maxFileSize {
		file.Close()
		oldFilePath := path.Join(f.filePath, fileInfo.Name())
		newFilePath := path.Join(f.filePath, time.Now().Format("20060102150405")+"_"+fileInfo.Name())
		err := os.Rename(oldFilePath, newFilePath)
		if err != nil {
			panic(errors.New("rename log file filed"))
			return nil, err
		}

		fmt.Println(fileInfo.Name() + "文件大小分割log文件，进来了")

		// 打开文件
		fobj, err := os.OpenFile(oldFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("open log file failed,err:", err)
			return nil, err
		}

		return fobj, nil

	} else {
		return file, nil
	}

}

func (f *FileLogger) fileTimeSplit() {
	old_key := strings.Split(f.fileobj.Name(), ".")[1]
	now_key := time.Now().Format(dateformat)

	if old_key != now_key { // 关闭文件，打开新的文件
		fmt.Println("时间分割文件，进来了")
		f.close()
		f.fileName = "sys." + time.Now().Format(dateformat) + ".log"
		f.errFileName = "err." + time.Now().Format(dateformat) + ".log"

		err := f.initfile()
		if err != nil {
			panic(err)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log("DEBUG", format, a...)

}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log("Trace", format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log("Warning", format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log("Info", format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log("Error", format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log("Fatal", format, a...)
}

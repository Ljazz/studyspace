package logger

import (
	"fmt"
	"time"
)

// 往文件里面写日志相关代码

// FileLogger 结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1:= &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile()	// 
	if err!=nil{
		panic(err)
	}
	return f1
}

func (f *FileLogger)initFile()(error){
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err!=nil{
		fmt.Printf("open log file failed, err:%v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName + ".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err!=nil{
		fmt.Printf("open err log file failed, err:%v", err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errFileObj
}

func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Frintf(, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("20006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log()
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	if f.enable(INFO) {
		// now := time.Now()
		// fmt.Fprintln("[%s] [INFO] %s\n", now.Format("20006-01-02 15:04:05"), msg)

		log(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	if f.enable(WARNING) {
		// now := time.Now()
		// fmt.Fprintln("[%s] [WARNING] %s\n", now.Format("20006-01-02 15:04:05"), msg)

		log(WARNING, format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	if f.enable(ERROR) {
		// now := time.Now()
		// fmt.Fprintln("[%s] [ERROR] %s\n", now.Format("20006-01-02 15:04:05"), msg)

		log(ERROR, format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	if f.enable(FATAL) {
		// now := time.Now()
		// fmt.Fprintln("[%s] [FATAL] %s\n", now.Format("20006-01-02 15:04:05"), msg)

		log(FATAL, format, a...)
	}
}

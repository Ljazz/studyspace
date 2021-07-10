package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("20006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

// Debug级别日志
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	// if c.enable(DEBUG) {
	// 	now := time.Now()
	// 	funcName, fileName, lineNo := getInfo(2)
	// 	fmt.Fprintln("[%s] [DEBUG] %s\n", now.Format("20006-01-02 15:04:05"), msg)

	// 	log(DEBUG, format, a...)
	// }
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	// if c.enable(INFO) {
	// 	now := time.Now()
	// 	fmt.Fprintln("[%s] [INFO] %s\n", now.Format("20006-01-02 15:04:05"), msg)

	// 	log(INFO, format, a...)
	// }
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	// if c.enable(WARNING) {
	// 	now := time.Now()
	// 	fmt.Fprintln("[%s] [WARNING] %s\n", now.Format("20006-01-02 15:04:05"), msg)

	// 	log(WARNING, format, a...)
	// }
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	// if c.enable(ERROR) {
	// 	now := time.Now()
	// 	fmt.Fprintln("[%s] [ERROR] %s\n", now.Format("20006-01-02 15:04:05"), msg)

	// 	log(ERROR, format, a...)
	// }
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	// if c.enable(FATAL) {
	// 	now := time.Now()
	// 	fmt.Fprintln("[%s] [FATAL] %s\n", now.Format("20006-01-02 15:04:05"), msg)

	// 	log(FATAL, format, a...)
	// }
	c.log(FATAL, format, a...)
}

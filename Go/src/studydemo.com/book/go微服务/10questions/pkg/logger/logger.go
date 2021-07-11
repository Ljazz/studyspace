package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"studydemo.com/book/go微服务/10questions/pkg/config"
)

var inforLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger
var debugLogger *log.Logger

func init() {
	logFilePath := filepath.Join(config.LogDir, fmt.Sprintf("share-%s.log", time.Now().Format("2006-01-02")))
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("fail to create share.log")
	}
	inforLogger = log.New(logFile, "[Info]", log.LstdFlags|log.Lshortfile)
	warnLogger = log.New(logFile, "[Warn]", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(logFile, "[Error]", log.LstdFlags|log.Lshortfile)
	debugLogger = log.New(logFile, "[Debug]", log.LstdFlags|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
	if v == nil {
		inforLogger.Println(format)
	} else {
		inforLogger.Println(format+"\n", v)
	}
}

func Warn(format string, v ...interface{}) {
	if v == nil {
		warnLogger.Println(format)
	} else {
		warnLogger.Println(format+"\n", v)
	}
}

func Error(format string, v ...interface{}) {
	if v == nil {
		errorLogger.Println(format)
	} else {
		errorLogger.Println(format+"\n", v)
	}
}

func Debug(format string, v ...interface{}) {
	if v == nil {
		debugLogger.Println(format)
	} else {
		debugLogger.Println(format+"\n", v)
	}
}

func Panic(v ...interface{}) {
	log.Panic(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

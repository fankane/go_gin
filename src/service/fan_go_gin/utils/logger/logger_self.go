package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var logger *log.Logger
var logFileBasePath = "/Users/klook/Documents/klook/workspace_hf/go_gin/src/service/fan_go_gin/"

func InitLog() error {
	if logger == nil {
		logger = &log.Logger{}
	}
	logFile := fmt.Sprintf("%s%s", logFileBasePath, "/fan_go.log")
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		//不存在，创建即可
	} else {
		//存在，需要新建
		logFile = fmt.Sprintf("%s%s", logFileBasePath, fmt.Sprintf("/fan_go-%d.log", time.Now().Unix()))
	}
	_, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("日志文件创建失败 file:%s, err:%s", logFile, err)
	}
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	mulF := io.MultiWriter(os.Stdout, f) //设置 multiwriter, 让日志能同时打印到 控制台 和 log 文件
	logger.SetOutput(mulF)
	Infof("初始化 log 写入文件成功, file:%s", logFile)
	return nil
}

func Info(format string, v ...interface{}) {
	newV := getNewV(v...)
	logger.Print(newV...)
}
func Infof(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%s %s", getCurrentTime(), format), v...)
}

func Warn(v ...interface{}) {
	newV := getNewV(v...)
	logger.Print(newV...)
}

func Warnf(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%s %s", getCurrentTime(), format), v...)
}

func Error(v ...interface{}) {
	newV := getNewV(v...)
	logger.Print(newV...)
}
func Errorf(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%s %s", getCurrentTime(), format), v...)
}

func Debug(v ...interface{}) {
	newV := getNewV(v...)
	logger.Print(newV...)
}

func Debugf(format string, v ...interface{}) {
	logger.Printf(fmt.Sprintf("%s %s", getCurrentTime(), format), v...)
}

func getNewV(v ...interface{}) []interface{} {
	newV := make([]interface{}, 0)
	newV = append(newV, getCurrentTime())
	newV = append(newV, v...)
	return newV
}
func getCurrentTime() string {
	return fmt.Sprintf("%s : ", time.Now().Format(time.RFC3339Nano))
}

package logger

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

func InitLog() error {
	if logger == nil {
		logger = &log.Logger{}
	}
	logFileFullPath := "/Users/klook/Documents/klook/workspace_hf/go_gin/src/service/fan_go_gin/fan_go.log"
	f, err := os.OpenFile(logFileFullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	mulF := io.MultiWriter(os.Stdout, f) //设置 multiwriter, 让日志能同时打印到 控制台 和 log 文件
	logger.SetOutput(mulF)
	Infof("初始化 log 写入文件成功, file:%s", logFileFullPath)
	return nil
}

func Infof(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Warn() {

}

func Error(v ...interface{}) {
	logger.Print(v...)
}
func Errorf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Debug() {

}

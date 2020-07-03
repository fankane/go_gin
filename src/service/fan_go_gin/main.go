package main

import (
	"log"
	"service/fan_go_gin/config"
	"service/fan_go_gin/router"
	"service/fan_go_gin/utils/logger"
)



func main() {

	err := logger.InitLog()
	if err != nil {
		log.Fatalf("logger 初始化失败 err: %s", err)
		return
	}

	err = config.InitConfig()
	if err != nil {
		log.Fatalf("config 初始化失败 err: %s", err)
		return
	}
	router := router.InitRouter()

	router.Run(":9000")
	logger.Infof("服务启动成功")

}

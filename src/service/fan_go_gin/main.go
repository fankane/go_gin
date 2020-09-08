package main

import (
	"fmt"
	"log"
	"os/exec"
	"service/fan_go_gin/config"
	"service/fan_go_gin/router"
	"service/fan_go_gin/utils/logger"
	"time"
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

	go func() {
		time.Sleep(time.Second * 3)
		err := openHtml()
		if err != nil {
			logger.Error("打开网页失败 err:", err)
			return
		}
	}()

	router.Run(fmt.Sprintf(":%d", config.Conf.HttpPort))
	logger.Infof("服务启动成功")

}

func openHtml() error {

	cmd := exec.Command("open", "http://localhost:9002/assets/index.html")
	return cmd.Run()
}

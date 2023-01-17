package main

import (
	"fmt"
	"go_gin/src/service/fan_go_gin/config"
	"go_gin/src/service/fan_go_gin/router"
	"go_gin/src/service/fan_go_gin/utils/logger"
	"log"
	"os/exec"
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
		if !config.Conf.IsLocal {
			return
		}
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
	cmd := exec.Command("open", "http://localhost:9001/assets/index.html")
	return cmd.Run()
}

package main

import "service/fan_go_gin/router"



func main() {
	router := router.InitRouter()

	router.Run(":9000")
}

package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/fan_go_gin/controller/user"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//静态资源路由
	router.Static("/assets", "./assets")

	//router.Handlers
	router.GET("/v0/user/create", gin.HandlerFunc(func(context *gin.Context) {}))


	userR := router.Group("/v1/user")
	{
		userR.Handle(http.MethodGet, "/user/list", HandlerFunc(user.GetUserList))
		userR.Handle(http.MethodPost, "/user/create", HandlerFunc(user.CreateUser))
	}



	return router
}

func HandlerFunc(f gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		requestID := time.Now().Unix()
		context.Set("RequestID", fmt.Sprintf("%d", requestID))
		f(context)
	}
}
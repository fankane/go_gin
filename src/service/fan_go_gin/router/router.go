package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/src/service/fan_go_gin/controller/inventory"
	"go_gin/src/service/fan_go_gin/controller/user"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//静态资源路由
	router.Static("/assets", "./assets")

	//router.Handlers
	//router.GET("/v1/user/create", gin.HandlerFunc(func(context *gin.Context) {}))
	//router.GET("/v1/user/create", gin.HandlerFunc(func(context *gin.Context) {}))

	fileR := router.Group("/v1/file")
	{
		fileR.Handle(http.MethodPost, "/upload/download/url", HandlerFunc(inventory.UploadCSVFile))
		fileR.Handle(http.MethodGet, "/upload/download/process", HandlerFunc(inventory.CheckProcess))
	}


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

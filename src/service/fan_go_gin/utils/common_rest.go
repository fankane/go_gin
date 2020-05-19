package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/fan_go_gin/model"
)

func GetRequestID(ctx *gin.Context) string {
	return ""
}



func ReturnSuccessJson(ctx *gin.Context, obj interface{}) {
	resp := model.FResponse{}
	resp.Result = obj
	ctx.JSON(http.StatusOK, resp)
}
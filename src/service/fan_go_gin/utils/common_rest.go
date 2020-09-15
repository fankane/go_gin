package utils

import (
	"github.com/gin-gonic/gin"
	"go_gin/src/service/fan_go_gin/model"
	"net/http"
)

func GetRequestID(ctx *gin.Context) string {
	return ""
}



func ReturnSuccessJson(ctx *gin.Context, obj interface{}) {
	resp := model.FResponse{}
	resp.Result = obj
	resp.Success = true
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, resp)
}

func ReturnFailedJson(ctx *gin.Context, code int ,message string) {
	resp := model.FResponse{}
	resp.Success = false
	resp.Error.Code = code
	resp.Error.Message = message
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, resp)
}

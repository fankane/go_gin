package user

import (
	"github.com/gin-gonic/gin"
	"service/fan_go_gin/utils"
	"service/fan_go_gin/utils/logger"
)

func CreateUser(ctx *gin.Context) {



}


func GetUserList(ctx *gin.Context) {
	reqID, exist :=ctx.Get("RequestID")
	if !exist {
		logger.Infof("RequestID 不存在")
		return
	}
	logger.Infof("%s 获取成功", reqID)
	utils.ReturnSuccessJson(ctx, "hello")
}
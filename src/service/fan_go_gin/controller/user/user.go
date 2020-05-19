package user

import (
	"github.com/gin-gonic/gin"
	"service/fan_go_gin/utils"
)

func CreateUser(ctx *gin.Context) {



}


func GetUserList(ctx *gin.Context) {
	utils.ReturnSuccessJson(ctx, "hello")
}
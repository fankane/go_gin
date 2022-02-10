package image

import (
	"github.com/gin-gonic/gin"
	"go_gin/src/service/fan_go_gin/config"
	"go_gin/src/service/fan_go_gin/model"
	"go_gin/src/service/fan_go_gin/utils"
	"go_gin/src/service/fan_go_gin/utils/logger"
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(ctx *gin.Context) {
	header, err := ctx.FormFile("file")
	if err != nil {
		logger.Errorf("获取文件失败 err:%s", err)
		utils.ReturnFailedJson(ctx, model.CodeParamsError, err.Error())
		return
	}
	dst := header.Filename
	// gin 简单做了封装,拷贝了文件流
	if err := saveUploadedFile(header, dst); err != nil {
		logger.Errorf("获取文件失败 err:%s", err)
		utils.ReturnFailedJson(ctx, model.CodeSystemError, err.Error())
		return
	}
	utils.ReturnSuccessJson(ctx, "success")
	return
}


func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	dstF := config.Conf.ImagePathPre + dst
	out, err := os.Create(dstF)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}
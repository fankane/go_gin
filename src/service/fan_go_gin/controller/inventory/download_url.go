package inventory

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"service/fan_go_gin/model"
	"service/fan_go_gin/utils"
	"service/fan_go_gin/utils/logger"
	"service/fan_go_gin/utils/pdf"
	"strconv"
	"strings"
	"time"
)

var (
	total   = 0
	success = 0
	failed  = 0
	hasTask = false
)

type ProcessResp struct {
	Total   int     `json:"total"`
	Success int     `json:"success"`
	Failed  int     `json:"failed"`
	Percent float64 `json:"percent"`
}

func UploadCSVFile(ctx *gin.Context) {
	if hasTask {
		utils.ReturnFailedJson(ctx, model.CodeSystemError, "有任务正在处理")
		return
	}
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
	go func() {
		now := time.Now()
		err := startDownload(ctx.Copy(), dst)
		if err != nil {
			logger.Errorf("开始下载文件失败 err:", err)
			return
		}
		logger.Infof("下载完成, 耗时:%s", time.Since(now))
	}()
	utils.ReturnSuccessJson(ctx, "success")
	hasTask = true
	return
}

func CheckProcess(ctx *gin.Context) {
	logger.Infof("total:%d, success:%d, failed:%d", total, success, failed)
	percent := float64(success+failed) / float64(total) * 100.0
	logger.Infof("当前percent:%f", percent)
	percent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", percent), 64) //保留2位小数

	resp := &ProcessResp{}
	resp.Total = total
	resp.Success = success
	resp.Failed = failed
	resp.Percent = percent

	utils.ReturnSuccessJson(ctx, resp)
	return
}

func startDownload(ctx context.Context, file string) error {
	if strings.TrimSpace(file) == "" {
		return errors.New("file 为空")
	}
	urlList, err := getDownloadURL(file)
	if err != nil {
		return err
	}
	logger.Infof("总共url 数量:%d", len(urlList))
	total = len(urlList)
	downloadPath := "./download"
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		//不存在，创建即可
		err = os.Mkdir(downloadPath, 0755)
		if err != nil {
			return err
		}
	}

	for i, v := range urlList {
		name := fmt.Sprintf("System-%d.pdf", i+1)
		pdfT := fmt.Sprintf("%s/%s", downloadPath, name)
		err = pdf.WkhtmltoPDF(v, pdfT)
		if err != nil {
			logger.Errorf("下载失败 err:%s, file:%s, url:%s", err, pdfT, v)
			failed++
			continue
		}
		success++
	}

	return nil
}

func getDownloadURL(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("打开文件:%s 失败 err:%s", file, err)
	}
	defer f.Close()
	result := make([]string, 0)
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("读取csv 失败 err:%s", err)
		}
		result = append(result, record[0]) //读取第一列内容
	}
	return result, nil
}

func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}

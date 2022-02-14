package image

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/src/service/fan_go_gin/config"
	"go_gin/src/service/fan_go_gin/model"
	"go_gin/src/service/fan_go_gin/utils"
	"go_gin/src/service/fan_go_gin/utils/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ImgListInfo struct {
	Total       int        `json:"imgTotal"`
	ImgInfos    []*ImgInfo `json:"imgInfos"`
	PreviewURLs []string   `json:"previewURLs"`
}

type ImgInfo struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	Fit          string `json:"fit"`
	FileSize     string `json:"fileSize"`
	CreateTime   string `json:"cTime"`
	CreateTimeTs int64  `json:"-"`
}

func UploadImage(ctx *gin.Context) {
	go func() {
		imgFileGlobal = FileGlobal{}
	}()
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

func ImgList(ctx *gin.Context) {
	page, err := getQueryInt(ctx, "page")
	if err != nil {
		logger.Errorf("page 参数不存在，默认为1 err:%s", err)
		page = 1
	}
	pageSize, err := getQueryInt(ctx, "pageSize")
	if err != nil {
		logger.Errorf("pageSize 参数不存在，默认为10 err:%s", err)
		pageSize = 10
	}
	logger.Infof("page:%d, pageSize:%d", page, pageSize)

	fullList, err := getFileList(config.Conf.ImagePathPre)
	if err != nil {
		logger.Errorf("获取文件失败 err:%s", err)
		utils.ReturnFailedJson(ctx, model.CodeSystemError, err.Error())
		return
	}

	resp := &ImgListInfo{
		Total: len(fullList),
	}

	resp.ImgInfos = getResultFile(page, pageSize)
	resp.PreviewURLs = getPreviewURLs(resp.ImgInfos)
	logger.Infof("返回数据:%+v", resp)
	utils.ReturnSuccessJson(ctx, resp)
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

	exists, err := IsExists(dstF)
	if err != nil {
		return err
	}
	if exists {
		old := dstF
		dstA := strings.Split(dst, ".")
		logger.Infof("dstA:%+v", dstA)
		if len(dstA) > 1 {
			dstF = fmt.Sprintf("%s.%d.%s",
				strings.Join(dstA[0:len(dstA)-1], "."),
				time.Now().Unix(), dstA[len(dstA)-1])
		}
		dstF = config.Conf.ImagePathPre + dstF
		logger.Warnf("文件%s存在,重命名为:%s", old, dstF)
	}
	out, err := os.Create(dstF)
	if err != nil {
		return err
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, src)
	return err
}

func getPreviewURLs(iis []*ImgInfo) []string {
	urls := make([]string, 0)
	for _, t := range iis {
		urls = append(urls, t.URL)
	}
	return urls
}

func IsExists(f string) (bool, error) {
	_, err := os.Stat(f)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var notExist = errors.New("not exits")

func getQueryInt(ctx *gin.Context, key string) (int, error) {
	str, ok := ctx.GetQuery(key)
	if !ok {
		return 0, notExist
	}
	return strconv.Atoi(str)
}

var imgFileGlobal = FileGlobal{}

type FileGlobal []*ImgInfo

func (g FileGlobal) Len() int {
	return len(g)
}
func (g FileGlobal) Less(i, j int) bool {
	return g[i].CreateTimeTs > g[j].CreateTimeTs
}
func (g FileGlobal) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func getFileList(path string) (FileGlobal, error) {
	if imgFileGlobal.Len() > 0 {
		return imgFileGlobal, nil
	}
	logger.Infof("开始读取目录:%s", path)
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	const fit = "contain"
	iis := make([]*ImgInfo, 0)
	for _, f := range fs {
		fileName := f.Name()
		if !strings.HasSuffix(fileName, "jpg") &&
			!strings.HasSuffix(fileName, "png") &&
			!strings.HasSuffix(fileName, "jpeg") {
			continue
		}
		temp := &ImgInfo{
			Name:         f.Name(),
			URL:          config.Conf.FileSysPre + f.Name(),
			Fit:          fit,
			FileSize:     getSizeStr(f.Size()),
			CreateTime:   f.ModTime().Format("2006-01-02 15:04:05"),
			CreateTimeTs: f.ModTime().Unix(),
		}
		iis = append(iis, temp)
	}
	sortedFile := FileGlobal(iis)
	sort.Sort(sortedFile) //按创建时间倒序排序
	imgFileGlobal = sortedFile
	return iis, nil
}

func getResultFile(page, pageSize int) []*ImgInfo {
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > len(imgFileGlobal) {
		end = len(imgFileGlobal)
	}
	return imgFileGlobal[start:end]
}

func getSizeStr(size int64) string {
	const kbint = 1024
	const mbint = 1024 * 1024
	if size < kbint {
		return fmt.Sprintf("%d B", size)
	} else if size < mbint {
		return fmt.Sprintf("%0.3f KB", float64(size)/kbint)
	} else {
		return fmt.Sprintf("%0.3f MB", float64(size)/mbint)
	}
}

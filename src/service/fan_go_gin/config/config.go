package config

import (
	"flag"
	"github.com/Unknwon/goconfig"
	"service/fan_go_gin/utils/logger"
)

var (
	Conf = new(Config)
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "cf", "./conf.ini", "默认配置文件路径")
}

// 初始化配置文件
func InitConfig() error {
	//confFile := "/Users/klook/Documents/klook/workspace_hf/go_gin/src/service/fan_go_gin/conf.ini"
	confFile := "./conf.ini"
	cf, err := goconfig.LoadConfigFile(confFile)
	if err != nil {
		return err
	}

	err = ParseConfig(cf, Conf)
	if err != nil {
		return err
	}
	Conf.Print()
	return nil
}

type Config struct {
	// 系统配置
	HttpPort int `config:"SYSTEM:port"`

	// 数据库配置
	MySQLHost string `config:"DB:mysql.host"`

	PdfBasePath string `config:"BUSINESS:download.pdf.basepath"`
	DownloadTask int `config:"BUSINESS:download.task"`
}

func (c *Config) Print() {
	logger.Infof("http.port:%d", c.HttpPort)
	logger.Infof("mysql.host:%s", c.MySQLHost)
	logger.Infof("download.pdf.basepath:%s", c.PdfBasePath)
	logger.Infof("download.task:%s", c.DownloadTask)
}

package config

import (
	"github.com/Unknwon/goconfig"
	"service/fan_go_gin/utils/logger"
)

var (
	Conf = new(Config)
)

// 初始化配置文件
func InitConfig() error {
	cf, err := goconfig.LoadConfigFile("/Users/klook/Documents/klook/workspace_hf/go_gin/src/service/fan_go_gin/conf.ini")
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

}

func (c *Config) Print() {
	logger.Infof("http.port:%d", c.HttpPort)
	logger.Infof("mysql.host:%s", c.MySQLHost)
}

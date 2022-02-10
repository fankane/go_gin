package config

import (
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"go_gin/src/service/fan_go_gin/utils/logger"
	"reflect"
	"strconv"
	"strings"
)

/**
把配置文件的配置数据 解析到 Config 结构体里
 */
func ParseConfig(c *goconfig.ConfigFile, config interface{}) error {

	if c == nil {
		return errors.New("goconfig.ConfigFile is nil")
	}

	if config == nil {
		return errors.New("待序列化的配置结构体为空")
	}

	value := reflect.ValueOf(config)

	valueElum := value.Elem() //conf 的值
	valueFieldQuantity := valueElum.NumField() //conf 的值的个数
	valueFieldType := valueElum.Type() //conf 的值的类型

	for i := 0;i<valueFieldQuantity;i++ {
		tempElumField := valueElum.Field(i)
		tempElumFieldType := valueFieldType.Field(i)
		filedTag := tempElumFieldType.Tag
		section, key := getSectionAndKey(filedTag.Get("config"))
		if section == "" || key == "" {
			logger.Infof("获取配置 config tag 为空")
			continue
		}
		configValue, err := getValueFromFile(c, section ,key)
		if err != nil {
			return fmt.Errorf("getValueFromFile err:%s", err)
		}
		err = setValue(tempElumField, configValue)
		if err != nil {
			return fmt.Errorf("配置赋值失败 err:%s", err)
		}
	}
	return nil
}

/**
从 struct 标签里面，获取到 配置项的 section 和 key
 */
func getSectionAndKey(tag string) (section, key string) {
	arr := strings.SplitN(tag,":", 3)
	if len(arr) <2 {
		return
	}
	section, key = arr[0], arr[1]
	return
}

/**
根据 section，key 从配置文件里面获取配置值
 */
func getValueFromFile(c *goconfig.ConfigFile, section, key string) (string, error) {
	if c == nil {
		return "", errors.New("configFile is nil")
	}

	// 根据 section 获取到 map[string]string 的 key-value 键值对
	sMap, err := c.GetSection(section)
	if err != nil {
		return "", fmt.Errorf("获取 section 失败 err:%s", err)
	}
	value, ok := sMap[key]
	if !ok {
		// key 不存在
		return "", nil
	}
	return value, nil
}

/**
设置 Config的字段具体值
配置文件里面的值从 string 类型，反射到具体字段类型
 */
func setValue(rf reflect.Value, value string) error {
	switch rf.Kind() {
	case reflect.Bool:
		rf.SetBool(getBoolFromStr(value))
	case reflect.String:
		rf.SetString(value)
	case reflect.Int, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Errorf("value:%s, 不是int, 无法赋值给:%s", value, rf.String())
		}
		rf.SetInt(intVal)
	default:
		return fmt.Errorf("不支持的类型:%s", rf.Kind().String())
	}
	return nil
}


func getBoolFromStr(v string) bool {
	if strings.ToLower(v) == "true" {
		return true
	}
	return false
}



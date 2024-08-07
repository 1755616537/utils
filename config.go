package utils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gyaml"
	"io/ioutil"
)

/*
type Config struct {}

func (Config)GetFuWuQiJava1()string  {
	return Util.GetConfig("FuWuQi.Java1").(string)
}
*/

var (
	_config *gjson.Json
)

type GetConfigs struct {
	//配置文件路径
	fileUrl string
}

// 加载配置文件数据
func (this *GetConfigs) Ini() error {
	if this.fileUrl == "" {
		this.fileUrl = "config.yml"
	}
	configByte, err := ioutil.ReadFile(this.fileUrl)
	if err != nil {
		return errors.New(fmt.Sprint("读取配置信息错误", err.Error()))
	}
	configYml, err := gyaml.ToJson(configByte)
	if err != nil {
		return errors.New(fmt.Sprint("解析配置信息错误", err.Error()))
	}

	_config, err = gjson.DecodeToJson(configYml)
	if err != nil {
		return errors.New(fmt.Sprint("解析配置信息错误", err.Error()))
	}

	return nil
}

// 读配置文件信息
func RunConfig() error {
	//是否是Linux环境
	filename := "config2.yml"
	{
		_LinuxOn, err := ZhengShiHuanJingOn()
		if err != nil {
			return err
		}
		if _LinuxOn {
			filename = "config.yml"
		}
	}

	configByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New(fmt.Sprint("读取配置信息错误", err.Error()))
	}
	configYml, err := gyaml.ToJson(configByte)
	if err != nil {
		return errors.New(fmt.Sprint("解析配置信息错误", err.Error()))
	}

	_config, err = gjson.DecodeToJson(configYml)
	if err != nil {
		return errors.New(fmt.Sprint("解析配置信息错误", err.Error()))
	}

	return nil
}

// 取指定配置信息
func GetConfig(name string) interface{} {
	if _config == nil {
		return nil
	}

	return _config.Get(name)
}

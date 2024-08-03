package utils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gyaml"
	"io/ioutil"
	"os"
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

// 读配置文件信息
func RunConfig() error {
	riJiAlL, _ := GetRiJi()
	riJi := riJiAlL[0]

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
		errValue := fmt.Sprint("读取配置信息错误 - ", err.Error())
		riJi.RiJiShuChuJingGaoFatal(errValue)
		return errors.New(errValue)
	}
	configYml, err := gyaml.ToJson(configByte)
	if err != nil {
		errValue := fmt.Sprint("解析配置信息错误 - ", err.Error())
		riJi.RiJiShuChuJingGaoFatal(errValue)
		return errors.New(errValue)
	}

	_config, err = gjson.DecodeToJson(configYml)
	if err != nil {
		errValue := fmt.Sprint("解析配置信息错误 - ", err.Error())
		riJi.RiJiShuChuJingGaoFatal(errValue)
		return errors.New(errValue)
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

// 是否是Linux环境
func ZhengShiHuanJingOn() (bool, error) {
	//获取当前目录路径
	str, err := os.Getwd()
	if err != nil {
		return false, errors.New("获取当前目录路径失败")
	}
	if str[:1] == "/" {
		return true, nil
	}
	return false, nil
}

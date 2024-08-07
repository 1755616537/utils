package utils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gyaml"
	"github.com/mdobak/go-xerrors"
	"io/ioutil"
	"log/slog"
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
	riJiAlL, _ := GetRiJi()
	riJi := riJiAlL[0]

	if this.fileUrl == "" {
		this.fileUrl = "config.yml"
	}
	configByte, err := ioutil.ReadFile(this.fileUrl)
	if err != nil {
		errValue := fmt.Sprint("读取配置信息错误 - ", err.Error())
		riJi.RiJiShuChuJingGaoFatal(errValue)
		slog.Error("读取配置信息错误")
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

// 原生error转换成携带错误信息的error
func ToXerror(err error) error {
	return xerrors.New(err)
}

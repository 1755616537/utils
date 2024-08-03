package utils

import (
	"errors"
	"time"
)

func StringTotime(timeLayout, toBeCharge string) (*time.Time, error) {
	loc, err := time.LoadLocation("Local") //重要：获取时区
	if err != nil {
		return nil, errors.New("时区获取失败")
	}
	theTime, err := time.ParseInLocation(toBeCharge, timeLayout, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return nil, errors.New("时间转换失败")
	}
	return &theTime, nil
}

package cache

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

// 缓存 限制
var Cache *cache.Cache

// 获取 Cache 缓存实例 限制
func GetCache() *cache.Cache {
	if Cache == nil {
		//创建一个默认1秒的缓存 每2秒清除过期项目
		Cache = cache.New(1*time.Second, 2*time.Second)
	}
	return Cache
}

// 设置缓存 公共 key=存储位置 value=需要存储内容
func SetHuanCun(标志key, 识别key string, value interface{}, 是否无限期 bool) error {
	var Time time.Duration
	if 是否无限期 {
		//使用无限期
		Time = cache.NoExpiration
	} else {
		//使用默认时间限制
		Time = cache.DefaultExpiration
	}
	GetCache().Set(fmt.Sprint(标志key, "-", 识别key), value, Time)
	return nil
}

// 获取缓存 公共 key=存储位置
func _GetHuanCun(标志key, 识别key string) (interface{}, error) {
	data, errbool := GetCache().Get(fmt.Sprint(标志key, "-", 识别key))
	if !errbool {
		return "", errors.New("获取失败(或存储位置无数据")
	}
	return data, nil
}

package utils

import (
	"github.com/mojocn/base64Captcha"
)

// 开辟一个验证码使用的存储空间
var store = base64Captcha.DefaultMemStore

// 获取验证码
func GetCaptcha() (string, string, error) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}

package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// 运行目录
func RunDirectory() (string, error) {
	var ex string
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// 运行程序时是否编程时的环境
func IfTestEnv() bool {
	var ex string
	ex, err := RunDirectory()
	if err != nil {
		panic(err)
	}
	if GetRStr(ex, 4) == "Temp" {
		return true
	}
	if GetRStr(ex, 10) == "tmp\\GoLand" {
		return true
	}
	return false
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

// 公钥加密
func RsaEncrypt2(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func RsaDecrypt2(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 计算 总页码 total=总数量 pageSize=多少个分一页的数量
func JiSuanZongYeMa(total, pageSize int) int {
	if total%pageSize != 0 {
		return (total / pageSize) + 1
	} else {
		return total / pageSize
	}
}

// 使用正则表达式清理字符串 [^a-zA-Z0-9]
func CleanString(input string) string {
	// Regular expression to match non-alphanumeric characters.
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	// Replace matched characters with an empty string.
	return re.ReplaceAllString(input, "")
}

// 转换uint32类型
func TyUint32(str string) (uint32, error) {
	dateUint64Value, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, errors.New("【时间类型转换失败】")
	}
	return uint32(dateUint64Value), nil
}

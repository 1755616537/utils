package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/grand"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// 32位随机字符串
func GetNonceStr(n int) string {
	return grand.Str("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890", n)
}

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

// 涨跌幅，计算
func CalculateChange(currentPrice float64, previousClosePrice float64) float64 {
	change := (currentPrice - previousClosePrice) / previousClosePrice * 100
	return change
}

// 涨停价 计算 10%是0.10
// fmt.Printf("%.2f", limitUpPrice)
func CalculateLimitUpPrice(price, limit float64) float64 {
	// 计算涨停价增值
	limitUpIncrement := price * limit
	// 四舍五入涨停价增值到最近的分位数
	limitUpIncrementRounded := math.Round(limitUpIncrement*100) / 100
	// 计算涨停价
	limitUpPrice := price + limitUpIncrementRounded
	return limitUpPrice
}

// 涨停价 计算
func CalculateLimitUpPrice2(price, limit float64) float64 {
	lastClose := Decimal(price)
	upStopPrice := Decimal(lastClose * (1.0000 + limit))
	return upStopPrice
}

// 跌停价 计算 10%是0.10
// fmt.Printf("%.2f", limitUpPrice)
func CalculateLimitDownPrice(price, limit float64) float64 {
	// 计算跌停价减少
	limitDownDecrement := price * limit
	// 四舍五入跌停价减少到最近的分位数
	limitDownDecrementRounded := math.Round(limitDownDecrement*100) / 100
	// 计算跌停价
	limitDownPrice := price - limitDownDecrementRounded
	return limitDownPrice
}

// 跌停价 计算
func CalculateLimitDownPrice2(price, limit float64) float64 {
	lastClose := Decimal(price)
	upStopPrice := Decimal(lastClose * (1.0000 - limit))
	return upStopPrice
}

// 保留两位数
func Float64To2u(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

// 保留两位数
func Float64To2u2(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

// 把一个[]string按N个一份分成map

func SplitSliceIntoChunks(slice []string, chunkSize int) map[int][]string {
	chunks := make(map[int][]string)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks[i/chunkSize] = slice[i:end]
	}
	return chunks
}

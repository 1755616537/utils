package utils

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gogf/gf/util/grand"
	"io/ioutil"
	"math/rand"
	"time"
)

// SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}

// 生成随机字符串
func GenerateSubId(_len int) string {
	b := make([]rune, _len)
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := range b {
		b[i] = letterRunes[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(letterRunes))]
	}
	return string(b)
}

// SHA256加密 >>> Base64加密
func SHA256WithRsaBase64(origData string, key []byte, keypath string) (sign string, err error) {
	if key == nil {
		key, err = ioutil.ReadFile(keypath)
		if err != nil {
			return "", err
		}
	}

	blocks, _ := pem.Decode(key)
	if blocks == nil || blocks.Type != "PRIVATE KEY" {
		fmt.Println("无法解码私钥")
		return

	}
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	if err != nil {
		return "", err
	}

	sign = base64.StdEncoding.EncodeToString(s)

	return sign, err
}

// AES-256-GCM解密
func RsaDecrypt(ciphertext, nonce2, associatedData2 string) (plaintext string, err error) {
	key := []byte("UJATIMB38cHO5X4ABekT4FZT0V7O0Pv3") //key是APIv3密钥，长度32位，由管理员在商户平台上自行设置的
	additionalData := []byte(associatedData2)
	nonce := []byte(nonce2)

	block, err := aes.NewCipher(key)
	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaindata, err := aesgcm.Open(nil, nonce, cipherdata, additionalData)
	//fmt.Println("plaintext: ", string(plaindata))

	return string(plaindata), err
}

// 验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 加密SHA统一格式 salt=显式(会出现在加密字符串里面) salt2=隐式(不会出现在加密字符串里面)
func JiaMiSHATongYiGeShi(data, salt, salt2 string) string {
	if salt == "" {
		//取16位随机字符串
		salt = grand.Str("qwertyuioplkjhgfdsazxcvbnm1234567890", 16)
	}
	//加密密码
	data = GetSHA256HashCode([]byte(data))
	data = GetSHA256HashCode([]byte(fmt.Sprint(data, salt)))
	if salt2 != "" {
		data = GetSHA256HashCode([]byte(fmt.Sprint(data, salt2)))
	}
	data = fmt.Sprint("$", "SHA", "$", salt, "$", data)
	return data
}

// 校验SHA统一格式 salt=显式(会出现在加密字符串里面) salt2=隐式(不会出现在加密字符串里面)
func JiaoYanSHATongYiGeShi(data, JiaMidata, salt, salt2 string) bool {
	//加密密码
	data = GetSHA256HashCode([]byte(data))
	data = GetSHA256HashCode([]byte(fmt.Sprint(data, salt)))
	if salt2 != "" {
		data = GetSHA256HashCode([]byte(fmt.Sprint(data, salt2)))
	}
	data = fmt.Sprint("$", "SHA", "$", salt, "$", data)
	if data == JiaMidata {
		return true
	}
	return false
}

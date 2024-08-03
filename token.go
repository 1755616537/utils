package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//token

type PayloadType struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	//用户ID
	UserID int `json:"UserID"`
	//用户名称
	UserName string `json:"UserName"`
	//加密的密码
	UserPassword string `json:"UserPassword"`
	//用户IP
	IP string `json:"IP"`
	//申请地址
	ShenQingDiZhi string `json:"ShenQingDiZhi"`
	//签发者
	QianFaZhe string `json:"QianFaZhe"`
	//权限
	QuanXian string `json:"QuanXian"`
	//自定义信息
	Data interface{} `json:"Data"`
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

// 获取Token
func SetToken(payload *PayloadType, salt string) (string, error) {
	//加密密码
	{
		if payload.UserPassword != "" {
			//哈希算法
			UserPasswordByte, err := bcrypt.GenerateFromPassword([]byte(payload.UserPassword), bcrypt.DefaultCost)
			if err != nil {
				return "", errors.New("加密密码出错")
			}
			//组合算法
			payload.UserPassword = JiaMiSHATongYiGeShi(string(UserPasswordByte), "", salt)
		}
	}

	if payload.ExpiresAt == 0 {
		payload.IssuedAt = time.Now().Unix()
		payload.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	} else {
		payload.ExpiresAt = time.Now().Unix() + (payload.ExpiresAt - payload.IssuedAt) + 1
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New("获取令牌失败-" + err.Error())
	}
	return signedToken, nil
}

// 校验Token
func GetToken(strToken string, context *gin.Context) (*PayloadType, error) {
	token, err := jwt.ParseWithClaims(strToken, &PayloadType{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if token == nil {
		if err != nil {
			return nil, errors.New("校验令牌失败,或已失效")
		}
		return nil, errors.New("校验令牌失败,或已失效")
	}
	claims, ok := token.Claims.(*PayloadType)
	if !ok {
		return nil, errors.New("无效令牌")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("无效令牌")
	}
	if err != nil {
		return claims, errors.New("校验令牌失败,或已失效")
	}
	if context != nil {
		//ip, _ := GetPublic().GetClientIp()
		if claims.IP != fmt.Sprintf("%s-%s", context.ClientIP(), "ip") {
			return claims, errors.New("校验令牌失败")
		}
	}
	return claims, nil
}

// 校验Token是否正确，不校验IP地址
func GetTokenString(token string) (*PayloadType, error) {
	tokenPayloadType, err := GetToken(token, nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return tokenPayloadType, nil
}

// 更新Token
func RefreshToken(strToken string) (string, error) {
	var context *gin.Context
	payload, err := GetToken(strToken, context)
	if err != nil {
		return "", err
	}
	Token, err := SetToken(payload, "")
	if err != nil {
		return "", err
	}
	if Token == "" {
		return "", errors.New("更新令牌失败")
	}
	return Token, nil
}

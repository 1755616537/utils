package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/encoding/gjson"
)

// 取post参值
func PostGetValue(c *gin.Context) (*gjson.Json, error) {
	//从缓存中取值
	RawDataByte, errBoll := c.Get("PostGetRawData")
	var (
		PostGetRawData []byte
		err            error
	)
	if !errBoll {
		//取post参值
		PostGetRawData, err = c.GetRawData()
		if err != nil {
			return nil, errors.New(fmt.Sprint("取参值失败", err.Error()))
		}
		fmt.Println("【地址】", c.FullPath(), "【数据被取出】", "===", string(PostGetRawData))
	} else {
		PostGetRawData = []byte(fmt.Sprint(RawDataByte))
	}
	//转换类型
	//{
	//	//校验Json格式
	//	var PostData map[string]interface{}
	//	err = json.Unmarshal(PostGetRawData, &PostData)
	//	if err != nil {
	//		return nil, errors.New(fmt.Sprint("参值转换json类型失败",err.Error()))
	//	}
	//}

	if !errBoll {
		//存储值,防止下次取不到
		c.Set("PostGetRawData", string(PostGetRawData))
	}

	PostGetRawDataJson, err := gjson.DecodeToJson(string(PostGetRawData))
	if err != nil {
		return nil, errors.New(fmt.Sprint("参值转换json类型2失败", err.Error()))
	}

	return PostGetRawDataJson, nil
}

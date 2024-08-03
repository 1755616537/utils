package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// json字符串转换map[string]
func Json_MapString(data string) (map[string]interface{}, error) {
	var mapData map[string]interface{}
	err := json.Unmarshal([]byte(data), &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

// HTTP请求 map类型 method=请求类型
func HTTPGet(method string, url string, data map[string]interface{}) (*http.Response, string, error) {
	var (
		bytesData []byte
		err       error
	)
	if data != nil {
		bytesData, err = json.Marshal(data)
		if err != nil {
			return nil, "", err
		}
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
	}
	client := &http.Client{}
	var bytesData2 io.Reader
	if data == nil {
		bytesData2 = nil
	} else {
		bytesData2 = bytes.NewReader(bytesData)
	}
	req, err := http.NewRequest(method, url, bytesData2)
	if err != nil {
		return nil, "", err
	}
	for i, i2 := range headers {
		req.Header.Add(i, i2)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	fmt.Println("HTTP请求【结果】===", string(body))
	return resp, string(body), nil
}

// HTTP请求 interface类型 method=请求类型
func HTTPGet2(method string, url string, data interface{}) (*http.Response, string, error) {
	var (
		bytesData []byte
		err       error
	)
	if data != nil {
		bytesData, err = json.Marshal(data)
		if err != nil {
			return nil, "", err
		}
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
	}
	client := &http.Client{}
	var bytesData2 io.Reader
	if data == nil {
		bytesData2 = nil
	} else {
		bytesData2 = bytes.NewReader(bytesData)
	}
	req, err := http.NewRequest(method, url, bytesData2)
	if err != nil {
		return nil, "", err
	}
	for i, i2 := range headers {
		req.Header.Add(i, i2)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return resp, string(body), nil
}

// HTTP请求 interface类型 method=请求类型 自定义headers
func HTTPGet2Headers(method string, url string, data interface{}, headers map[string]string) (*http.Response, string, error) {
	var (
		bytesData []byte
		err       error
	)
	if data != nil {
		bytesData, err = json.Marshal(data)
		if err != nil {
			return nil, "", err
		}
	}
	client := &http.Client{}
	var bytesData2 io.Reader
	if data == nil {
		bytesData2 = nil
	} else {
		bytesData2 = bytes.NewReader(bytesData)
	}
	req, err := http.NewRequest(method, url, bytesData2)
	if err != nil {
		return nil, "", err
	}
	for i, i2 := range headers {
		req.Header.Add(i, i2)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return resp, string(body), nil
}

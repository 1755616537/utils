package utils

import (
	"errors"
	"net"
)

// 获取访问者外网IP
func GetClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if Inet, ok := address.(*net.IPNet); ok && !Inet.IP.IsLoopback() {
			if Inet.IP.To4() != nil {
				return Inet.IP.String(), nil
			}
		}
	}
	return "", errors.New("找不到客户端ip地址!")
}

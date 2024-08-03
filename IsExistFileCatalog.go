package utils

import "os"

// 是否存在文件或目录
func IsExistFileCatalog(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

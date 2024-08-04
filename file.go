package utils

import (
	"bufio"
	"fmt"
	cyanfile "github.com/dablelv/cyan/file"
	"io"
	"log"
	"os"
	"strings"
)

// 读取文本
func LineByLine(file string) ([]string, error) {

	var err error

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var res []string
	for {
		line, err := r.ReadString('\n')
		line = CleanString(line)
		if err == io.EOF {
			if line != "" {
				res = append(res, line)
			}
			break
		} else if err != nil {
			log.Printf("error reading file %s", err)
			break
		}
		res = append(res, line)
	}
	return res, nil
}

// 保存文本到文件
func Setfile(data, fileNamem string) error {
	dstFile, err := os.Create(fileNamem)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(data)
	if err != nil {
		return err
	}
	return err
}

// 读入文件
func Getfile(fileNamem string) (string, error) {
	fp, err := os.OpenFile(fileNamem, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
	if err != nil {
		return "", err
	}
	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。

	var data []byte
	_, err = fp.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 获取目录下所有文件和子目录名称（不会递归）
func ListDir(dir string) ([]string, error) {
	return cyanfile.ListDir(dir)
}

// 递归列出目录中的所有文件或目录路径。如果cur为真，则结果将包括当前目录。
// 请注意，如果子目录是符号链接，则GetDirAllEntryPaths不会跟在符号链接后面。
func ListDirEntryPaths(dir string, cur bool) ([]string, error) {
	return cyanfile.ListDirEntryPaths(dir, cur)
}

// 递归列出目录中的所有文件或目录路径。如果cur为真，则结果将包括当前目录。
func ListDirEntryPathsSymlink(dir string, cur bool) ([]string, error) {
	return cyanfile.ListDirEntryPathsSymlink(dir, cur)
}

// 获取目录下所有文件和子目录名称（不会递归） 指定后缀名
func ListDirType(dir string, suffix string) ([]string, error) {
	list, err := cyanfile.ListDir(dir)
	if err != nil {
		return nil, err
	}
	var resList []string
	for i := 0; i < len(list); i++ {
		if strings.HasSuffix(list[i], suffix) {
			resList = append(resList, list[i])
		}
	}
	return resList, nil
}

// 去除后缀名
func RemoveSuffix(str string) string {
	strArr := strings.Split(str, ".")
	if len(strArr) > 0 {
		str = ""
		for i := 0; i < len(strArr)-1; i++ {
			if i+1 == len(strArr)-1 {
				str = fmt.Sprint(str, strArr[i])
			} else {
				str = fmt.Sprint(str, strArr[i], ".")
			}
		}
	}

	return str
}

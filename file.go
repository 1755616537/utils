package utils

import (
	"bufio"
	"io"
	"log"
	"os"
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
		if err == io.EOF {
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

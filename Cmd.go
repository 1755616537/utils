package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// 命令行调用 字符串分割
func CmdString(com string) (string, []string) {
	args := strings.Split(com, " ")
	return args[0], args[1:]
}

// 直接在当前目录使用并返回结果
func Cmd(commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	//fmt.Println("Cmd", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

// 在命令位置使用并返回结果
func CmdAndChangeDir(dir string, commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	//fmt.Println("CmdAndChangeDir", dir, cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

// 在命令位置使用并实时输出每行结果
func CmdAndChangeDirToShow(dir string, commandName string, params []string) error {
	cmd := exec.Command(commandName, params...)
	//fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println("CMD===【输出】>>>", line)
	}
	err = cmd.Wait()
	return err
}

// 在命令位置使用并实时写入每行结果到文件
func CmdAndChangeDirToFile(fileName, dir, commandName string, params []string) error {
	var f *os.File
	//判断文件是否存在
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			return err
		}
		defer f.Close()
	}
	cmd := exec.Command(commandName, params...)
	//fmt.Println("CmdAndChangeDirToFile", dir, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		_, err = f.WriteString(line) //写入文件(字节数组)
		_ = f.Sync()
	}
	//_, err = f.WriteString("=================处理完毕========================") //写入文件(字节数组)
	//_ = f.Sync()
	err = cmd.Wait()
	return err
}

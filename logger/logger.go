package logger

import (
	"context"
	"errors"
	"fmt"
	"github.com/1755616537/utils"
	"io"
	"log/slog"
	"os"
	"time"
)

var (
	//日记目录
	logUrl string = "./logFile"
	//是否打印日志到文件
	onLogFile bool = false
	//log文件
	fileOpenFile *os.File = os.Stdout
	//使用error错误栈打印
	//slog.ErrorContext(ctx, "err", slog.Any("error", err))
	printErrorStack bool = true
)

func init() {
	//err := IniLog()
	//if err != nil {
	//	slog.Error(err.Error())
	//}
}

func Exit() error {
	if fileOpenFile != os.Stdout {
		return fileOpenFile.Close()
	}
	return nil
}

// 初始化slog配置
func IniLog() error {
	err := Exit()
	if err != nil {
		return err
	}

	var appEnv = os.Getenv("APP_ENV")

	getTime := time.Now().Format("2006-01-02-15-04-05")
	//初始化文件保存地址
	url := fmt.Sprintf("%s/%s.log", logUrl, getTime)

	if onLogFile {
		if !utils.IsExistFileCatalog(logUrl) {
			//创建目录
			err := os.MkdirAll(logUrl, os.ModePerm)
			if err != nil {
				return errors.New(fmt.Sprint("MkdirLogUrlErr", err.Error()))
			}
		}
		wfileOpenFile, err := os.Create(url)
		//wfileOpenFile, err := os.OpenFile(url+".log", os.O_APPEND|os.O_CREATE, 666)
		if err != nil {
			slog.Error("OpenFileLogUrlErr", err)
		} else {
			fileOpenFile = wfileOpenFile
		}
	}

	opts := slog.HandlerOptions{
		AddSource: printErrorStack,
		Level:     slog.LevelDebug,
		//ReplaceAttr: replaceAttr,
	}

	var logger *slog.Logger
	if appEnv == "development" {
		logger = slog.New(slog.NewTextHandler(fileOpenFile, &opts))
	} else {
		//appEnv == "production"
		popts := PrettyHandlerOptions{
			SlogOpts: opts,
		}
		logger = slog.New(NewPrettyHandler(fileOpenFile, popts))

		//customHandler := &CustomHandler{}
		//logger = slog.New(customHandler)

		//logger = slog.New(slog.NewJSONHandler(fileOpenFile, &opts))
	}

	//使用 SetDefault() 方法还会改变 log 包使用的默认 log.Logger。
	//这种行为允许利用旧 log 包的现有应用程序无缝过渡到结构化日志记录。
	slog.SetDefault(logger)

	return nil
}

// 是否打印 error错误栈
func SetPrintErrorStack(b bool) {
	printErrorStack = b
}

// 日记目录
func SetLogUrl(url string) {
	logUrl = url
}

// 日记目录
func GetLogUrl() string {
	return logUrl
}

// 是否打印日志到文件
func SetOnLogFile(on bool) {
	onLogFile = on
}

// log文件IoWriter
func GetIoWriter() io.Writer {
	return fileOpenFile
}

// ctx := context.Background().
// 可以使用类似 xerrors 的库来创建带有堆栈跟踪的错误.
// err := xerrors.New("something happened").
func ErrorContext(ctx context.Context, ctxMsg string, err error) {
	slog.ErrorContext(ctx, ctxMsg, slog.Any("error", err))
}

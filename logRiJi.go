package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

//日记

const (
	//日记目录
	logUrl = "log"
	//输出前缀名
	QianZhuiMing string = "[日记系统]"
)

var (
	//RiJi日记池
	_riJi []*RiJi
)

type RiJi struct {
	//日记目录
	logUrl string
	//输出前缀名
	QianZhuiMing string
	//日记文件路径
	url string
	//创建时间
	Time string
	//当前在Gin池的位置
	WeiShu int
	//gin文件
	FileCreate *os.File
	//log文件
	FileOpenFile *os.File
	//正常日记输出
	LoggerZhengChang *log.Logger
	//提示日记输出
	LoggerTiShi *log.Logger
	//警告日记输出
	LoggerJingGao *log.Logger
}

// 启动日记
func RunRiJi(onFile bool) (*RiJi, error) {
	var riJi *RiJi

	riJi.RiJiShuChuTiShiPrintln(QianZhuiMing, "初始化中...")

	//获取当前时间
	getTime := time.Now().Format("2006-01-02-15-04-05")

	//初始化文件保存地址
	url := fmt.Sprintf("./%s/%s.log", logUrl, getTime)

	//gin日记只能初始化一次
	var FileCreate *os.File
	if len(_riJi) == 0 {
		//禁用控制台颜色禁用控制台中的颜色输出。
		gin.DisableConsoleColor()
		//控制台中的强制颜色输出。
		//gin.ForceConsoleColor()
		if onFile {
			//是否存在日记目录=>不存在就创建
			if !IsExistFileCatalog(logUrl) {
				//创建目录
				err := os.Mkdir(fmt.Sprintf("./%s", logUrl), os.ModePerm)
				if err != nil {
					return nil, errors.New("创建日记目录失败-" + err.Error())
				}
			}
			// 创建记录日志的文件
			fileCreate, err := os.Create(url)
			if err != nil {
				return nil, errors.New("打开日记(控制台)Url失败-" + err.Error())
			}
			FileCreate = fileCreate
			// 将日志同时写入文件和控制台
			gin.DefaultWriter = io.MultiWriter(fileCreate, os.Stdout)
		}
	} else {
		FileCreate = _riJi[0].FileCreate
	}

	var FileOpenFile *os.File
	var LoggerZhengChang, LoggerTiShi, LoggerJingGao *log.Logger
	if onFile {
		//打开文件
		fileOpenFile, err := os.OpenFile(url+".log", os.O_APPEND|os.O_CREATE, 666)
		if err != nil {
			return nil, errors.New("打开日记(log)Url失败-" + err.Error())
		}
		FileOpenFile = fileOpenFile
		//初始化日记格式
		LoggerZhengChang = log.New(fileOpenFile, "[ZhengChang][正常]", log.Ldate|log.Lmicroseconds|log.Lshortfile) // 日志文件格式:log包含时间及文件行数
		LoggerTiShi = log.New(fileOpenFile, "[TiShi][提示]", log.Ldate|log.Lmicroseconds|log.Lshortfile)           // 日志文件格式:log包含时间及文件行数
		LoggerJingGao = log.New(fileOpenFile, "[JingGao][警告]", log.Ldate|log.Lmicroseconds|log.Lshortfile)       // 日志文件格式:log包含时间及文件行数
	}

	riJi.RiJiShuChuTiShiPrintln(QianZhuiMing, "初始化完毕!")

	//获取当前池中的数量
	lenGin := len(_riJi)
	if lenGin == 0 {
		lenGin = 1
	} else {
		lenGin = lenGin + 1
	}

	riJi = &RiJi{
		//日记目录
		logUrl: logUrl,
		//输出前缀名
		QianZhuiMing: QianZhuiMing,
		//日记文件路径
		url: url,
		//创建时间
		Time: getTime,
		//当前在Gin池的位置
		WeiShu: lenGin,
		//gin文件
		FileCreate: FileCreate,
		//log文件
		FileOpenFile: FileOpenFile,
		//正常日记输出
		LoggerZhengChang: LoggerZhengChang,
		//提示日记输出
		LoggerTiShi: LoggerTiShi,
		//警告日记输出
		LoggerJingGao: LoggerJingGao,
	}

	_riJi = append(_riJi, riJi)

	return riJi, nil
}

// 获取RiJi
func GetRiJi() ([]*RiJi, error) {
	if _riJi != nil {
		if len(_riJi) != 0 {
			return _riJi, nil
		}
	}
	var riJi []*RiJi

	riJi0 := &RiJi{
		//日记目录
		logUrl: "",
		//输出前缀名
		QianZhuiMing: "",
		//日记文件路径
		url: "",
		//创建时间
		Time: "",
		//当前在Gin池的位置
		WeiShu: -1,
		//gin文件
		FileCreate: nil,
		//log文件
		FileOpenFile: nil,
		//正常日记输出
		LoggerZhengChang: nil,
		//提示日记输出
		LoggerTiShi: nil,
		//警告日记输出
		LoggerJingGao: nil,
	}

	riJi = append(riJi, riJi0)

	return riJi, errors.New("未初始化过一个RiJi日记")
}

// 移除RiJi
func (riJi *RiJi) DelRiJi() error {
	//声明一个互斥锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	if riJi.WeiShu == -1 {
		return errors.New("默认RiJi无法移除")
	}

	//判断池中数量是否比传递过来的位置要小,防止报错
	if len(_riJi) < riJi.WeiShu {
		return errors.New("RiJi池中数量比传递过来的数位要小")
	}
	err := riJi._Close()
	if err != nil {
		return err
	}
	//从Gin池中移除
	_riJi = append(_riJi[:riJi.WeiShu-1], _riJi[riJi.WeiShu:]...)

	//重新排序Gin池中的WeiShu位置
	for i := 0; i < len(_riJi); i++ {
		_riJi[i].WeiShu = i + 1
	}

	return nil
}

// 关闭资源RiJi 内部方法
func (riJi *RiJi) _Close() error {
	//程序关闭后关闭打开文件-gin日记
	if riJi.FileCreate != nil && riJi.Time != "" {
		//判断是否剩最后一个
		if len(_riJi) == 1 {
			err := riJi.FileCreate.Close()
			if err != nil {
				riJi.RiJiShuChuJingGaoPrintln(QianZhuiMing, "关闭打开文件-gin日记失败", fmt.Sprint("初始化时间:", riJi.Time), err.Error())
			} else {
				riJi.RiJiShuChuTiShiPrintln(QianZhuiMing, fmt.Sprint("初始化时间:", riJi.Time), "释放gin日记文件资源成功")
			}
		} else {
			riJi.RiJiShuChuTiShiPrintln(QianZhuiMing, fmt.Sprint("初始化时间:", riJi.Time), "释放gin日记文件资源成功")
		}
	}
	//程序关闭后关闭打开文件-log日记
	if riJi.FileOpenFile != nil && riJi.Time != "" {
		err := riJi.FileOpenFile.Close()
		if err != nil {
			riJi.RiJiShuChuJingGaoPrintln(QianZhuiMing, fmt.Sprint("初始化时间:", riJi.Time), "关闭打开文件-log日记失败", err.Error())
		} else {
			riJi.RiJiShuChuTiShiPrintln(QianZhuiMing, fmt.Sprint("初始化时间:", riJi.Time), "释放log日记文件资源成功")
		}
	}
	return nil
}

// 正常日记输出Println 到控制台和log日记
func (riJi *RiJi) RiJiShuChuZhengChangPrintln(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	if riJi != nil {
		if riJi.LoggerZhengChang != nil {
			riJi.LoggerZhengChang.Println(value...)
		}
	}
	value = append([]interface{}{"[ZhengChang][正常]"}, value...)
	GetDaiYanSeShuChu().ColorPrintlnPrint(FontColor.Green, true, true, value...)
}

// 提示日记输出Println 到控制台和log日记
func (riJi *RiJi) RiJiShuChuTiShiPrintln(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	if riJi != nil {
		if riJi.LoggerTiShi != nil {
			riJi.LoggerTiShi.Println(value...)
		}
	}
	value = append([]interface{}{"[TiShi][提示]"}, value...)
	GetDaiYanSeShuChu().ColorPrintlnPrint(FontColor.Yellow, true, true, value...)
}

// 警告日记输出Println 到控制台和log日记
func (riJi *RiJi) RiJiShuChuJingGaoPrintln(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	if riJi != nil {
		if riJi.LoggerJingGao != nil {
			riJi.LoggerJingGao.Println(value...)
		}
	}
	value = append([]interface{}{"[JingGao][警告]"}, value...)
	GetDaiYanSeShuChu().ColorPrintlnPrint(FontColor.Red, true, true, value...)
}

// 正常日记输出Fatal=>输出完后结束程序 到控制台和log日记
func (riJi *RiJi) RiJiShuChuZhengChangFatal(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	log.Println(value...)
	if riJi != nil {
		if riJi.LoggerZhengChang != nil {
			riJi.LoggerZhengChang.Fatal(value...)
			return
		}
	}
	log.Fatal(value...)
}

// 提示日记输出Fatal=>输出完后结束程序 到控制台和log日记
func (riJi *RiJi) RiJiShuChuTiShiFatal(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	log.Println(value...)
	if riJi != nil {
		if riJi.LoggerTiShi != nil {
			riJi.LoggerTiShi.Fatal(value...)
			return
		}
	}
	log.Fatal(value...)
}

// 警告日记输出Fatal=>输出完后结束程序 到控制台和log日记
func (riJi *RiJi) RiJiShuChuJingGaoFatal(value ...interface{}) {
	//value=append([]interface{}{QianZhuiMing},value...)
	log.Println(value...)
	if riJi != nil {
		if riJi.LoggerJingGao != nil {
			riJi.LoggerJingGao.Fatal(value...)
			return
		}
	}
	log.Fatal(value...)
}

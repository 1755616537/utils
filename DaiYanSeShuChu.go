package utils

import (
	"fmt"
	"sync"
	"time"
)

//带颜色字体控制台输出

var (
	//kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	//proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	//CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	_daiYanSeShuChu *DaiYanSeShuChu
)

type Color struct {
	Black       int // 黑色
	Blue        int // 蓝色
	Green       int // 绿色
	Cyan        int // 青色
	Red         int // 红色
	Purple      int // 紫色
	Yellow      int // 黄色
	LightGray   int // 淡灰色（系统默认值）
	Gray        int // 灰色
	LightBlue   int // 亮蓝色
	LightGreen  int // 亮绿色
	LightCyan   int // 亮青色
	LightRed    int // 亮红色
	LightPurple int // 亮紫色
	LightYellow int // 亮黄色
	White       int // 白色
}

type DaiYanSeShuChu struct {
}

func GetDaiYanSeShuChu() *DaiYanSeShuChu {
	return _daiYanSeShuChu
}

// 输出有颜色的字体 输出完成后不会恢复白色将继续输出当前颜色
func (daiYanSeShuChu *DaiYanSeShuChu) ColorPrintln(color int, timeBool bool, callerBool bool, value ...interface{}) {
	//声明一个互斥锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	//handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))

	if callerBool {
		mutex.Unlock()
		//添加调用文件行
		//var ok bool
		//_, file, line, ok := runtime.Caller(3)
		//if !ok {
		//	file = "???"
		//	line = 0
		//}
		//value=append([]interface{}{fmt.Sprint("\x1b[36;1m[",file,"^",line,"行]\x1b[0m")},value...)
		mutex.Lock()
	}

	if timeBool {
		//添加时间
		t := time.Now()
		value = append([]interface{}{fmt.Sprint("\u001B[33m[", t.Format("2006/01/02 15:04:05."), t.UTC().Nanosecond()/1e3, "]\u001B[0m")}, value...)
	}

	fmt.Println(value...)

	//_, _, _ = CloseHandle.Call(handle)
}

// 输出有颜色的字体 输出完成后恢复白色字体
func (daiYanSeShuChu *DaiYanSeShuChu) ColorPrintlnPrint(color int, timeBool bool, callerBool bool, value ...interface{}) {
	daiYanSeShuChu.ColorPrintln(color, timeBool, callerBool, value...)
	daiYanSeShuChu.ColorPrint(FontColor.White, false, false, "\r")
}

// 输出有颜色的字体 输出完成后不会恢复白色将继续输出当前颜色
func (daiYanSeShuChu *DaiYanSeShuChu) ColorPrint(color int, timeBool bool, callerBool bool, value ...interface{}) {
	//声明一个互斥锁
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	//handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))

	if callerBool {
		mutex.Unlock()
		//添加调用文件行
		//var ok bool
		//_, file, line, ok := runtime.Caller(3)
		//if !ok {
		//	file = "???"
		//	line = 0
		//}
		//value=append([]interface{}{file,line},value...)
		mutex.Lock()
	}

	if timeBool {
		//添加时间
		t := time.Now()
		value = append([]interface{}{fmt.Sprint(t.Format("2006/01/02 15:04:05."), t.UTC().Nanosecond()/1e3)}, value...)
	}

	fmt.Print(value...)

	//_, _, _ = CloseHandle.Call(handle)
}

package Ffmpeg

import (
	"errors"
	"fmt"
	"os"
	"strings"
	utils "util"
)

// 是否开启日记
func _ConfigOn() string {
	if true {
		return ""
	}
	return " -loglevel quiet"
}

// 判断文件或目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 分割视频成图片
func FenGeShiPinChengTuPian(mp4, path, s string) error {
	cmdStr := fmt.Sprint("ffmpeg", _ConfigOn(), " -ss 00:00 -i ", mp4, " -f image2 -r ", s, " -c copy ", path, "%8d.jpg")
	args := strings.Split(cmdStr, " ")
	msg, err := utils.Cmd(args[0], args[1:])
	if err != nil {
		return errors.New(fmt.Sprintf("videoMerge failed, %v, output: %v\n", err, msg))
	}
	return nil
}

// 执行自定义脚本文件 txt=脚本文件路径 mp4=生成文件路径
func ZhiXingZiDingYiMingLingWenJian(txt, mp4 string) error {
	cmdStr := fmt.Sprint("ffmpeg", _ConfigOn(), " -f concat -safe 0 -i ", txt, " ", "-c copy -vsync vfr -pix_fmt yuv420p ", mp4)
	args := strings.Split(cmdStr, " ")
	msg, err := utils.Cmd(args[0], args[1:])
	if err != nil {
		return errors.New(fmt.Sprintf("videoMerge failed, %v, output: %v\n", err, msg))
	}
	return nil
}

// 视频合成音频 mp4=视频路径 mp3=生成文件路径 mp42=生成文件路径
func ShiPinHeChengYinPin(mp4, mp3, mp42 string) error {
	cmdStr := fmt.Sprint("ffmpeg", _ConfigOn(), " -i ", mp4, " -i ", mp3, " ", "-c:v copy -af apad -shortest ", mp42)
	args := strings.Split(cmdStr, " ")
	msg, err := utils.Cmd(args[0], args[1:])
	if err != nil {
		return errors.New(fmt.Sprintf("videoMerge failed, %v, output: %v\n", err, msg))
	}
	return nil
}

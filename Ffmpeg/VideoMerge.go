package Ffmpeg

import (
	"errors"
	"fmt"
	"github.com/1755616537/utils"
	"strings"
)

// 视频合成
func VideoMerge(in []string, out string) error {
	cmdStr := fmt.Sprintf("ffmpeg%s -i concat:%s -loglevel quiet -c copy -absf aac_adtstoasc -movflags faststart %s",
		_ConfigOn(), strings.Join(in, "|"), out)
	args := strings.Split(cmdStr, " ")
	msg, err := utils.Cmd(args[0], args[1:])
	if err != nil {
		return errors.New(fmt.Sprintf("videoMerge failed, %v, output: %v\n", err, msg))
	}
	return nil
}

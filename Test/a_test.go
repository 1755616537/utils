package Test

import (
	"fmt"
	"github.com/1755616537/utils"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	timeArr, err := utils.DateAoArr(time.Now(), 20)
	if err != nil {
		return
	}
	fmt.Println(timeArr)
}

func Test_file(t *testing.T) {
	line, err := utils.LineByLine("C:\\Users\\17556\\Desktop/1.txt")
	fmt.Println(line, len(line), err)
}

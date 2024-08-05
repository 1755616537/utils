package Test

import (
	"fmt"
	"github.com/1755616537/utils"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	timeArr, err := utils.DateAoArr(time.Now(), 20, true)
	fmt.Println(timeArr, err)
	timeArr, err = utils.DateAoArr(time.Now(), 20, false)
	fmt.Println(timeArr, err)
}

func Test_file(t *testing.T) {
	line, err := utils.LineByLine("C:\\Users\\17556\\Desktop/1.txt")
	fmt.Println(line, len(line), err)
}

func Test_file2(t *testing.T) {
	fmt.Println(utils.RemoveSuffix("600501.a.json"))
}

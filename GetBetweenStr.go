package utils

import (
	"fmt"
	"strings"
)

// 取文本（字符串）中间
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

// 取最右文本
func GetRStr(str string, l int) string {
	runeStr := []rune(str)
	if len(runeStr) > 0 && len(str) >= l {
		return fmt.Sprint(string(runeStr[len(runeStr)-l:]))
	}
	return ""
}

// 取最左文本
func GetLStr(str string, l int) string {
	runeStr := []rune(str)
	if len(runeStr) > 0 && len(str) >= l {
		return fmt.Sprint(string(runeStr[:l]))
	}
	return ""
}

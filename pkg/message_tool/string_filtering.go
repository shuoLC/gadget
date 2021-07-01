package message_tool

import (
	"strings"
	"sync"
	"tour/pkg/transformation"
)

var lock sync.RWMutex

// Filter 字符串过滤空格换行
func Filter(str string) string {
	lock.Lock()
	result := transformation.StrTo(str).String()
	result = strings.ToTitle(strings.ReplaceAll(result," ",""))
	result = strings.ReplaceAll(result,"\n","")
	result = strings.ReplaceAll(result,"\r","")
	lock.Unlock()
	return result
}
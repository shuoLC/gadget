package transformation

import (
	"fmt"
	"strconv"
)

// Hex2Dec 十六进制字符串转十进制 int
func Hex2Dec(val string) (int,error) {
	n,err := strconv.ParseUint(val,16,32)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// MustHex2Dec 十六进制字符串转十进制 int 不返回 error
func MustHex2Dec(val string) int {
	v,_ := Hex2Dec(val)
	return v
}

// StrHex2Dec 16进制转10进制字符串
func StrHex2Dec(val string) string {
	n,err := strconv.ParseUint(val,16,64)
	if err != nil {
		fmt.Println(err)
	}
	res := fmt.Sprintf("%d",n)
	num := len(res)
	startNum := 16-num
	for startNum > 0 {
		res = "0"+res
		startNum--
	}
	return res
}
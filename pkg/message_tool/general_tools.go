package message_tool

import (
	"github.com/spf13/cast"
	"strconv"
	"strings"
	"tour/pkg/transformation"
)

// ByteInversion 字节反转
func ByteInversion(data string) string {
	data = strings.ReplaceAll(data, " ", "")
	var result []string
	length := len(data)
	num := 0
	for num < length {
		s := data[num : num+2]
		result = append([]string{s}, result...)
		num = num + 2
	}
	return strings.Join(result, "")
}

// SnInversion 集中器地址计算
func SnInversion(data string) string {
	return ByteInversion(data[0:4]) + ByteInversion(data[4:8])
}

// ReCalculatedAddress 测量点计算
func ReCalculatedAddress(position string) int {
	right, _ := strconv.ParseUint(position[2:4], 16, 32)
	left := position[0:2]
	var rLeft int
	switch left {
	case "01":
		rLeft = 1
	case "02":
		rLeft = 2
	case "04":
		rLeft = 3
	case "08":
		rLeft = 4
	case "10":
		rLeft = 5
	case "20":
		rLeft = 6
	case "40":
		rLeft = 7
	case "80":
		rLeft = 8
	default:
		rLeft = 8
	}
	res := (int(right)-1)*8 + rLeft
	return res
}

// BaseCalculation 数据标识解析 每两位倒序 -33
func BaseCalculation(data string) string {
	var result []string
	length := len(data)
	num := 0
	for num < length {
		s := data[num : num+2]
		i1 := transformation.StrTo(s[0:1]).MustInt64() - 3
		i2 := transformation.StrTo(s[1:2]).MustInt64() - 3
		str := cast.ToString(i1) + cast.ToString(i2)
		switch str {
		case "0-1":
			str = "ff"
		case "-10":
			str = "ff"
		case "-1-1":
			str = "ff"
		}
		result = append([]string{str}, result...)
		num = num + 2
	}
	return strings.Join(result, "")
}

// TimeStitching 年 月 日 时 分 秒 拼接 无逆转
func TimeStitching(data string) string {
	str := data[0:2] + "年" + data[2:4] + "月" + data[4:6] + "日" + data[6:8] + "时" + data[8:10] + "分" + data[10:12] + "秒"
	return str
}

// TimeStitchingF 年 月 日 时 分 时间拼接
func TimeStitchingF(data string) string {
	reTime := ByteInversion(data)
	return reTime[0:2] + "年" + reTime[2:4] + "月" + reTime[4:6] + "日" + reTime[6:8] + "时" + reTime[8:10] + "分"
}

// TimeStitchingR 年 月 日 时间拼接
func TimeStitchingR(data string) string {
	reTime := ByteInversion(data)
	return reTime[0:2] + "年" + reTime[2:4] + "月" + reTime[4:6] + "日"
}

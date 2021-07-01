package message_tool

import (
	"github.com/spf13/cast"
	"strconv"
	"strings"
	"sync"
	"tour/pkg/mutex_map"
	"tour/pkg/transformation"
)

// ReadCurrentData 645 数据区域计算
func ReadCurrentData(data string, numb, position int) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	var relic []string
	str := BaseCalculation(data)
	num := 0
	length := len(str)
	for num < length {
		s := str[num : num+numb]
		if length > position {
			s = s[0:position] + "." + s[position:]
		} else {
			s = s[0:]
		}
		relic = append([]string{s}, relic...)
		num = num + numb
	}
	for k, v := range relic {
		result.Set(k, v)
	}
	return result.Data()
}

// LastOpeningRecord 上次开盖记录
func LastOpeningRecord(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	var slice []string
	str := BaseCalculation(data)
	num := 0
	length := len(str)
	for num < length {
		s := TimeStitching(str[num : num+12])
		slice = append([]string{s}, slice...)
		num = num + 12
	}
	for k, v := range slice {
		result.Set(k, v)
	}
	return result.Data()
}

// PowerFailure 掉电事件
func PowerFailure(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	start := BaseCalculation(data[0:12])
	end := BaseCalculation(data[12:24])
	result.Set(0, TimeStitching(start))
	result.Set(1, TimeStitching(end))
	return result.Data()
}

// TotalTimesOfOverCurrent 过流总次数
func TotalTimesOfOverCurrent(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	var slice []string
	str := BaseCalculation(data)
	num := 0
	length := len(str)
	for num < length {
		slice = append([]string{str[num : num+2]}, slice...)
		num = num + 2
	}
	result.Set(0, strings.Join(slice, ""))
	return result.Data()
}

// OverCurrentTime 过流时刻
func OverCurrentTime(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	result.Set(0, TimeStitching(BaseCalculation(data)))
	return result.Data()
}

// LastHourData 上次整点冻结数据块
func LastHourData(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	str := strings.ReplaceAll(data, "DD", "")
	timer := str[0:10]
	totalTimesOfOverCurrent := TotalTimesOfOverCurrent(timer)
	timer = TimeStitchingF(totalTimesOfOverCurrent[0])
	result.Mp = ReadCurrentData(str[10:], 8, 6)
	result.Set(2, timer)
	return result.Data()
}

// PowerDataBlock 功率计算
func PowerDataBlock(data string, numb, position int) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	var slice []string
	str := BaseCalculation(data)
	num := 0
	length := len(str)
	for num < length {
		s := str[num : num+numb]
		if length > position {
			s = s[0:position] + "." + s[position:]
		} else {
			s = s[0:]
		}
		slice = append([]string{s}, slice...)
		num = num + numb
	}
	for k, v := range slice {
		result.Set(k, v)
	}
	return result.Data()
}

// TrippingAndClosingEvents 跳合闸事件
func TrippingAndClosingEvents(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	str := BaseCalculation(data[:12])
	str = str[0:2] + "年" + str[2:4] + "月" + str[4:6] + "日" + str[6:8] + "时" + str[8:10] + "分" + str[10:12] + "秒"
	result.Mp = ReadCurrentData(data[20:68], 8, 6)
	result.Set(6, str)
	return result.Data()
}

// MaximumDemandAndTime 上一日正向有功最大需量及发生时间
func MaximumDemandAndTime(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	MaximumDemand := BaseCalculation(data[0:6])
	MaximumDemand = MaximumDemand[0:2] + "." + MaximumDemand[2:]
	timer := TimeStitchingF(BaseCalculation(data[6:16]))
	result.Set(0, MaximumDemand)
	result.Set(1, timer)
	return result.Data()
}

// TotalPowerFactor 总功率因数
func TotalPowerFactor(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	factor := BaseCalculation(data)
	result.Set(0, factor[0:1]+"."+factor[1:])
	return result.Data()
}

// HarmonicContentData 谐波含量数据块
func HarmonicContentData(data string) map[int]string {
	result := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	var slice []string
	num := 0
	length := 21
	for num < length {
		s := BaseCalculation(data[num*4:num*4+4])
		s = s[0:2]+"."+s[2:]
		slice = append(slice,s)
		num++
	}
	for k,v := range slice{
		result.Set(k,v)
	}
	return result.Data()
}

// 日数据处理 13761 报文处理
func dailFrozenDataProcessing(data string) map[int]string {
	result := mutex_map.MutexMapInt{
		Mp: make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	count := ByteInversion(data[0:10])
	a := ByteInversion(data[10:20])
	b := ByteInversion(data[20:30])
	c := ByteInversion(data[30:40])
	d := ByteInversion(data[40:50])
	result.Set(0,count[0:6]+"."+count[6:])
	result.Set(1,a[0:6]+"."+a[6:])
	result.Set(2,b[0:6]+"."+b[6:])
	result.Set(3,c[0:6]+"."+c[6:])
	result.Set(4,d[0:6]+"."+d[6:])
	return result.Data()
}

// IpApn IP 端口处理
func IpApn(data string) map[int]string {
	result := mutex_map.MutexMapInt{
		Mp: make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	port1,_ := strconv.ParseUint(ByteInversion(data[8:12]),16,32)
	port2,_ := strconv.ParseUint(ByteInversion(data[20:24]),16,32)
	apn := func(data string) string {
		var result []string
		var apn [] string
		length := len(data)
		num := 0
		for num < length {
			s := data[num : num+2]
			result = append(result,s)
			num = num + 2
		}
		for _,char := range result{
			hex,_ := strconv.ParseUint(char,16,32)
			apn = append(apn,cast.ToString(hex))
		}
		return strings.Join(apn,"")
	}(data)
	strPort1 := cast.ToString(port1)
	strPort2 := cast.ToString(port2)
	ip1 := calculationIp(data[0:8])
	ip2 := calculationIp(data[12:20])
	result.Set(0,ip1)
	result.Set(1,ip2)
	result.Set(2,strPort1)
	result.Set(3,strPort2)
	result.Set(4,apn)
	return result.Data()
}

// 计算 IP
func calculationIp(data string) string {
	var result []string
	length := len(data)
	num := 0
	for num < length {
		s := data[num : num+2]
		result = append(result, s)
		num = num + 2
	}
	res := make(map[int]string)
	for k, v := range result {
		hex, _ := strconv.ParseInt(v, 16, 10)
		str := cast.ToString(hex)
		res[k] = str
	}
	return res[0] + "." + res[1] + "." + res[2] + "." + res[3]
}

// CheckGetFile13761 召测电表档案解析
func CheckGetFile13761(data string, number string) map[int]map[string]string {
	var result []string
	num,_ := strconv.Atoi(number)
	var iu int
	for i := 0; i < num; i++ {
		s := data[iu:iu+54]
		iu = iu+54
		result = append(result,s)
	}
	res := mutex_map.MutexMapIntMap{
		Mp: make(map[int]map[string]string),
		RWMutex: new(sync.RWMutex),
	}
	for k,v := range result{
		r := mutex_map.MutexMapStr{
			Mp: make(map[string]string),
			RWMutex: new(sync.RWMutex),
		}
		strToInt,_ := strconv.Atoi(transformation.StrHex2Dec(ByteInversion(v[0:4])))
		r.Set("measuringPoint",strconv.Itoa(strToInt))
		two := v[8:10]
		var rate string
		var port string
		switch two {
		case "7F":
			rate = "2400bps"
			port = "载波"
		case "62":
			rate = "2400bps"
			port = "485-1"
		case "63":
			rate = "2400bps"
			port = "485-2"
		case "C2":
			rate = "9600bps"
			port = "485-1"
		default:
			rate = "速率异常"
			port = "端口异常"
		}
		r.Set("rate",rate)
		r.Set("port",port)
		var statute string
		if v[10:12] == "1E" {
			statute = "645-07"
		}else if v[10:12] == "20" {
			statute = "188"
		}else {
			statute = "未知规约"
		}
		r.Set("statute",statute)
		r.Set("address",ByteInversion(v[12:24]))
		val := v[36:38]
		flat,_ := strconv.ParseUint(val,10,32)
		r.Set("flat",strconv.FormatUint(flat,10))
		res.Set(k,r.Mp)
	}
	return res.Data()
}
package electric

import (
	"github.com/spf13/cast"
	"strconv"
	"sync"
	"tour/pkg/message_tool"
	"tour/pkg/mutex_map"
)

// Analysis 13761、13761 透明转发 645 报文解析
func Analysis(str string) map[string]interface{} {
	// 报文过滤
	str = message_tool.Filter(str)
	result := &mutex_map.MutexMap{
		Mp:      make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	// 区分协议
	if str[10:12] == "68" { // 13761
		if str[24:26] == "10" { // 透明转发
			result.Mp = sectionAnalysisPenetrate(str)
			section := result.Get("section").(map[string]string)
			result.Set("command645", Analysis(section["command"]))
			delete(section,"command")
		} else { // 常规 13761
			result.Mp = sectionAnalysis3761(str)
		}
	} else if str[14:16] == "68" { // 645
		result.Mp = sectionAnalysis(str)
		if result.Get("length") == "00" {
			result.Set("status", 0)
			return result.Data()
		}
		type645 := message_tool.Type645(result.Get("flag").(string), str)
		result.Set("messageData", type645["messageData"])
		result.Set("status", 0)
		return result.Data()
	}
	return result.Data()
}

// 13761 报文区段解析
func sectionAnalysis3761(str string) map[string]interface{} {
	result := &mutex_map.MutexMap{
		Mp: make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	section := message_tool.Type13761(public3761(str),str)
	result.Set("section",section["section"])
	result.Set("messageType",section["messageType"])
	result.Set("messageData",section["messageData"])
	result.Set("originalMessage",str)
	return result.Data()
}

// 645 报文区段解析
func sectionAnalysis(str string) map[string]interface{} {
	result := &mutex_map.MutexMap{
		Mp:      make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	// 起始帧 1
	result.Set("startCharacter", str[0:2])
	// 表号
	result.Set("address", message_tool.ByteInversion(str[2:14]))
	// 起始符 2
	result.Set("startCharacter2", str[14:16])
	// 控制码
	result.Set("controlCode", str[16:18])
	// 数据长度
	result.Set("length", str[18:20])
	// 校验位
	result.Set("check", str[len(str)-4:len(str)-2])
	// 结束符
	result.Set("endCharacter", str[len(str)-2:])
	// 原始报文
	result.Set("originalMessage",str)
	if result.Get("controlCode") == "9C" && result.Get("length") == "00" {
		result.Set("messageType", "阀控成功")
	} else {
		// 数据标识
		result.Set("flag", str[20:28])
		message := message_tool.Type645(result.Get("flag").(string), str)
		result.Set("messageType", message["messageType"])
	}
	return result.Data()
}

// 13761 公共区段
func public3761(data string) map[string]string {
	public3761 := &mutex_map.MutexMapStr{
		Mp: make(map[string]string),
		RWMutex: new(sync.RWMutex),
	}
	// 起始符1
	public3761.Set("startCharacter",data[0:2])
	// 长度位
	public3761.Set("lengthBit",data[2:10])
	// 起始符2
	public3761.Set("startCharacter2",data[10:12])
	// 控制码
	public3761.Set("controlCode",data[12:14])
	// 终端地址
	public3761.Set("sn",message_tool.SnInversion(data[14:24]))
	// afn
	public3761.Set("afn",data[24:26])
	// 帧序列
	public3761.Set("frameSequence",data[16:28])
	// 测量点号
	measuringPoint := message_tool.ReCalculatedAddress(data[28:32])
	public3761.Set("measuringPoint",cast.ToString(measuringPoint))
	// 功能码
	public3761.Set("functionCode",data[32:36])
	return public3761.Data()
}

// 透明转发报文区段解析
func sectionAnalysisPenetrate(data string) map[string]interface{} {
	result := mutex_map.MutexMap{
		Mp: make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	section := mutex_map.MutexMapStr{
		Mp: make(map[string]string),
		RWMutex: new(sync.RWMutex),
	}
	section.Mp = public3761(data)
	section.Set("port",data[36:38])
	length645 := data[38:42]
	section.Set("length645",length645)
	uin,_ := strconv.ParseUint(length645[0:2],16,10)
	section.Set("command",data[42:42+uin*2])
	section.Set("check",data[len(data)-4:len(data)-2])
	section.Set("endCharacter",data[len(data)-2:])

	result.Set("section",section.Mp)
	result.Set("originalMessage",data)
	return result.Data()
}

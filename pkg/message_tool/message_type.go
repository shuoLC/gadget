package message_tool

import (
	"strconv"
	"sync"
	"tour/pkg/mutex_map"
)

// Type645 645 报文类型判断
func Type645(flag string, str string) map[string]interface{} {
	result := &mutex_map.MutexMap{
		Mp:      make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	messageData := &mutex_map.MutexMapInt{
		Mp:      make(map[int]string),
		RWMutex: new(sync.RWMutex),
	}
	switch flag {
	case "33323433":
		result.Set("messageType", "当前正向有功电能数据块")
		messageData.Mp = ReadCurrentData(str[28:68], 8, 6)
	case "33323533":
		result.Set("messageType", "当前反向有功电能数据块")
		messageData.Mp = ReadCurrentData(str[28:68], 8, 6)
	case "33323535":
		result.Set("messageType", "电流数据块")
		messageData.Mp = ReadCurrentData(str[28:46], 6, 3)
	case "33323435":
		result.Set("messageType", "电压数据块")
		messageData.Mp = ReadCurrentData(str[28:40], 4, 3)
	case "34343938":
		result.Set("messageType", "上一次日冻结正向有功电能数据块")
		messageData.Mp = ReadCurrentData(str[28:68], 8, 6)
	case "34353938":
		result.Set("messageType", "上一次日冻结反向有功电能数据块")
		messageData.Mp = ReadCurrentData(str[28:68], 8, 6)
	case "3433B335":
		result.Set("messageType", "零线电流")
		messageData.Mp = ReadCurrentData(str[28:34], 6, 3)
	case "33406336":
		result.Set("messageType", "电表开盖次数")
		messageData.Mp = ReadCurrentData(str[28:34], 6, 6)
	case "34406336":
		result.Set("messageType", "上次开盖记录")
		messageData.Mp = LastOpeningRecord(str[28:52])
	case "33334436":
		result.Set("messageType", "掉电总次数")
		messageData.Mp = ReadCurrentData(str[28:34], 6, 6)
	case "34334436":
		result.Set("messageType", "上一次掉电事件时间")
		messageData.Mp = PowerFailure(str[28:52])
	case "3433344C":
		result.Set("messageType", "过流总次数")
		messageData.Mp = TotalTimesOfOverCurrent(str[28:34])
	case "3434344C":
		result.Set("messageType", "上一次过流发生时刻")
		messageData.Mp = OverCurrentTime(str[28:40])
	case "3454344C":
		result.Set("messageType", "上一次过流结束时刻")
		messageData.Mp = OverCurrentTime(str[28:40])
	case "34323738":
		result.Set("messageType", "上一次整点冻结数据块")
		messageData.Mp = LastHourData(str[28:60])
	case "33323635":
		result.Set("messageType", "瞬时有功功率数据块")
		messageData.Mp = PowerDataBlock(str[28:52], 6, 2)
	case "33323735":
		result.Set("messageType", "瞬时无功功率数据块")
		messageData.Mp = PowerDataBlock(str[28:52], 6, 2)
	case "34323350":
		result.Set("messageType", "上一次跳闸事件")
		//messageData.Mp = TrippingAndClosingEvents(str[28:96])
	case "34323351":
		result.Set("messType", "上一次合闸事件")
		//messageData.Mp = TrippingAndClosingEvents(str[28:96])
	case "34333434":
		result.Set("messageType", "上一日正向有功最大需量及发生事件")
		//messageData.Mp = MaximumDemandAndTime(str[28:44])
	case "33333935":
		result.Set("messageType", "总功率因数")
		//messageData.Mp = TotalPowerFactor(str[28:32])
	case "32343D35":
		result.Set("messageType", "A相电压谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	case "32353D35":
		result.Set("messageType", "B相电压谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	case "32363D35":
		result.Set("messageType", "C相电压谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	case "32343E35":
		result.Set("messageType", "A相电流谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	case "32353E35":
		result.Set("messageType", "B相电流谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	case "36383337":
		result.Set("messageType", "C相电流谐波含量数据块")
		//messageData.Mp = HarmonicContentData(str[28:112])
	}
	result.Set("messageData", messageData.Mp)
	return result.Data()
}

// Type13761 13761 类型报文解析
func Type13761(oldMap map[string]string,data string) map[string]interface{} {
	result := mutex_map.MutexMap{
		Mp: make(map[string]interface{}),
		RWMutex: new(sync.RWMutex),
	}
	section := mutex_map.MutexMapStr{
		Mp: make(map[string]string),
		RWMutex: new(sync.RWMutex),
	}
	section.Mp = oldMap
	switch section.Get("functionCode") {
	case "0110":
		result.Set("messageType","正向有功")
		section.Set("time",TimeStitchingF(data[36:46]))
		section.Set("rate",data[46:48])
		section.Set("type","0")
		result.Set("messageData",dailFrozenDataProcessing(data[48:98]))
	case "0410":
		result.Set("messageType","反向有功")
		section.Set("time",TimeStitchingF(data[36:46]))
		section.Set("rate",data[46:48])
		section.Set("type","0")
		result.Set("messageData",dailFrozenDataProcessing(data[48:98]))
	case "0114":
		result.Set("messageType","上一日冻结正向有功电能")
		section.Set("timeScale",TimeStitchingR(data[36:42]))
		section.Set("time",TimeStitchingF(data[42:52]))
		section.Set("type","0")
		result.Set("messageData",dailFrozenDataProcessing(data[54:104]))
	case "0414":
		result.Set("messageType","上一日冻结反向有功电能")
		section.Set("timeScale",TimeStitchingR(data[36:42]))
		section.Set("time",TimeStitchingF(data[42:52]))
		section.Set("type","0")
		result.Set("messageData",dailFrozenDataProcessing(data[54:104]))
	case "0400":
		result.Set("messageType","主站IP")
		section.Set("type","0")
		result.Set("messageData",IpApn(data[36:92]))
	case "0201":
		result.Set("messageType","电表档案")
		number,_ := strconv.ParseUint(ByteInversion(data[36:40]),16,32)
		section.Set("number",strconv.FormatUint(number,10))
		result.Set("messageData",CheckGetFile13761(data[40:],section.Get("number")))
		section.Set("type","0")
	case "0100":
		result.Set("messageType","命令操作")
		section.Set("type","1")
		sMap := make(map[int]string)
		lock.Lock()
		if data[24:26] == "00" {
			sMap[0] = "成功"
		}else {
			sMap[0] = "失败"
		}
		lock.Unlock()
		result.Set("messageData",sMap)
	}
	result.Set("section",section.Data())
	return result.Data()
}
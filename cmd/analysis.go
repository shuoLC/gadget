package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/agreement/electric"
	"tour/internal/word"
)

var message string
var typer int8

const (
	TyperElectric = iota + 1
	TyperWater
)

var docs = strings.Join([]string{
	"该自命令支持解析协议如下：",
	"1：376.1协议，645协议，13761透传协议",
	"2：CJT188协议",
},"\n")

var anaCmd = &cobra.Command{
	Use: "analysis",
	Short: "报文解析",
	Long: docs,
	Run: func(cmd *cobra.Command, args []string) {
		var content map[string]interface{}
		switch typer {
		case TyperElectric:
			message = word.RemoveAllSpaces(message)
			content = electric.Analysis(message)
		case TyperWater:
			log.Fatalf("该协议暂未完成")
		default:
			log.Fatalf("暂不支持该协议，请执行 help analysisi 查看帮助文档")
		}
		byt,_ := json.Marshal(content)
		log.Printf("输出结果：%s",string(byt))
	},
}

func init()  {
	anaCmd.Flags().StringVarP(&message,"message","m","","请输入原始报文")
	anaCmd.Flags().Int8VarP(&typer,"type","t",0,"请输入报文解析类型")
}

package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/aes"
)

var msg string

var msgCmd  = &cobra.Command{
	Use: "aes",
	Short: "报文加密",
	Long: "报文加密",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		msg = strings.ReplaceAll(msg," ","")
		content,_ = aes.Encrypt(msg)
		log.Printf("输出结果：%s", content)
	},
}

func init()  {
	msgCmd.Flags().StringVarP(&msg,"msg","m","","请输入报文")
}

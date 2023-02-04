package cmd

import (
	"zt/service"

	"github.com/spf13/cobra"
)

var WebInitProject bool

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web项目工具",
	Args:  cobra.ExactArgs(1),
	Long:  `一个个人封装的gin项目，详情请了解：https://github.com/ZTaboo/go-gin.`,
	Run:   service.Init,
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webCmd.Flags().BoolVarP(&WebInitProject, "init", "i", false, "<project name>初始化web项目")
}

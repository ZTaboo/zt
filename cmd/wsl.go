package cmd

import (
	s "zt/service"

	"github.com/spf13/cobra"
)

var WslOpenStatus bool

// wslCmd represents the wsl command
var wslCmd = &cobra.Command{
	Use:   "wsl",
	Short: "操作wsl",
	Long:  `open:进入wsl并进入当前目录`,
	Run:   s.WslInit,
}

func init() {
	rootCmd.AddCommand(wslCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wslCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wslCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	wslCmd.Flags().BoolVarP(&WslOpenStatus, "open", "o", false, "打开wsl并进入当前目录")
}

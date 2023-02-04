package service

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func Init(cmd *cobra.Command, args []string) {
	f := cmd.Flag("init")
	if f.Value.String() == "true" {
		fmt.Println("开始创建，请稍等")
		cmd := exec.Command("git", "clone", "https://github.com/ZTaboo/go-gin.git", args[0])
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("err:", err)
		}
	} else {
		fmt.Println("请输入flag;使用-h查看帮助文档")
	}
}

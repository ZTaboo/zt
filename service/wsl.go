package service

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"zt/public"
)

func WslInit(cmd *cobra.Command, args []string) {
	if public.WslOpenStatus {
		cmd := exec.Command("powershell", "wt -p \"Ubuntu-20.04\" -d .")
		bytes, err := cmd.Output()
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(string(bytes))
		}
	} else {
		fmt.Println("请输入flags;使用-h查看帮助文档")
	}
}

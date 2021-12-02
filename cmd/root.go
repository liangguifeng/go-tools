package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// debug模式
var debug bool

// 命令行超时时间
var commandTimeout int

// 命令行的执行目录
var commandDir string

// stark-tools配置所在文件路径
var toolsCnfPath string

// rootCmd 声明命令
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "go-tools is help developers tools",
	Long:  `go-tools is a CLI library for developer.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

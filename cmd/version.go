package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 声明当前版本
const GO_TOOLS_VERSION = "v1.1.0"

// versionCmd 声明命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show current go-tools version",
	Long:  "show current go-tools version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-tools version %s\n", GO_TOOLS_VERSION)
		fmt.Println(`
                           __                .__          
   ____   ____           _/  |_  ____   ____ |  |   ______
  / ___\ /  _ \   ______ \   __\/  _ \ /  _ \|  |  /  ___/
 / /_/  >  <_> ) /_____/  |  | (  <_> |  <_> )  |__\___ \ 
 \___  / \____/           |__|  \____/ \____/|____/____  >
/_____/                                                \/ `)
	},
}

// init 默认执行
func init() {
	rootCmd.AddCommand(versionCmd)
}

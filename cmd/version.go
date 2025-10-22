package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version  = "0.9.6"        // 版本递增
	codename = "MR"           // 修改 codename
	intro    = "A MR backend that supports many panels" // 修改介绍文字
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print current version of MR",
		Run: func(cmd *cobra.Command, args []string) {
			showVersion()
		},
	})
}

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}

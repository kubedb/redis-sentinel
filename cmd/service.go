package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createServiceCMD = &cobra.Command{
	Use: "redis-service",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Created Service  for redis server ")
		instance.CreateHeadlessService()
	},
}


var createSentinelSvcCMD = &cobra.Command{
	Use: "senti-service",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create senti service cmd called")
		instance.CreateSentinelHeadlessService()
	},
}

func init() {
	createCMD.AddCommand(createServiceCMD)
	createCMD.AddCommand(createSentinelSvcCMD)
	fmt.Println(".....")
}

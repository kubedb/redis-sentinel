package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createServiceCMD = &cobra.Command{
	Use: "service",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Service cmd called")
		instance.CreateHeadlessService()
	},
}


var createSentinelSvcCMD = &cobra.Command{
	Use: "svc1",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Svc1 cmd called")
		instance.CreateSentinelHeadlessService()
	},
}

func init() {
	createCMD.AddCommand(createServiceCMD)
	createCMD.AddCommand(createSentinelSvcCMD)
	fmt.Println(".....")
}

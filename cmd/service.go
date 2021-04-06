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

var createSvc1CMD = &cobra.Command{
	Use: "svc1",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Svc1 cmd called")
		instance.CreateSvc1()
	},
}

func init() {
	createCMD.AddCommand(createServiceCMD)
	createCMD.AddCommand(createSvc1CMD)
	fmt.Println(".....")
}

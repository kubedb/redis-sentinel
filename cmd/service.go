package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createServiceCMD = &cobra.Command{
	Use: "new-Service",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Service cmd called")
		instance.CreateHeadlessService()
	},
}

func init() {
	createCMD.AddCommand(createServiceCMD)
}

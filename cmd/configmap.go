package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createConfigCMD = &cobra.Command{
	Use: "new-Configmap",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Configmap cmd called")
		instance.CreateConfigmap()
	},
}

func init() {
	createCMD.AddCommand(createConfigCMD)
}

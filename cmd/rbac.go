package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createRBAC = &cobra.Command{
	Use: "role",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Created Service  for redis server ")
		instance.CreateRole("default", "predis-role")
	},
}

var createBinding = &cobra.Command{
	Use: "binding",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Created Service  for redis server ")
		// 3rd default is service account name.
		instance.CreateRoleBinding("default", "predis-role", "predis-account")

	},
}


var createSA = &cobra.Command{
	Use: "sa",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Created Service  for redis server ")
		// 3rd default is service account name.
		instance.CreateServiceAccount("default", "predis-account")

	},
}


func init() {
	createCMD.AddCommand(createRBAC)
	createCMD.AddCommand(createBinding)
	createCMD.AddCommand(createSA)

	fmt.Println(".....")
}

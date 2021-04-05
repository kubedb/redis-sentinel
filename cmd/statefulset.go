package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var (
	defaultImage   string
	defaultReplica int32
)

var createStateCMD = &cobra.Command{
	Use: "new-statefulset",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create statefulset cmd called")
		instance.CreateStatefulset(defaultImage, defaultReplica)
	},
}

var listStateCMD = &cobra.Command{
	Use: "list-statefulset",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List statefulset cmd called")
		instance.ListStatefulSet()
	},
}

var deleteStateCMD = &cobra.Command{
	Use: "delete-statefulset",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete statefulset cmd called")
		instance.DeleteStatefulSet()
	},
}

func init() {
	createCMD.AddCommand(createStateCMD)
	createCMD.AddCommand(listStateCMD)
	createCMD.AddCommand(deleteStateCMD)

	createStateCMD.PersistentFlags().StringVarP(&defaultImage, "image", "i", "pranganmajumder/go-basic-restapi:1.0.0", "It sets the custom image you want")
	createStateCMD.PersistentFlags().Int32VarP(&defaultReplica, "replica", "r", 3, "It sets the number of replica user want")
}
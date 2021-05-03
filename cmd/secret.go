package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)


var createSecretCMD = &cobra.Command{
	Use: "secret",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling CreateSecret()  for redis server ")
		instance.CreateSecret()
	},
}


func init() {
	createCMD.AddCommand(createSecretCMD)
	fmt.Println(".....created Secret 1........... ")
}
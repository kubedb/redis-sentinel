package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//CreateCMD is our main command
var createCMD = &cobra.Command{
	Use:   "create",
	Short: "It will create a new resource you want",
	Long: `It will create a new resources like: ./app create deploy . Here deploy is our another child/
			sub-command of create command . You need to create the deploy sub-command manually and add it to createCMD`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create command called ")
	},
}

func init() {
	rootCmd.AddCommand(createCMD)

}

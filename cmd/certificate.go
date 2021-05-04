


package cmd

import (
	"fmt"
	"github.com/pranganmajumder/predis/instance"
	"github.com/spf13/cobra"
)

var createCertificateCMD = &cobra.Command{
	Use: "cert",

	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Created Service  for redis server ")
		_, err := instance.CreateServerCert()
		if err != nil {
			panic(err)
		}
	},
}




func init() {
	createCMD.AddCommand(createCertificateCMD)
	fmt.Println(".....")
}

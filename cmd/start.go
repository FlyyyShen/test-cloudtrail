package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"test-cloudtrail/impl"
)




var principalIdCmd = &cobra.Command{
	Use: "principalId",
	Short: "cloudtrail principalId",
	Long:  `cloudtrail principalId`,
	RunE: func(cmd *cobra.Command, args []string) error {
		impl.CompressGz()
		pwd, _ := os.Getwd()
		pathname := pwd+"/cloudtrail/logs/"
		impl.GetprincipalId(pathname)

		return nil
	},
}

var resourceCmd = &cobra.Command{
	Use: "resource",
	Short: "cloudtrail resource",
	Long:  `cloudtrail created resource research`,
	RunE: func(cmd *cobra.Command, args []string) error {
		impl.CompressGz()
		pwd, _ := os.Getwd()
		pathname := pwd+"/cloudtrail/logs/"
		impl.GetCreateResource(pathname)

		return nil
	},
}


func init()  {
	RootCmd.AddCommand(principalIdCmd)
	RootCmd.AddCommand(resourceCmd)
}
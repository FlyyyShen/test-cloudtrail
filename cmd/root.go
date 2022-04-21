package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var vers  bool

var (
	configType string
	confFile string
)

var RootCmd = &cobra.Command{
	Use: "demo test cloudtrail",
	Short: "cloudtrail",
	Long: "cloudtrail...",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("0.0.1")
			return nil
		}
		return errors.New("no flags find")

	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the test-cloudtrail version")
}
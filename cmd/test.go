package cmd

import (
	"github.com/lcmps/ExodiaLibrary/app"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use: "test",
	RunE: func(cmd *cobra.Command, args []string) error {

		app.DownloadImages()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

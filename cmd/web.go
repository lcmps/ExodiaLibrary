package cmd

import (
	"github.com/lcmps/ExodiaLibrary/web"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use: "web",
	RunE: func(cmd *cobra.Command, args []string) error {

		w, err := web.New()
		if err != nil {
			return err
		}

		w.Host()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}

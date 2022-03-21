package cmd

import (
	"fmt"

	"github.com/lcmps/ExodiaLibrary/db"
	"github.com/lcmps/ExodiaLibrary/web"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	RunE: func(cmd *cobra.Command, args []string) error {

		conn, err := db.InitConnection()
		if err != nil {
			fmt.Println(err.Error())
		}
		conn.ImportCards()

		w, err := web.New()
		if err != nil {
			return err
		}

		w.Host()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

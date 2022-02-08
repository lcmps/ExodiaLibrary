package cmd

import (
	"fmt"

	"github.com/lcmps/ExodiaLibrary/db"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {

		conn, err := db.InitConnection()
		if err != nil {
			fmt.Println(err.Error())
		}
		conn.CreateTables()
		conn.ImportCards()
		// fmt.Println("Migrate and populate Databases;")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Do database(s) migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run database(s) migration")
		// Do db(s) migration here
	},
}

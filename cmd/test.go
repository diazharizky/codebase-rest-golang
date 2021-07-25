package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Do functionality test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run functionality test")
		// Call the test function here
	},
}

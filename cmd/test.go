package cmd

import (
	"fmt"

	"github.com/fomiller/scribe/utils"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		test := utils.GetFolderName(Name)
		fmt.Println(test)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	// Here you will define your flags and configuration settings.
}

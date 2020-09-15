package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list out files in your drive.",
	Long:  `List will return you a list of the files located in your google drive without a search criteria.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileNameList := drive.ListFileNames()
		file := ""
		prompt := &survey.Select{
			Message:  "Select a file:",
			Options:  fileNameList,
			PageSize: 10,
		}

		survey.AskOne(prompt, &file, survey.WithKeepFilter(true))
		fmt.Printf("You chose: %v", file)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	// Here you will define your flags and configuration settings.
}

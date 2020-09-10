package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// multideleteCmd represents the multidelete command
var multideleteCmd = &cobra.Command{
	Use:   "multidelete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileNameList := drive.ListFileNames()
		filesToDelete := []string{}
		prompt := &survey.MultiSelect{
			Message:  "What files do you want to delete:",
			Options:  fileNameList,
			PageSize: 15,
		}
		survey.AskOne(prompt, &filesToDelete)

		// delete files that were selected to be deleted
		for _, v := range filesToDelete {
			// get docId for file to be deleted
			docId, err := drive.GetFileId(v)
			if err != nil {
				log.Fatal(err)
			}
			// delete file
			err = drive.DeleteFile(docId)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v Deleted\n", v)
		}
	},
}

func init() {
	rootCmd.AddCommand(multideleteCmd)
	// Here you will define your flags and configuration settings.
}

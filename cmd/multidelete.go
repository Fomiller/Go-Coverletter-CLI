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
	Short: "Delete multiple files by selecting from a list.",
	Run: func(cmd *cobra.Command, args []string) {
		fileNameList := drive.ListFileNames()
		filesToDelete := []string{}
		prompt := &survey.MultiSelect{
			Message:  "What files do you want to delete:",
			Options:  fileNameList,
			PageSize: 10,
		}

		survey.AskOne(prompt, &filesToDelete, survey.WithKeepFilter(true))

		// If the user selects ALL Files to be deleted double check if they would like to delete ALL files
		if len(fileNameList) == len(filesToDelete) {
			// deletePrompt := &survey.Confirm{
			// 	Message: fmt.Sprintf("You have selected ALL of the files in your drive, %v files. Are You sure you want to delete all of these files? You CAN NOT undo these changes", len(filesToDelete)),
			// }
			// deleteAll := false
			// survey.AskOne(deletePrompt, deleteAll)
			// if deleteAll == true {
			// 	fmt.Println("Deleting all files")
			// }
			fmt.Println(`
	Sorry! As a saftey precaution Scribe is not allowed to delete all of your files.
	
	If you wish to delete all of your files please do so using the Google Drive/Docs application.
	
	`)
		} else {
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
		}
	},
}

func init() {
	rootCmd.AddCommand(multideleteCmd)
	// Here you will define your flags and configuration settings.
}

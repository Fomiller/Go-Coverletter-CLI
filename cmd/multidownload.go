package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/config"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// multidownloadCmd represents the multidownload command
var multidownloadCmd = &cobra.Command{
	Use:   "multidownload",
	Short: "Download multiple files from selecting from a list.",
	Run: func(cmd *cobra.Command, args []string) {
		fileNameList := drive.ListFileNames()
		filesToDownload := []string{}
		prompt := &survey.MultiSelect{
			Message:  "What files do you want to download:",
			Options:  fileNameList,
			PageSize: 10,
		}
		survey.AskOne(prompt, &filesToDownload, survey.WithKeepFilter(true))

		// if path flag is not set prompt if you would like to set the output path default to "NO"
		if Path == "" {
			// init setPath variable
			setPath := false
			prompt := &survey.Confirm{
				Message: "Do you want to set a custom output path for this file?",
				Default: false,
			}
			survey.AskOne(prompt, &setPath)
			// if the user confirms to set an output path enter the path here and set to config.Scribe.Download.Path
			if setPath == true {
				prompt := &survey.Input{
					Message: "What path would you like to use to download your file to?",
				}
				survey.AskOne(prompt, &config.Scribe.Download.Path)
			}
		}

		for _, file := range filesToDownload {
			// get file id from name variable
			docId, err := drive.GetFileId(file)
			if err != nil {
				log.Fatal(err)
			}
			// download file
			drive.DownloadFile(docId, file)
			fmt.Printf("%v Downloaded\n", file)
		}

	},
}

func init() {
	rootCmd.AddCommand(multidownloadCmd)
	// Here you will define your flags and configuration settings.
}

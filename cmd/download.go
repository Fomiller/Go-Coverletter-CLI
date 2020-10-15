package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/config"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file.",
	Run: func(cmd *cobra.Command, args []string) {
		// if path flag is set use as download path
		if Path != "" {
			config.Scribe.Download.Path = Path
		}
		// if name flag is not set ask for the file name
		if Name == "" {
			prompt := &survey.Input{
				Message: "What is the name of the file you want to download",
			}
			survey.AskOne(prompt, &Name)
		}

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

		// get file id from name variable
		docId, err := drive.GetFileId(Name)
		if err != nil {
			log.Fatal(err)
		}
		// download file
		drive.DownloadFile(docId, Name)
		fmt.Printf("%v Downloaded", Name)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

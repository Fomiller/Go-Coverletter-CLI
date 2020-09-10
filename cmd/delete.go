/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// var deleteFile string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete File from Google Drive",
	Run: func(cmd *cobra.Command, args []string) {
		if NewFileName == "" {
			prompt := &survey.Input{
				Message: "What is the name of the file you want to delete?",
			}
			survey.AskOne(prompt, &NewFileName)
		}
		// get docId for file to be deleted
		docId, err := drive.GetFileId(NewFileName)
		if err != nil {
			log.Fatal(err)
		}
		// delete file
		err = drive.DeleteFile(docId)
		// handle error if necessary
		if err != nil {
			log.Fatal(err)
		} else {
			// if err is nil
			fmt.Printf("%v Deleted", NewFileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// deleteCmd.Flags().StringVarP(&deleteFile, "delete", "p", "", "Insert your file name here.")
}

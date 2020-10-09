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

	"github.com/fomiller/scribe/utils"

	"github.com/fomiller/scribe/docs"
	"github.com/fomiller/scribe/drive"
	"github.com/fomiller/scribe/sheets"
	"github.com/spf13/cobra"
)

// createfromsheetCmd represents the createfromsheet command
var createfromsheetCmd = &cobra.Command{
	Use:   "createfromsheet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileId, err := drive.GetFileId(SheetName)
		if err != nil {
			log.Fatal(err)
		}
		fieldNames := sheets.GetSpreadsheetColumnNames(fileId)
		rowData := sheets.GetRowData(fileId)
		spreadsheetData := sheets.FmtSpreadsheetData(fieldNames, rowData)

		// initialize FieldMap
		FieldMap := make(map[string]string)
		// create map[string]string to hold key value pairs from spreadsheet data such as "Name":"Bob"
		for _, v := range spreadsheetData {
			for _, vv := range v {
				// convert from interface{} to string
				fkey := fmt.Sprint(vv.FieldName)
				// convert from interface{} to string
				fvalue := fmt.Sprint(vv.FieldValue)
				FieldMap[fkey] = fvalue
			}

			// create file name by adding the TemplateName and the Unique File ID value together
			Name = fmt.Sprintf("%v - %v", TemplateName, FieldMap["Unique File ID"])
			// strings to be removed from the file name
			removeStrings := []string{"TEMPLATE", "template", "Template"}
			// removing strings from file name
			Name := utils.FmtFileName(Name, removeStrings...)

			// create file for each row in spreadsheetData
			docs.CreateFile(Name, TemplateName, FieldMap, DlFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(createfromsheetCmd)
}

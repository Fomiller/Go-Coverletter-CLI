package cmd

import (
	"fmt"

	"github.com/fomiller/scribe/sheets"
	"github.com/spf13/cobra"
)

// createfromfileCmd represents the createfromfile command
var createfromfileCmd = &cobra.Command{
	Use:   "createfromfile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fieldNames := sheets.GetSpreadsheetColumnNames()
		rowData := sheets.GetRowData()
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
			// create file for each row in spreadsheetData
			CreateFile(fmt.Sprintf("test-%v", FieldMap["Student Name"]), "FROM FILE TEMPLATE", FieldMap, false)
			// print the field map
			// fmt.Println(FieldMap)
		}
	},
}

func init() {
	rootCmd.AddCommand(createfromfileCmd)
}

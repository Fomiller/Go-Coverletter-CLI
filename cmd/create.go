package cmd

import (
	"fmt"

	"github.com/fomiller/scribe/docs"
	"github.com/fomiller/scribe/drive"

	"github.com/spf13/cobra"
)

var FieldMap map[string]string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new file using a template",
	Long: `Create command will create a new file based off of a template and field arguments provided by the field flag.

	Use the field flag -f to specify fields inside your template that need to be replaced with data.
	for example:
	fields takes in a map[string]string
	Single field example:
		--fields 'name=Myname'
	Multiple fields example:
	comma seperated single string
		-f 'name=Myname, date=12/10/1993'
	comma seperated single string with substrings
		-f '"name=Myname", "date=12/10/1993"' OR -f '"name"="Myname", "date"="12/10/1993"' OR -f '"name=Myname" -f '"date=12/10/1993"
	
	*All keys are automatically capitalized to match fields in Google doc template ex: '{{NAME}}'`,

	Run: func(cmd *cobra.Command, args []string) {
		// print out the name of the file being downloaded
		fmt.Printf("Creating: %v\n", NewFileName)
		// print out the name of the template being used
		fmt.Printf("Using template: %v\n", TemplateName)
		// Get Template Id from the template name
		templateId := drive.GetFileId(TemplateName)
		// create and return docId for new file using NewFileName and the templateID from TemplateName,
		docId := drive.NewTemplate(NewFileName, templateId)
		// create replace struct from field flags
		// **fields to be changed inside the document/template
		replaceStruct := docs.CreateRequestStruct(FieldMap)
		// update the newfile using the docId with the replace struct
		docs.NewUpdateTemplateFile(docId, replaceStruct)

		fmt.Println("New File Created")

		if DlFile == true {
			drive.DownloadFile(docId, NewFileName)
			fmt.Println("New File Downloaded")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringToStringVarP(&FieldMap, "field", "f", nil, "use this to fill out custom fields")
}

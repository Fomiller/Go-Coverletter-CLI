package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/docs"
	"github.com/fomiller/scribe/drive"

	"github.com/spf13/cobra"
)

var createQuestions = []*survey.Question{
	{
		Name:     "fileName",
		Prompt:   &survey.Input{Message: "What is your new file name"},
		Validate: survey.Required,
	},
	{
		Name:     "templateName",
		Prompt:   &survey.Input{Message: "What is the name of the template you want to use"},
		Validate: survey.Required,
	},
	// {
	// 	Name: "download",
	// 	Prompt: &survey.Confirm{
	// 		Message: "Do you want to download this file?",
	// 		Default: false,
	// 	},
	// 	Validate: survey.Required,
	// },
}

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

		// execute the create commands when all variables are predefined with flags ie: scribe create -n "newFile" -t "fromTemplate" -f "date=1/01/2020"
		if NewFileName != "" && TemplateName != "" && FieldMap != nil {
			CreateFile(NewFileName, TemplateName, FieldMap, DlFile)
		}

		// if no arguments specified
		if NewFileName == "" || TemplateName == "" {
			// check if only template argument is missing return an error
			if NewFileName == "" && TemplateName != "" {
				fmt.Println("Template argument missing")
				return
			}
			// check if only NewfileName argument is missing return an error
			if TemplateName == "" && NewFileName != "" {
				fmt.Println("Filename argument missing ")
				return
			}

			// define answer struct
			answers := struct {
				FileName     string // survey will match the question and field names
				TemplateName string `survey:"templateName"` // or you can tag fields to match a specific name
				Download     bool   // if the types don't match, survey will convert it
			}{}

			// perform the questions
			err := survey.Ask(createQuestions[:2], &answers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			// set NewFileName, TemplateName, DlFile to recorded answers from survey
			NewFileName = answers.FileName
			TemplateName = answers.TemplateName
			parseCmd.Run(cmd, args)
			// // Unmarshal answer.Fields into type FieldMap map[string]string
			// json.Unmarshal([]byte(answers.Fields), &FieldMap)
			prompt := &survey.Confirm{
				Message: "Do you want to download this file?",
				Default: false,
			}
			survey.AskOne(prompt, &answers.Download)

			// err = survey.Ask(createQuestions[2:], &answers.Download)
			// if err != nil {
			// 	fmt.Println(err.Error())
			// 	return
			// }
			DlFile = answers.Download

			// create file from survey
			CreateFile(NewFileName, TemplateName, FieldMap, DlFile)

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

func CreateFile(NewFileName string, TemplateName string, FieldMap map[string]string, DlFile bool) {
	// print out the name of the file being downloaded
	fmt.Printf("Creating: %v\n", NewFileName)
	// print out the name of the template being used
	fmt.Printf("Using template: %v\n", TemplateName)
	// Get Template Id from the template name
	templateId, err := drive.GetFileId(TemplateName)
	if err != nil {
		log.Fatal(err)
	}
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
		fmt.Printf("%v Downloaded", NewFileName)
	}
}

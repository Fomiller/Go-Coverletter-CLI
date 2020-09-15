package cmd

import (
	"fmt"
	"log"

	"github.com/fomiller/scribe/utils"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Return a list of fields from a template",
	Long:  `Call Parse command to return a list of fields inside the specified template.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if template flag is not used prompt for the template name
		if TemplateName == "" {
			prompt := &survey.Input{
				Message: "What Template would you like to use",
			}
			survey.AskOne(prompt, &TemplateName)
		}
		// get the docID of the template that needs to be parsed
		templateId, err := drive.GetFileId(TemplateName)
		if err != nil {
			// log.Fatal(err)
			log.Fatalf("File could not be found, %v", err)
		}
		// insert parsedID and return []string of fields in the template
		parsedFields := drive.ParseTemplateFields(templateId)
		// base command needs to stop here.
		for _, v := range parsedFields {
			fmt.Println(v)
		}

		// All extra needs to be in seperate command or own function
		qs := []*survey.Question{}
		for _, v := range parsedFields {
			q := &survey.Question{
				Name: v,
				Prompt: &survey.Input{
					Message: fmt.Sprintf("%v:", v),
				},
				Validate: survey.Required,
			}
			qs = append(qs, q)
		}
		fmt.Println("Enter data for the corresponding fields in your template")
		err = survey.Ask(qs, &TemplateData)
		if err != nil {
			log.Fatal(err)
		}

		FieldMap = utils.StrIntfToStrStr(TemplateData)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
	// Here you will define your flags and configuration settings.
}

package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fomiller/scribe/drive"
	"github.com/spf13/cobra"
)

// showfieldsCmd represents the showfields command
var showfieldsCmd = &cobra.Command{
	Use:   "showfields",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if template flag is not used prompt for the template name
		if TemplateName == "" {
			prompt := &survey.Input{
				Message: "What Template would you like to use",
			}
			survey.AskOne(prompt, &TemplateName)
			fmt.Println("\n")
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
	},
}

func init() {
	rootCmd.AddCommand(showfieldsCmd)
	// Here you will define your flags and configuration settings.
}

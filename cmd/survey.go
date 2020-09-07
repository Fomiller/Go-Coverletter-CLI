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

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// the questions to ask
var qs = []*survey.Question{
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
	{
		Name: "download",
		Prompt: &survey.Confirm{
			Message: "Do you want to download this file?",
			Default: false,
		},
		Validate: survey.Required,
	},
}

// surveyCmd represents the survey command
var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			FileName     string // survey will match the question and field names
			TemplateName string `survey:"templateName"` // or you can tag fields to match a specific name
			Download     bool   // if the types don't match, survey will convert it
		}{}

		// perform the questions
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("%s chose %s, and download was set to %v.", answers.FileName, answers.TemplateName, answers.Download)

	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surveyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surveyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

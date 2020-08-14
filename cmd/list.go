package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list out files in your drive.",
	Long: `List will return you a list of the files located in your google drive without a search criteria.
	
	Use the Fields flag -f to specify fields inside your template that need to be replaced with data.
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
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

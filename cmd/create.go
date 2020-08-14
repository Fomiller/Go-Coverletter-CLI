package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var FieldMap map[string]string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new file using a template",
	Long: `Create will create a new file based off of a template name.

Use the list command flags to automatically fill in placeholders inside the template.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		for k, val := range FieldMap {
			fmt.Println("Key: ", strings.ToUpper(k), "Value: ", val)
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

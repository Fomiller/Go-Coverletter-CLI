package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name    string
	company string
	job     string

	// printCmd represents the print command
	printCmd = &cobra.Command{
		Use:   "print",
		Short: "Prints different arguments",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name ", name)
			fmt.Println("company ", company)
			fmt.Println("job", job)
		},
	}
)

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().StringVarP(&name, "name", "n", "", "Insert your name here.")
	// printCmd.Flags().StringVarP(&company, "company", "c", "", "Insert company name.")
	// printCmd.Flags().StringVarP(&job, "job", "j", "", "Insert job title.")
	printCmd.Flags().StringP("name", "n", "", "Insert your name here.")
	printCmd.Flags().StringP("company", "c", "", "Insert company name.")
	printCmd.Flags().StringP("job", "j", "", "Insert job title.")
}

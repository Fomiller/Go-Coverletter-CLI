package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fomiller/scribe/config"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Uservalue:%v\n", config.Scribe.Download.Path)
		if config.Scribe.Download.Path != "" {
			log.Printf("Using user download Path:%v\n", config.Scribe.Download.Path)
		} else {
			dir, _ := os.UserHomeDir()
			log.Printf("Using Default Download path:%v", dir)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	// Here you will define your flags and configuration settings.
}

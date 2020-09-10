package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var userLicense string
var TemplateName string
var NewFileName string
var DlFile bool
var FieldMap map[string]string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scribe",
	Short: "A CLI to help you work with templates.",
	Long:  `Scribe is a CLI for interfacing with templates using the Google Docs and Google Drive API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var command string
		prompt := &survey.Select{
			Message: "Choose a command to run:",
			Options: []string{"create", "parse", "delete"},
		}
		survey.AskOne(prompt, &command)

		// based on survey selection run the corresponding command
		switch command {
		case "create":
			createCmd.Run(cmd, args)
		case "parse":
			parseCmd.Run(cmd, args)
		case "delete":
			deleteCmd.Run(cmd, args)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.coverletter.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&DlFile, "download", "d", false, "download file")
	rootCmd.PersistentFlags().StringVarP(&TemplateName, "template", "t", "", "Enter the name of the template you would like to use.")
	rootCmd.PersistentFlags().StringVarP(&NewFileName, "name", "n", "", "Enter the name of the new file.")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".coverletter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".coverletter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

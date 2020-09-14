package cmd

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
)

type ScribeConfig struct {
	DownloadOutput struct {
		Default          string `yaml:"default,omitempty"`
		User             string `yaml:"user,omitempty"`
		FolderGeneration bool   `yaml:"folderGeneration,omitempty"`
	} `yaml:"downloadOutput,omitempty"`
	DocCredentials struct {
		Installed struct {
			ClientID                string   `yaml:"client_id,omitempty"`
			ProjectID               string   `yaml:"project_id,omitempty"`
			AuthURI                 string   `yaml:"auth_uri,omitempty"`
			TokenURI                string   `yaml:"token_uri,omitempty"`
			AuthProviderX509CertURL string   `yaml:"auth_provider_x509_cert_url,omitempty"`
			ClientSecret            string   `yaml:"client_secret,omitempty"`
			RedirectUris            []string `yaml:"redirect_uris,omitempty"`
		} `yaml:"installed,omitempty"`
	} `yaml:"docCredentials,omitempty"`
	DriveCredentials struct {
		Installed struct {
			ClientID                string   `yaml:"client_id,omitempty"`
			ProjectID               string   `yaml:"project_id,omitempty"`
			AuthURI                 string   `yaml:"auth_uri,omitempty"`
			TokenURI                string   `yaml:"token_uri,omitempty"`
			AuthProviderX509CertURL string   `yaml:"auth_provider_x509_cert_url,omitempty"`
			ClientSecret            string   `yaml:"client_secret,omitempty"`
			RedirectUris            []string `yaml:"redirect_uris,omitempty"`
		} `yaml:"installed,omitempty"`
	} `yaml:"driveCredentials,omitempty"`
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		filename, err := filepath.Abs("./config/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		var cfg ScribeConfig
		err = yaml.Unmarshal(f, &cfg)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	// Here you will define your flags and configuration settings.
}

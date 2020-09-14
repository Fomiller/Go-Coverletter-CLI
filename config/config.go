package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ScribeConfig struct {
	DownloadOutput struct {
		Default          string `yaml:"default,omitempty"`
		User             string `yaml:"user,omitempty"`
		FolderGeneration bool   `yaml:"folderGeneration,omitempty"`
	} `yaml:"downloadOutput,omitempty"`
	DocCredentials   string `yaml:"docCredentials,omitempty"`
	DriveCredentials string `yaml:"driveCredentials,omitempty"`
}

var Config = ScribeConfig{}

func init() {
	filename, err := filepath.Abs("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, &Config)
	if err != nil {
		log.Fatal(err)
	}
}

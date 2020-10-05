package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ScribeConfig struct {
	Download struct {
		// Default          string `yaml:"default,omitempty"`
		Path             string `yaml:"path,omitempty"`
		FolderGeneration bool   `yaml:"folderGeneration,omitempty"`
		UsePrefix        bool   `yaml:"usePrefix,omitempty"`
	} `yaml:"download,omitempty"`
	Credentials struct {
		Docs   string `yaml:"docs,omitempty"`
		Drive  string `yaml:"drive,omitempty"`
		Sheets string `yaml:"sheets,omitempty"`
	} `yaml:"credentials,omitempty"`
}

var Scribe = ScribeConfig{}

func init() {
	filename, err := filepath.Abs("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, &Scribe)
	if err != nil {
		log.Fatal(err)
	}
}

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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
var GoPath = strings.Trim(fmt.Sprintln(os.Getenv("GOPATH")), "\n")

func init() {
	// this allows scribe to be run in any directory by giving it a full path to read the config.yaml file
	configPath := filepath.Join(GoPath, "\\src\\github.com\\fomiller\\scribe\\config\\config.yaml")
	filename, err := filepath.Abs(configPath)
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

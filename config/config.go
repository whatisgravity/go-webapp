package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var State ConfigState

type ConfigState struct {
	Debug        bool   `yaml:"debug"`
	Port         string `yaml:"port"`
	APIVersion   string `yaml:"api_version"`
	DatabasePath string `yaml:"database_path"`
}

func Load(filePath string) error {
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configFile, &State)
	return err
}

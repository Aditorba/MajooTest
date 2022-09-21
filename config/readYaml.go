package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type ConfigData struct {
	Database struct {
		DBName string `yaml:"db"`
	} `yaml:"database"`

	Secret struct {
		KeyGenerate string `yaml:"keyGenerate"`
	} `yaml:"secret"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var configResult *ConfigData

func PopulateConfigData(filePath string) *ConfigData {
	fmt.Println("read yml config")

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	err = yaml.Unmarshal(content, &configResult)
	check(err)

	return configResult
}

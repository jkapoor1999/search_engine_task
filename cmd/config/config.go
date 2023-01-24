package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configure struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Protocol   string `yaml:"protocol"`
		DBName     string `yaml:"dbname"`
		Collection string `yaml:"collection"`
	} `yaml:"database"`
}

var Config Configure

func LoadConfig() {

	data, err := os.ReadFile("./cmd/config/config.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}
}

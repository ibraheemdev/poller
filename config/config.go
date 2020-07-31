package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Config :
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

// ReadFile :
func ReadFile() Config {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	return cfg
}

package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// Config :
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// DatabaseConfig :
type DatabaseConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

// ServerConfig :
type ServerConfig struct {
	Host   string       `yaml:"host"`
	Port   string       `yaml:"port"`
	API    APIConfig    `yaml:"api"`
	Static StaticConfig `yaml:"static"`
}

// APIConfig :
type APIConfig struct {
	Domain string `yaml:"domain"`
}

// StaticConfig :
type StaticConfig struct {
	Domain    string `yaml:"domain"`
	HomePage  string `yaml:"homepage"`
	BuildPath string `yaml:"buildpath"`
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

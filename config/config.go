package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// EnvironmentConfig :
type EnvironmentConfig struct {
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
func ReadFile() EnvironmentConfig {
	file := fmt.Sprintf("config/environments/%s.yml", GetEnv())
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer f.Close()

	var cfg EnvironmentConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	return cfg
}

// GetEnv :
func GetEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("POLLER_ENV")
}

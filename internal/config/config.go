package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerPort string `yaml:"ServerPort"`
	Database   struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Name     string `yaml:"Name"`
	} `yaml:"Database"`
	JWTSecret string `yaml:"JWTSecret"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

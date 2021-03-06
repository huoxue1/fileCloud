package config

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed default.config.yml
var defaultConfig []byte

type Config struct {
	Address     string `json:"address" yaml:"address"`
	Port        int    `json:"port" yaml:"port"`
	StoragePath string `json:"storage_path" yaml:"storage_path"`
	DBPath      string `json:"db_path" yaml:"db_path"`
}

func GetConfig() *Config {
	return config
}

var config = new(Config)

func LoadConfig(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return err
	}
	return err
}

func CheckConfigFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err := os.WriteFile(path, defaultConfig, 0666)
		if err != nil {
			return
		}
		return
	}
}

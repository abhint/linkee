package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		DatabaseConfig `yaml:"databaseConfig"`
		HostConfig     `yaml:"hostConfig"`
	}
	DatabaseConfig struct {
		DataSourceName string `yaml:"dataSourceName"`
	}
	HostConfig struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
)

func LoadConfig(filename string) (config *Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return config, nil
}

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database"`
	DatabaseRead struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"database-read"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Aws struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"aws"`
}

func NewConfig(folder string) (*Config, error) {
	file, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

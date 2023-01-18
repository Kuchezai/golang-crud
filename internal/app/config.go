package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBbname  string `yaml:"dbname"`
		SSLmode  string `yaml:"sslmode"`
	}
}

func New() *Config {
	config := &Config{}
	return config
}
func Load() *Config {
	configuration := New()
	cfgFile, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(cfgFile, configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

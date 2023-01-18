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

func Load() *Config {
	configuration := &Config{}
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

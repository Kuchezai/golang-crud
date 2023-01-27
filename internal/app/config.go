package config

import (
	logger "CRUD/internal/app/logs"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBbname  string `yaml:"dbname"`
		SSLmode  string `yaml:"sslmode"`
	}
	JWT struct {
		Key string `yaml:"key"`
	}
}

var instantiated *config = nil

func GetConfig() **config {
	if instantiated == nil {
		instantiated = load()
	}
	return &instantiated
}

func load() *config {
	config := &config{}
	cfgFile, err := ioutil.ReadFile("configs/config.yaml")
	if err != nil {
		logger.Error.Println("Error in parsing config!")
		panic(err)
	}
	err = yaml.Unmarshal(cfgFile, config)
	if err != nil {
		panic(err)
	}
	return config
}

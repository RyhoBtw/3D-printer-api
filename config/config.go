package config

import (
	"io/ioutil"

	"github.com/RyhoBtw/3D-printer-api/log"
	"github.com/pelletier/go-toml"
)

var config *Config

type Config struct {
	RaspberryPiIp string
	DbIp          string
	LogLevel      string
}

func LoadConfig() {
	c := Config{}

	b, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Log().Infoln("Config not found, creating new")
		data, err := toml.Marshal(c)
		if err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			panic(err)
		}
		return
	}
	if err := toml.Unmarshal(b, &c); err != nil {
		panic(err)
	}
	if c.LogLevel == "" {
		c.LogLevel = "info"
	}
	log.Log().Infoln("Config loaded")
	config = &c
}

func GetConfig() *Config {
	return config
}

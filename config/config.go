package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Bot struct {
		Token string `json:"token"`
	} `json:"bot"`

	Database struct {
		Dsn string `json:"dsn"`
	} `json:"database"`
}

var Config config

func Setup(configPath string) {

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Cannot read a config file. Path: %v. Error: %v", configPath, err.Error())
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Cannot unmarshal a json config. Error: %v", err.Error())
	}
}

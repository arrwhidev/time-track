package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	spreadsheetId string `json:spreadsheetId`
	email         string `json:email`
	privateKey    string `json:privateKey`
	privateKeyId  string `json:privateKeyId`
}

func loadConfig() *Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

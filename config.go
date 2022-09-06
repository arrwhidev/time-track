package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	SpreadsheetId string   `json:spreadsheetId`
	Email         string   `json:email`
	PrivateKey    string   `json:privateKey`
	PrivateKeyId  string   `json:privateKeyId`
	Mappings      []string `json:mappings`
}

func loadConfig() *Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	bytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

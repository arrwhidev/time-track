package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Add(key string, value int) {
	ensureTodayExists()

	data := readToday()
	if val, ok := data[key]; ok {
		data[key] = value + val
	} else {
		data[key] = value
	}

	writeToday(data)
}

func Set(key string, value int) {
	ensureTodayExists()

	data := readToday()
	data[key] = value
	writeToday(data)
}

func today() string {
	return time.Now().Format("01-02-2006")
}

func todayPath() string {
	return "/tmp/tt/" + today()
}

func touchToday() {
	err := os.MkdirAll("/tmp/tt/", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(todayPath())
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("{}")
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func readToday() map[string]int {
	byteValue, err := ioutil.ReadFile(todayPath())
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]int
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func writeToday(data map[string]int) {
	byteValue, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(todayPath(), byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ensureTodayExists() {
	_, err := os.Stat(todayPath())
	if !os.IsNotExist(err) {
		touchToday()
	}
}

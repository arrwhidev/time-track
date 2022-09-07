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

func GetLastSyncTime() time.Time {
	ensureLastSyncExists()

	bytes, err := ioutil.ReadFile(lastSyncPath())
	if err != nil {
		log.Fatal(err)
	}

	var lastSync time.Time
	err = lastSync.UnmarshalText(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return lastSync
}

func UpdateLastSync() {
	ensureLastSyncExists()

	now := time.Now()
	bytes, err := now.MarshalText()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(lastSyncPath(), bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func today() string {
	return time.Now().Format("01-02-2006")
}

func todayDataFilePath() string {
	return "/tmp/tt/" + today()
}

func lastSyncPath() string {
	return "/tmp/tt/.last_sync"
}

func createFile(path string, initialValue string) {
	err := os.MkdirAll("/tmp/tt/", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(initialValue)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func readToday() map[string]int {
	bytes, err := ioutil.ReadFile(todayDataFilePath())
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]int
	err = json.Unmarshal(bytes, &result)
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

	err = ioutil.WriteFile(todayDataFilePath(), byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ensureTodayExists() {
	_, err := os.Stat(todayDataFilePath())
	if os.IsNotExist(err) {
		createFile(todayDataFilePath(), "{}")
	}
}

func ensureLastSyncExists() {
	_, err := os.Stat(lastSyncPath())
	if os.IsNotExist(err) {
		createFile(lastSyncPath(), "")
	}
}

func isDirty() bool {
	dataFileInfo, err := os.Stat(todayDataFilePath())
	if err != nil {
		log.Fatal(err)
	}
	dataFileModTime := dataFileInfo.ModTime()
	lastSyncTime := GetLastSyncTime()
	return dataFileModTime.After(lastSyncTime)
}

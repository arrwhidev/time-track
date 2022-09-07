package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

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

func lastSyncPath() string {
	return "/tmp/tt/.last_sync"
}

func ensureLastSyncExists() {
	_, err := os.Stat(lastSyncPath())
	if os.IsNotExist(err) {
		createFile(lastSyncPath(), "")
	}
}

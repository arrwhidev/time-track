package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type JsonFile[T any] struct {
	path string
	data *T
}

func OpenJsonFile[T any](path string) *JsonFile[T] {
	_, err := os.Stat(path)

	data := new(T)
	if os.IsNotExist(err) {
		createFile(path, "{}")
	} else {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bytes, &data)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &JsonFile[T]{path, data}
}

func (f *JsonFile[T]) Write() {
	byteValue, err := json.Marshal(f.data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(f.path, byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *JsonFile[T]) ModTime() time.Time {
	fileInfo, err := os.Stat(f.path)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfo.ModTime()
}

func createFile(path string, initialValue string) {
	err := os.MkdirAll("/tmp/tt/", os.ModePerm) // TODO: grab the path from `path`
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

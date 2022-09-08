package main

import (
	"time"
)

type DataFileWrapper struct {
	jf *JsonFile[map[string]int]
}

func GetDataFile() *DataFileWrapper {
	jf := OpenJsonFile[map[string]int]("/tmp/tt/" + today())
	return &DataFileWrapper{jf}
}

func (w *DataFileWrapper) Data() *map[string]int {
	if len(*w.jf.data) == 0 {
		*w.jf.data = make(map[string]int)
	}

	return w.jf.data
}

func (w *DataFileWrapper) IsDirty(appConfig AppConfig) bool {
	return w.jf.ModTime().After(appConfig.LastSync)
}

func (w *DataFileWrapper) Add(key string, value int) {
	data := (*w.Data())
	if val, ok := data[key]; ok {
		data[key] = value + val
	} else {
		data[key] = value
	}

	w.jf.Write()
}

func (w *DataFileWrapper) Set(key string, value int) {
	(*w.Data())[key] = value
	w.jf.Write()
}

func today() string {
	return time.Now().Format("01-02-2006")
}

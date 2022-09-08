package main

import "time"

const (
	appConfigFilename = ".config.json"
)

type AppConfig struct {
	LastSync time.Time `json:lastSync`
}

type AppConfigWrapper struct {
	jf *JsonFile[AppConfig]
}

func GetAppConfig() *AppConfigWrapper {
	jf := OpenJsonFile[AppConfig]("/tmp/tt/" + appConfigFilename)
	return &AppConfigWrapper{jf}
}

func (w *AppConfigWrapper) Data() *AppConfig {
	return w.jf.data
}

func (w *AppConfigWrapper) Write() {
	w.jf.Write()
}

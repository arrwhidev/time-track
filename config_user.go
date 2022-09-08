package main

const (
	userConfigFilename = "config.json"
)

type UserConfig struct {
	SpreadsheetId string   `json:spreadsheetId`
	Email         string   `json:email`
	PrivateKey    string   `json:privateKey`
	PrivateKeyId  string   `json:privateKeyId`
	Mappings      []string `json:mappings`
}

type UserConfigWrapper struct {
	jf *JsonFile[UserConfig]
}

func GetUserConfig() *UserConfigWrapper {
	jf := OpenJsonFile[UserConfig]("/tmp/tt/" + userConfigFilename)
	return &UserConfigWrapper{jf}
}

func (w *UserConfigWrapper) Data() *UserConfig {
	return w.jf.data
}

package main

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	srv    *sheets.Service
	config Config
}

func (ss *SheetsService) Sync(data map[string]int) {
	numMappings := len(ss.config.Mappings)
	rangeData := "Sheet1!A:" + string('A'+numMappings)
	values := [][]interface{}{
		{today()},
	}

	for _, k := range ss.config.Mappings {
		if val, ok := data[k]; ok {
			values[0] = append(values[0], val)
		} else {
			values[0] = append(values[0], 0)
		}
	}

	valueRange := &sheets.ValueRange{
		Range:  rangeData,
		Values: values,
	}

	_, err := ss.srv.Spreadsheets.Values.Append(ss.config.SpreadsheetId, rangeData, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatal(err)
	}
}

func NewSheetsService(config Config) *SheetsService {
	conf := &jwt.Config{
		Email:        config.Email,
		PrivateKey:   []byte(config.PrivateKey),
		PrivateKeyID: config.PrivateKeyId,
		TokenURL:     "https://oauth2.googleapis.com/token",
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
	}

	client := conf.Client(oauth2.NoContext)
	service, err := sheets.New(client)
	if err != nil {
		log.Fatal(err)
	}

	return &SheetsService{service, config}
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	srv    *sheets.Service
	config UserConfigWrapper
}

func NewSheetsService(config UserConfigWrapper) *SheetsService {
	data := config.Data()
	conf := &jwt.Config{
		Email:        data.Email,
		PrivateKey:   []byte(data.PrivateKey),
		PrivateKeyID: data.PrivateKeyId,
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

func (ss *SheetsService) Sync(data map[string]int, appConfig AppConfigWrapper) {
	config := ss.config.Data()
	numMappings := len(config.Mappings)
	values := [][]interface{}{
		{today()},
	}

	for _, k := range config.Mappings {
		if val, ok := data[k]; ok {
			values[0] = append(values[0], val)
		} else {
			values[0] = append(values[0], 0)
		}
	}

	if appConfig.Data().LastSync.Before(midnight()) {
		// Last sync was yesterday, so need to append a new row
		rangeData := "Sheet1!A:" + string('A'+numMappings)
		valueRange := &sheets.ValueRange{
			Range:  rangeData,
			Values: values,
		}

		row, err := ss.srv.Spreadsheets.Values.Append(config.SpreadsheetId, rangeData, valueRange).ValueInputOption("USER_ENTERED").Do()
		if err != nil {
			log.Fatal(err)
		}

		rowNumber := getRowNumberFromRange(row.Updates.UpdatedRange)
		appConfig.Data().RowNumber = rowNumber
		appConfig.Write()
	} else {
		// Last sync was today, so need to update existing row
		rowNumber := appConfig.Data().RowNumber
		rangeData := fmt.Sprint("Sheet1!A", rowNumber, ":", string('A'+numMappings), rowNumber)
		valueRange := &sheets.ValueRange{
			Range:  rangeData,
			Values: values,
		}

		_, err := ss.srv.Spreadsheets.Values.Update(config.SpreadsheetId, rangeData, valueRange).ValueInputOption("USER_ENTERED").Do()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getRowNumberFromRange(rangeValue string) int {
	parts := strings.Split(rangeValue, ":D") // TODO: compute `D` instead of hardcoding.
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func midnight() time.Time {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal(err)
	}
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

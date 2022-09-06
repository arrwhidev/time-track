package main

import (
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct {
	srv           *sheets.Service
	spreadsheetId string
}

func NewSheetsService(config Config) *SheetsService {
	spreadsheetId := config.spreadsheetId
	conf := &jwt.Config{
		Email:        config.email,
		PrivateKey:   []byte(config.privateKey),
		PrivateKeyID: config.privateKeyId,
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

	return &SheetsService{service, spreadsheetId}
}

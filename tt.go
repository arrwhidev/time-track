package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "set",
				Usage: "Set time",
				Action: func(cCtx *cli.Context) error {
					category := cCtx.Args().First()
					time, err := ParseInputToMinutes(cCtx.Args().Get(1))
					if err != nil {
						log.Fatal(err)
					}

					dataFile := GetDataFile()
					dataFile.Set(category, time)

					return nil
				},
			}, {
				Name:  "add",
				Usage: "Add time",
				Action: func(cCtx *cli.Context) error {
					category := cCtx.Args().First()
					time, err := ParseInputToMinutes(cCtx.Args().Get(1))
					if err != nil {
						log.Fatal(err)
					}

					dataFile := GetDataFile()
					dataFile.Add(category, time)

					return nil
				},
			}, {
				Name:  "sync",
				Usage: "Sync to Google Sheet",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "force", Aliases: []string{"F"}},
				},
				Action: func(cCtx *cli.Context) error {
					appConfig := GetAppConfig()
					dataFile := GetDataFile()

					if dataFile.IsDirty(*appConfig.Data()) || cCtx.Bool("force") {
						userConfig := GetUserConfig()
						ss := NewSheetsService(*userConfig)
						if ss == nil {
							log.Fatal("Failed to create sheets service")
						}
						ss.Sync(*dataFile.Data(), *appConfig)

						appConfig.Data().LastSync = time.Now()
						appConfig.Write()
					} else {
						fmt.Println("Already up to date!")
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

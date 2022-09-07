package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
					time, err := strconv.Atoi(cCtx.Args().Get(1))
					if err != nil {
						log.Fatal(err)
					}
					Set(category, time)

					return nil
				},
			}, {
				Name:  "add",
				Usage: "Add time",
				Action: func(cCtx *cli.Context) error {
					category := cCtx.Args().First()
					time, err := strconv.Atoi(cCtx.Args().Get(1))
					if err != nil {
						log.Fatal(err)
					}
					Add(category, time)

					return nil
				},
			}, {
				Name:  "sync",
				Usage: "Sync to Google Sheet",
				Action: func(cCtx *cli.Context) error {
					if isDirty() {
						config := loadConfig()
						ss := NewSheetsService(*config)
						if ss == nil {
							log.Fatal("Failed to create sheets service")
						}
						ss.Sync(readToday())
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

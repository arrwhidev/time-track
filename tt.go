package main

import (
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
			},
			{
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

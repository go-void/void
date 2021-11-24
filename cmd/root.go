package cmd

import (
	"os"

	"github.com/go-void/void/internal/app"

	"github.com/urfave/cli/v2"
)

var configFilePath string

func Execute() error {
	app := &cli.App{
		Name:     "void",
		Usage:    "void runs a DNS level adblocker",
		Commands: []*cli.Command{},
		Action: func(c *cli.Context) error {
			a, err := app.New(configFilePath)
			if err != nil {
				return err
			}

			return a.Run()
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Path to config file",
				Destination: &configFilePath,
			},
		},
	}

	return app.Run(os.Args)
}

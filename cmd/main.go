package main

import (
	"context"
	"github.com/QinLiStudio/Conship/internal/app"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "conship"
	app.Usage = "Conship based on GIN + GORM + WIRE."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Error(err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run http server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Usage:    "App configuration file(.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("config")))
		},
	}
}

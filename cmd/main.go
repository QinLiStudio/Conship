package cmd

import (
	"github.com/QinLiStudio/Conship/internal/app"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func main() {
	ctx := app.SugarLogger

	app := cli.NewApp()
	app.Name = "conship"
	app.Usage = "Conship based on GIN + GORM + WIRE."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}

	err := app.Run(os.Args)
	if err != nil {
		ctx.Errorf(err.Error())
	}
}

func newWebCmd(ctx *zap.SugaredLogger) *cli.Command {

}

package main

import (
	"os"

	"github.com/k3s-io/k3s/pkg/cli/cmds"
	"github.com/k3s-io/k3s/pkg/cli/token"
	"github.com/k3s-io/k3s/pkg/configfilearg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		cmds.NewTokenCommands(
			token.Create,
			token.Delete,
			token.Generate,
			token.List,
			token.Rotate,
		),
	}

	cmds.MustRun(app, configfilearg.MustParse(os.Args))
}

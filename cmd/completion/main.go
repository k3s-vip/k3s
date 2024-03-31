package main

import (
	"os"

	"github.com/k3s-io/k3s/pkg/cli/cmds"
	"github.com/k3s-io/k3s/pkg/cli/completion"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cmds.NewApp()
	app.Commands = []*cli.Command{
		cmds.NewCompletionCommand(
			completion.Bash,
			completion.Zsh,
		),
	}

	cmds.MustRun(app, os.Args)
}

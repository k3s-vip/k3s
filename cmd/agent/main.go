package main

import (
	"os"

	"github.com/k3s-io/k3s/pkg/cli/agent"
	"github.com/k3s-io/k3s/pkg/cli/cmds"
	"github.com/k3s-io/k3s/pkg/configfilearg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cmds.NewApp()
	app.DisableSliceFlagSeparator = true
	app.Commands = []*cli.Command{
		cmds.NewAgentCommand(agent.Run),
	}

	cmds.MustRun(app, configfilearg.MustParse(os.Args))
}

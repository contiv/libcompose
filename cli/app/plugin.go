package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/contiv/deploy/hooks"
	"github.com/docker/libcompose/project"
)

func plugin(p *project.Project, context *cli.Context) error {
	cliLabels := ""
	event := project.NoEvent
	switch context.Command.Name {
	case "up", "start":
		event = project.EventProjectUpStart
	case "down", "delete", "kill", "rm", "stop":
		event = project.EventProjectDownStart
	case "create", "build", "ps", "port", "pull", "log", "restart":
	}

	if event == project.NoEvent {
		return nil
	}

	if event == project.EventProjectUpStart && cliLabels != "" {
		if err := hooks.PopulateEnvLabels(p, cliLabels); err != nil {
			logrus.Fatalf("Unable to insert environment labels. Error %v", err)
		}
	}

	if err := hooks.NetHooks(p, event); err != nil {
		logrus.Fatalf("Unable to generate network labels. Error %v", err)
	}

	return nil
}

package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/contiv/deploy/hooks"
	"github.com/docker/libcompose/project"
)

func pre_plugin(p *project.Project, context *cli.Context) error {
	cliLabels := ""

	// TODO: parse cliLabels from cli.Context
	if cliLabels != "" {
		if err := hooks.PopulateEnvLabels(p, cliLabels); err != nil {
			logrus.Fatalf("Unable to insert environment labels. Error %v", err)
		}
	}

	if err := hooks.PreHooks(p, context.Command.Name); err != nil {
		logrus.Fatalf("Unable to generate network labels. Error %v", err)
	}

	return nil
}

func post_plugin(p *project.Project, context *cli.Context) error {
	if err := hooks.PostHooks(p, context.Command.Name); err != nil {
		logrus.Fatalf("Unable to populate dns entries. Error %v", err)
	}
	return nil
}

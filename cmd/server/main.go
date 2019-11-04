package main

import (
	"context"
	"github.com/aimuz/fishecho/pkg/server"
	"github.com/aimuz/fishecho/pkg/server/cmd"
	"github.com/aimuz/fishecho/pkg/version"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	command := cmd.NewServerDefaultCommand()
	command.RunE = run
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}

func run(_ *cobra.Command, _ []string) error {
	logrus.Infof("Starting fishecho, version: %s, git commit: %s", version.Version, version.GitCommit)
	ctx := signals.SetupSignalHandler(context.Background())
	return server.Start(ctx)
}

package main

import (
	"context"
	"github.com/aimuz/fishecho/pkg/server"
	"github.com/aimuz/fishecho/pkg/server/cmd"
	"github.com/rancher/wrangler/pkg/signals"
	"github.com/spf13/cobra"
)

func main() {
	ServerCmd := cmd.NewDefaultCommand()
	ServerCmd.RunE = run
	err := ServerCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	ctx := signals.SetupSignalHandler(context.Background())
	return server.Start(ctx)
}

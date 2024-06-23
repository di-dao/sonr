package main

import (
	"github.com/di-dao/sonr/pkg/vault"
	"github.com/spf13/cobra"
)

func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:                        "vltd",
		Aliases:                    []string{"vault"},
		Short:                      "run the vault rest api and htmx frontend",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		Run: func(cmd *cobra.Command, args []string) {
			vault.Serve(cmd.Context())
			select {}
		},
	}
}

func main() {
	if err := serveCmd().Execute(); err != nil {
		panic(err)
	}
}

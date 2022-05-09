package api

import "github.com/spf13/cobra"

func NewAPICommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api [service]",
		Short: "start api's service",
	}

	cmd.AddCommand(newAPICommand())

	return cmd
}

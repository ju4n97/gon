package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCommand creates a new cobra.Command for printing the current version of the application.
func NewVersionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print the current version of the application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version: 0.0.1")
		},
	}

	return command
}

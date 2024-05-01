package main

import (
	"github.com/mesatechlabs/kitten/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A brief description of your application",
	}

	rootCmd.AddCommand(cmd.NewVersionCommand())
	rootCmd.AddCommand(cmd.NewServeCommand())

	rootCmd.Execute()
}

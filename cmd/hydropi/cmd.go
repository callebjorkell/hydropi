package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hydropi",
		Short: "raspberry pi based watering system",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().Lookup("debug").Changed {
				log.SetLevel(log.DebugLevel)
			}
		},
	}

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newStartCmd())
	rootCmd.PersistentFlags().Bool("debug", false, "Turn on debug logging.")

	return rootCmd
}

var (
	buildTime    = "unknown"
	buildVersion = "dev"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (built: %s)\n", buildVersion, buildTime)
		},
	}
}

func newStartCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}

	return &cmd
}

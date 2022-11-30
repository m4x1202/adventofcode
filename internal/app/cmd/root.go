package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:              "aoc",
		Short:            "All in one AdventOfCode binary",
		TraverseChildren: true,
	}
	verbose int = 0
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.Flags().CountVarP(&verbose, "verbose", "v", "verbose output")
}

func initConfig() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch verbose {
	case 0:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 3:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}
	log.WithLevel(zerolog.GlobalLevel()).Msgf("logging level set to %v", zerolog.GlobalLevel())
}

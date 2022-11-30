package cmd2021

import (
	"github.com/m4x1202/adventofcode/internal/app/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(cmd2021)
}

var (
	cmd2021 = &cobra.Command{
		Use:   "2021",
		Short: "2021 puzzles",
	}
)

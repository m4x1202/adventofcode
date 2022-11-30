package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(cmd2022)
}

var (
	cmd2022 = &cobra.Command{
		Use:   "2022",
		Short: "2022 puzzles",
	}
)

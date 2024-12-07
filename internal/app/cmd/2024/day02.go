package cmd2024

import (
	"github.com/m4x1202/adventofcode/internal/app/2024/day02"
	"github.com/spf13/cobra"
)

func init() {
	cmd2024.AddCommand(day02Cmd)
}

var (
	day02Cmd = &cobra.Command{
		Use:       "day02 <part1|part2>",
		Short:     "Day 02 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day02.ExecutePart(1)
			case "part2":
				day02.ExecutePart(2)
			}
		},
	}
)

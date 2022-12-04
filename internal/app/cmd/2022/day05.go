package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day05"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day05Cmd)
}

var (
	day05Cmd = &cobra.Command{
		Use:       "day05 <part1|part2>",
		Short:     "Day 05 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day05.ExecutePart(1)
			case "part2":
				day05.ExecutePart(2)
			}
		},
	}
)

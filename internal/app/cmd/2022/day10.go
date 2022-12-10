package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day10"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day10Cmd)
}

var (
	day10Cmd = &cobra.Command{
		Use:       "day10 <part1|part2>",
		Short:     "Day 10 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day10.ExecutePart(1)
			case "part2":
				day10.ExecutePart(2)
			}
		},
	}
)

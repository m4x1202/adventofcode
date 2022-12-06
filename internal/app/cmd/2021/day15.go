package cmd2021

import (
	"github.com/m4x1202/adventofcode/internal/app/2021/day15"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day15Cmd)
}

var (
	day15Cmd = &cobra.Command{
		Use:       "day15 <part1|part2>",
		Short:     "Day 15 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day15.ExecutePart(1)
			case "part2":
				day15.ExecutePart(2)
			}
		},
	}
)

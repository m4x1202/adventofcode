package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day09"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day09Cmd)
}

var (
	day09Cmd = &cobra.Command{
		Use:       "day09 <part1|part2>",
		Short:     "Day 09 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day09.ExecutePart(1)
			case "part2":
				day09.ExecutePart(2)
			}
		},
	}
)

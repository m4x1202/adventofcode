package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day07"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day07Cmd)
}

var (
	day07Cmd = &cobra.Command{
		Use:       "day07 <part1|part2>",
		Short:     "Day 07 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day07.ExecutePart(1)
			case "part2":
				day07.ExecutePart(2)
			}
		},
	}
)

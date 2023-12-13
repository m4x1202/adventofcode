package cmd2023

import (
	"fmt"

	"github.com/m4x1202/adventofcode/internal/app/2023/day04"
	"github.com/spf13/cobra"
)

func init() {
	cmd2023.AddCommand(day04Cmd)
}

var (
	day04Cmd = &cobra.Command{
		Use:       "day04 <part1|part2>",
		Short:     "Day 04 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			var result uint64
			switch args[0] {
			case "part1":
				result = day04.ExecutePart(1)
			case "part2":
				result = day04.ExecutePart(2)
			}
			fmt.Printf("Result: %d\n", result)
		},
	}
)

package cmd2023

import (
	"fmt"

	"github.com/m4x1202/adventofcode/internal/app/2023/day02"
	"github.com/spf13/cobra"
)

func init() {
	cmd2023.AddCommand(day02Cmd)
}

var (
	day02Cmd = &cobra.Command{
		Use:       "day02 <part1|part2>",
		Short:     "Day 02 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			var result uint64
			switch args[0] {
			case "part1":
				result = day02.ExecutePart(1)
			case "part2":
				result = day02.ExecutePart(2)
			}
			fmt.Printf("Result: %d\n", result)
		},
	}
)

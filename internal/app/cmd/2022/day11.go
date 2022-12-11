// day11 is the template package which can be copied, modified to its final location
package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day11"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day11Cmd)
}

var (
	day11Cmd = &cobra.Command{
		Use:       "day11 <part1|part2>",
		Short:     "Day 11 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day11.ExecutePart(1)
			case "part2":
				day11.ExecutePart(2)
			}
		},
	}
)

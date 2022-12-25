// day23 is the template package which can be copied, modified to its final location
package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day23"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day23Cmd)
}

var (
	day23Cmd = &cobra.Command{
		Use:       "day23 <part1|part2>",
		Short:     "Day 23 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day23.ExecutePart(1)
			case "part2":
				day23.ExecutePart(2)
			}
		},
	}
)

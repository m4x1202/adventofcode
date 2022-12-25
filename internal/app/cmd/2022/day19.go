// day19 is the template package which can be copied, modified to its final location
package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day19"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day19Cmd)
}

var (
	day19Cmd = &cobra.Command{
		Use:       "day19 <part1|part2>",
		Short:     "Day 19 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "part1":
				day19.ExecutePart(1)
			case "part2":
				day19.ExecutePart(2)
			}
		},
	}
)

// day00 is the template package which can be copied, modified to its final location
package cmd2023

import (
	"fmt"

	"github.com/m4x1202/adventofcode/internal/app/2023/day00"
	"github.com/spf13/cobra"
)

func init() {
	cmd2023.AddCommand(day00Cmd)
}

var (
	day00Cmd = &cobra.Command{
		Use:       "day00 <part1|part2>",
		Short:     "Day 00 Challenge",
		ValidArgs: []string{"part1", "part2"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			var result uint64
			switch args[0] {
			case "part1":
				result = day00.ExecutePart(1)
			case "part2":
				result = day00.ExecutePart(2)
			}
			fmt.Printf("Result: %d\n", result)
		},
	}
)

// day00 is the template package which can be copied, modified to its final location
package cmd2021

import (
	"github.com/m4x1202/adventofcode/internal/app/2021/day00"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day00Cmd)

	day00Cmd.AddCommand(day00part1Cmd)
	day00Cmd.AddCommand(day00part2Cmd)
}

var (
	day00Cmd = &cobra.Command{
		Use:   "day00",
		Short: "Day 00 Challenge",
	}
	day00part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 00 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day00.Part1(args)
		},
	}
	day00part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 00 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day00.Part2(args)
		},
	}
)

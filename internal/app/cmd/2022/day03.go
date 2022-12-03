package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day03"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day03Cmd)

	day03Cmd.AddCommand(day03part1Cmd)
	day03Cmd.AddCommand(day03part2Cmd)
}

var (
	day03Cmd = &cobra.Command{
		Use:   "day03",
		Short: "Day 03 Challenge",
	}
	day03part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 03 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day03.Part1(args)
		},
	}
	day03part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 03 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day03.Part2(args)
		},
	}
)

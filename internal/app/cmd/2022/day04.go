package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day04"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day04Cmd)

	day04Cmd.AddCommand(day04part1Cmd)
	day04Cmd.AddCommand(day04part2Cmd)
}

var (
	day04Cmd = &cobra.Command{
		Use:   "day04",
		Short: "Day 04 Challenge",
	}
	day04part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 04 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day04.Part1(args)
		},
	}
	day04part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 04 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day04.Part2(args)
		},
	}
)

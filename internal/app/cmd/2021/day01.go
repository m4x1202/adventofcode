package cmd2021

import (
	"github.com/m4x1202/adventofcode/internal/app/2021/day01"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day01Cmd)

	day01Cmd.AddCommand(day01part1Cmd)
	day01Cmd.AddCommand(day01part2Cmd)
}

var (
	day01Cmd = &cobra.Command{
		Use:   "day01",
		Short: "Day 01 Challenge",
	}
	day01part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 01 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day01.Part1(args)
		},
	}
	day01part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 01 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day01.Part2(args)
		},
	}
)

package cmd

import (
	"github.com/m4x1202/adventofcode/internal/app/day15"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day15Cmd)

	day15Cmd.AddCommand(day15part1Cmd)
	day15Cmd.AddCommand(day15part2Cmd)
}

var (
	day15Cmd = &cobra.Command{
		Use:   "day15",
		Short: "Day 15 Challenge",
	}
	day15part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 15 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day15.Part1(args)
		},
	}
	day15part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 15 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day15.Part2(args)
		},
	}
)

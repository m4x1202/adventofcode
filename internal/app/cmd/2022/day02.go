package cmd2022

import (
	"github.com/m4x1202/adventofcode/internal/app/2022/day02"
	"github.com/spf13/cobra"
)

func init() {
	cmd2022.AddCommand(day02Cmd)

	day02Cmd.AddCommand(day02part1Cmd)
	day02Cmd.AddCommand(day02part2Cmd)
}

var (
	day02Cmd = &cobra.Command{
		Use:   "day02",
		Short: "Day 02 Challenge",
	}
	day02part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 02 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day02.Part1(args)
		},
	}
	day02part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 02 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day02.Part2(args)
		},
	}
)

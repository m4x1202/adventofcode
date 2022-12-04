package cmd2021

import (
	"github.com/m4x1202/adventofcode/internal/app/2021/day14"
	"github.com/spf13/cobra"
)

func init() {
	cmd2021.AddCommand(day14Cmd)

	day14Cmd.AddCommand(day14part1Cmd)
}

var (
	day14Cmd = &cobra.Command{
		Use:   "day14",
		Short: "Day 14 Challenge",
	}
	day14part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 14 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day14.Part1()
		},
	}
)

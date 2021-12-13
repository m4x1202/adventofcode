package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func Abs(x int) uint32 {
	if x < 0 {
		return uint32(-x)
	}
	return uint32(x)
}

func nthTriangleNumber(x uint32) uint32 {
	return x * (x + 1) / 2
}

type CrabPositions []int

func (c *CrabPositions) ShiftLeft() {
	for i := 0; i < len(*c); i++ {
		(*c)[i]--
	}
}

func (c *CrabPositions) TotalFuel() uint32 {
	var fuel uint32
	for _, pos := range *c {
		fuel += Abs(pos)
	}
	return fuel
}

func (c *CrabPositions) TotalFuelExp() uint64 {
	var fuel uint64
	for _, pos := range *c {
		fuel += uint64(nthTriangleNumber(Abs(pos)))
	}
	return fuel
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	day7Cmd.AddCommand(day7part1Cmd)
	day7Cmd.AddCommand(day7part2Cmd)
}

var (
	day7logger = log.With().
			Int("day", 7).
			Logger()
	day7Cmd = &cobra.Command{
		Use:   "day7",
		Short: "Day 7 Challenge",
	}
	day7part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 7 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day7part1logger := day7logger.With().
				Int("part", 1).
				Logger()
			day7part1logger.Info().Msg("Start")
			converted := prepareday7Input()

			lowestFuel := converted.TotalFuel()
			highestPos := (*converted)[len(*converted)-1]
			for i := 1; i < highestPos; i++ {
				converted.ShiftLeft()
				fuel := converted.TotalFuel()
				day7part1logger.Debug().Msgf("new fuel: %d | lowest: %d", fuel, lowestFuel)
				if fuel > lowestFuel {
					break
				}
				lowestFuel = fuel
			}

			fmt.Printf("lowest: %d\n", lowestFuel)
		},
	}
	day7part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 7 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day7part2logger := day7logger.With().
				Int("part", 2).
				Logger()
			day7part2logger.Info().Msg("Start")
			converted := prepareday7Input()

			lowestFuel := converted.TotalFuelExp()
			day7part2logger.Debug().Msgf("lowest: %d", lowestFuel)
			highestPos := (*converted)[len(*converted)-1]
			for i := 1; i < highestPos; i++ {
				converted.ShiftLeft()
				fuel := converted.TotalFuelExp()
				day7part2logger.Debug().Msgf("new fuel: %d | lowest: %d", fuel, lowestFuel)
				if fuel > lowestFuel {
					break
				}
				lowestFuel = fuel
			}

			fmt.Printf("lowest: %d\n", lowestFuel)
		},
	}
)

func prepareday7Input() *CrabPositions {
	content, err := os.ReadFile("resources/day7.txt")
	if err != nil {
		day7logger.Fatal().Err(err).Send()
	}

	input := strings.Split(strings.TrimSpace(string(content)), ",")
	day7logger.Info().Msgf("length of input file: %d", len(input))
	day7logger.Debug().Msgf("plain input: %v", input)

	converted := make(CrabPositions, 0, len(input))
	for i := 0; i < len(input); i++ {
		pos := cast.ToInt(input[i])
		converted = append(converted, pos)
	}
	sort.Slice(converted, func(i, j int) bool { return converted[i] < converted[j] })
	day7logger.Debug().Msgf("converted input: %v", converted)

	return &converted
}

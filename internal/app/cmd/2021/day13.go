package cmd2021

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

type FoldableMap utils.SingleSliceMap[uint64]

func (m *FoldableMap) FoldAlong(axis rune, pos uint) *FoldableMap {
	switch axis {
	case 'x':
		return m.FoldAlongX(pos)
	case 'y':
		return m.FoldAlongY(pos)
	default:
		day13logger.Error().Msgf("tried to fold along unknown axis")
		return nil
	}
}

func (m *FoldableMap) FoldAlongX(pos uint) *FoldableMap {
	foldedMap := utils.SingleSliceMap[uint64]{}
	for _, elem := range (*utils.SingleSliceMap[uint64])(m).Elems {
		newX := elem.X
		if elem.X > pos {
			newX = pos - (elem.X - pos)
		}
		foldedMap.ModifyElem(func(elem uint64) uint64 {
			return 0
		}, newX, elem.Y)
	}
	return (*FoldableMap)(&foldedMap)
}

func (m *FoldableMap) FoldAlongY(pos uint) *FoldableMap {
	foldedMap := utils.SingleSliceMap[uint64]{}
	for _, elem := range (*utils.SingleSliceMap[uint64])(m).Elems {
		newY := elem.Y
		if elem.Y > pos {
			newY = pos - (elem.Y - pos)
		}
		foldedMap.ModifyElem(func(elem uint64) uint64 {
			return 0
		}, elem.X, newY)
	}
	return (*FoldableMap)(&foldedMap)
}

func (m *FoldableMap) GetLen() uint {
	return uint(len(((*utils.SingleSliceMap[uint64])(m)).Elems))
}

func init() {
	cmd2021.AddCommand(day13Cmd)

	day13Cmd.AddCommand(day13part1Cmd)
	day13Cmd.AddCommand(day13part2Cmd)
}

var (
	day13logger = log.With().
			Int("day", 13).
			Logger()
	day13Cmd = &cobra.Command{
		Use:   "day13",
		Short: "Day 13 Challenge",
	}
	day13part1Cmd = &cobra.Command{
		Use:   "part1",
		Short: "Day 13 Part 1 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day13part1logger := day13logger.With().
				Int("part", 1).
				Logger()
			day13part1logger.Info().Msg("Start")
			foldableMap, foldInstructions := prepareday13Input()

			day13logger.Debug().Msgf("fold instruction to execute next: %s", foldInstructions[0])

			instruction := strings.Split(foldInstructions[0], "=")
			foldPos, _ := strconv.ParseUint(instruction[1], 10, 32)

			day13logger.Debug().Msgf("dots before folding: %d", foldableMap.GetLen())

			foldableMap = foldableMap.FoldAlong(rune(instruction[0][0]), uint(foldPos))

			fmt.Printf("%d\n", foldableMap.GetLen())
		},
	}
	day13part2Cmd = &cobra.Command{
		Use:   "part2",
		Short: "Day 13 Part 2 Challenge",
		Run: func(cmd *cobra.Command, args []string) {
			day13part2logger := day13logger.With().
				Int("part", 2).
				Logger()
			day13part2logger.Info().Msg("Start")
			foldableMap, foldInstructions := prepareday13Input()

			for _, nextInstruction := range foldInstructions {
				day13logger.Debug().Msgf("fold instruction to execute next: %s", nextInstruction)

				instruction := strings.Split(nextInstruction, "=")
				foldPos, _ := strconv.ParseUint(instruction[1], 10, 32)

				day13logger.Debug().Msgf("dots before folding: %d", foldableMap.GetLen())

				foldableMap = foldableMap.FoldAlong(rune(instruction[0][0]), uint(foldPos))

				day13logger.Debug().Msgf("dots after folding: %d", foldableMap.GetLen())
			}

			fmt.Printf("%s\n", (*utils.SingleSliceMap[uint64])(foldableMap).String())
		},
	}
)

func prepareday13Input() (*FoldableMap, []string) {
	content, err := os.ReadFile("resources/day13.txt")
	if err != nil {
		day13logger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
	day13logger.Info().Msgf("length of input file: %d", len(input))
	day13logger.Debug().Msgf("plain input: %v", input)

	indexOfEmpty := indexOfEmpty(input)
	foldInstructions := input[indexOfEmpty+1:]
	day13logger.Debug().Msgf("fold instructions: %v", foldInstructions)

	for i := 0; i < len(foldInstructions); i++ {
		foldInstructions[i] = strings.TrimLeft(foldInstructions[i], "fold along ")
	}

	input = input[:indexOfEmpty]
	day13logger.Debug().Msgf("map dots: %v", input)
	converted := utils.SingleSliceMap[uint64]{}

	for _, dot := range input {
		coordinates := strings.Split(dot, ",")

		converted.ModifyElem(func(elem uint64) uint64 {
			return 0
		}, cast.ToUint(coordinates[0]), cast.ToUint(coordinates[1]))
	}
	day13logger.Debug().Msgf("map dots parsed: %s", converted)

	return (*FoldableMap)(&converted), foldInstructions
}

func indexOfEmpty(in []string) int {
	for index, elem := range in {
		if elem == "" {
			return index
		}
	}
	return -1
}

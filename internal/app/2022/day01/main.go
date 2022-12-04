package day01

import (
	"fmt"
	"sort"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "01"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		part1Func(preparedInput)
	case 2:
		part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(carryingCaloriesPerElf []int) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	fmt.Printf("most carried calories: %d\n", carryingCaloriesPerElf[0])
	puzzleAnswer = cast.ToUint64(carryingCaloriesPerElf[0])
	return puzzleAnswer
}

func part2Func(carryingCaloriesPerElf []int) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	var totalCalories int
	for _, heavyLoadElf := range carryingCaloriesPerElf[:3] {
		totalCalories += heavyLoadElf
	}
	fmt.Printf("carried calories by top three elves: %d\n", totalCalories)
	puzzleAnswer = cast.ToUint64(totalCalories)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []int {
	input := strings.Split(rawInput, "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	converted := []int{0}
	elf := 0
	for _, line := range input {
		if line == "" {
			partLogger.Debug().Msgf("new elf, last elf carries %d calories", converted[elf])
			elf++
			converted = append(converted, 0)
			continue
		}
		calories := cast.ToInt(line)
		converted[elf] += calories
	}
	sort.Sort(sort.Reverse(sort.IntSlice(converted)))

	partLogger.Debug().Msgf("converted input: %v", converted)
	return converted
}

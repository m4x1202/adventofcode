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

func Part1(args []string) {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	carryingCaloriesPerElf := prepareInput()

	fmt.Printf("most carried calories: %d\n", carryingCaloriesPerElf[0])
}

func Part2(args []string) {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	carryingCaloriesPerElf := prepareInput()

	var totalCalories int
	for _, heavyLoadElf := range carryingCaloriesPerElf[:3] {
		totalCalories += heavyLoadElf
	}
	fmt.Printf("carried calories by top three elves: %d\n", totalCalories)
}

func prepareInput() []int {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}

	input := strings.Split(string(content), "\n")
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

package day01

import (
	"fmt"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
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
	depths := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		part1Func(depths)
	case 2:
		part2Func(depths)
	default:
		panic("part does not exist")
	}
}

func part1Func(depths []uint16) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for i := 1; i < len(depths); i++ {
		if depths[i-1] < depths[i] {
			puzzleAnswer++
		}
	}

	fmt.Printf("num of increases: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(depths []uint16) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	sliding := utils.SlidingWindow[[][]uint16](depths, 3)
	slidingSums := utils.CombineFunc(sliding, func(in []uint16) uint16 {
		var res uint16
		for _, v := range in {
			res += v
		}
		return res
	})

	for i := 1; i < len(slidingSums); i++ {
		if slidingSums[i-1] < slidingSums[i] {
			puzzleAnswer++
		}
	}

	fmt.Printf("num of sliding increases: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2021/day%s/input.txt", DAY))
	if err != nil {
		partLogger.Fatal().Err(err).Send()
	}
	return strings.TrimSpace(string(content))
}

func prepareInput(rawInput string) []uint16 {
	input := strings.Split(rawInput, "\n")
	partLogger.Info().Msgf("length of input file: %d", len(input))
	partLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]uint16, len(input))
	for i := 0; i < len(input); i++ {
		converted[i] = cast.ToUint16(input[i])
		partLogger.Trace().Int("index", i).Msgf("converted: %d", converted[i])
	}
	partLogger.Debug().Msgf("converted input: %v", converted)

	return converted
}

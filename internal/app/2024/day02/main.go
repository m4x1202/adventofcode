package day02

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "02"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger zerolog.Logger
)

func ExecutePart(p uint8) {
	preparedInput := prepareInput(readPuzzleInput())
	var result uint64
	switch p {
	case 1:
		result = part1Func(preparedInput)
	case 2:
		result = part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
	fmt.Printf("Result: %d\n", result)
}

func part1Func(preparedInput [][]uint) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	for _, report := range preparedInput {
		partLogger.Debug().Msgf("%v", report)
		sorted := slices.IsSorted(report)
		partLogger.Debug().Msgf("Sorted %v", sorted)
		if !sorted {
			reversed := make([]uint, len(report))
			copy(reversed, report)
			slices.Reverse(reversed)
			partLogger.Debug().Msgf("%v", reversed)
			sorted = slices.IsSorted(reversed)
			partLogger.Debug().Msgf("Reverse Sorted %v", sorted)
		}
		if !sorted {
			continue
		}
		safe := true
		previousLevel := report[0]
		for i := 1; i < len(report); i++ {
			difference := math.Abs(float64(previousLevel) - float64(report[i]))
			if difference > 3 || difference < 1 {
				safe = false
				break
			}
			previousLevel = report[i]
		}
		if safe {
			puzzleAnswer++
		}
	}

	return puzzleAnswer
}

func part2Func(preparedInput [][]uint) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	// Logic here
	puzzleAnswer = cast.ToUint64(0)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2024/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) [][]uint {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := make([][]uint, len(input))
	for i := range input {
		splitted := strings.Split(input[i], " ")
		converted[i] = make([]uint, len(splitted))
		for j := range splitted {
			converted[i][j] = cast.ToUint(splitted[j])
		}
	}

	return converted
}

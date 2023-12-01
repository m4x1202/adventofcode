package day01

import (
	"fmt"
	"slices"
	"strconv"
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

func part1Func(preparedInput []string) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")

	values := make([]uint64, len(preparedInput))
	for i := range preparedInput {
		calibrationLine := []rune(preparedInput[i])
		firstNumIndex := strings.IndexAny(string(calibrationLine), "0123456789")
		firstNum := calibrationLine[firstNumIndex]

		slices.Reverse(calibrationLine)
		lastNumIndex := strings.IndexAny(string(calibrationLine), "0123456789")
		lastNum := calibrationLine[lastNumIndex]

		finalNumStr := string(firstNum) + string(lastNum)
		if finalNum, err := strconv.ParseUint(finalNumStr, 10, 64); err == nil {
			values[i] = finalNum
		}
	}

	var finalSum uint64
	for _, num := range values {
		finalSum += num
	}

	return finalSum
}

func part2Func(preparedInput []string) uint64 {
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
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2023/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []string {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := make([]string, len(input))
	for i := range input {
		converted[i] = input[i]
	}

	return converted
}

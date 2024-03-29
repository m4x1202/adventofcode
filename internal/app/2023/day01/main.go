package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DAY = "01"
)

var (
	dayLogger = log.With().
			Str("day", DAY).
			Logger()
	partLogger    zerolog.Logger
	numberStrings = map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
)

func ExecutePart(p uint8) uint64 {
	preparedInput := prepareInput(readPuzzleInput())
	switch p {
	case 1:
		return part1Func(preparedInput)
	case 2:
		return part2Func(preparedInput)
	default:
		panic("part does not exist")
	}
}

func part1Func(preparedInput []string) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")

	values := make([]uint64, len(preparedInput))
	for i := range preparedInput {
		calibrationLine := []rune(preparedInput[i])
		firstNumIndex := strings.IndexAny(string(calibrationLine), "123456789")
		firstNum := calibrationLine[firstNumIndex]

		slices.Reverse(calibrationLine)
		lastNumIndex := strings.IndexAny(string(calibrationLine), "123456789")
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

	values := make([]uint64, len(preparedInput))
	for i := range preparedInput {
		replacedInput := ReplaceAll(preparedInput[i])
		partLogger.Debug().Msg(replacedInput)
		calibrationLine := []rune(replacedInput)

		firstNumIndex := strings.IndexAny(replacedInput, "123456789")
		firstNum := replacedInput[firstNumIndex]

		slices.Reverse(calibrationLine)
		lastNumIndex := strings.IndexAny(string(calibrationLine), "123456789")
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

func ReplaceAll(s string) string {
	result := s
	for k, v := range numberStrings {
		result = strings.ReplaceAll(result, k, v)
	}
	return result
}

package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

const (
	DAY = "03"
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
	var puzzleAnswer uint64

	for _, i := range preparedInput {
		puzzleAnswer += cast.ToUint64(runMul(i))
	}

	return puzzleAnswer
}

func part2Func(preparedInput []string) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	enabled := true
	for _, i := range preparedInput {
		if !enabled && i != "do()" {
			continue
		}
		switch i {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			puzzleAnswer += cast.ToUint64(runMul(i))
		}
	}

	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2024/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) []string {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)|do\\(\\)|don't\\(\\)")
	converted := r.FindAllString(strings.Join(input, ""), -1)

	return converted
}

func runMul(in string) int {
	trimmedIn := strings.TrimPrefix(in, "mul(")
	trimmedIn = strings.TrimSuffix(trimmedIn, ")")
	strA, strB, _ := strings.Cut(trimmedIn, ",")
	numA, _ := strconv.Atoi(strA)
	numB, _ := strconv.Atoi(strB)

	return numA * numB
}

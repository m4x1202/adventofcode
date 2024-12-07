package day01

import (
	"fmt"
	"math"
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
	var result uint
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

func part1Func(preparedInput [2][]uint) uint {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")

	sort.Slice(preparedInput[0], func(i, j int) bool { return preparedInput[0][i] < preparedInput[0][j] })
	sort.Slice(preparedInput[1], func(i, j int) bool { return preparedInput[1][i] < preparedInput[1][j] })

	var totalDistance uint
	for i := 0; i < len(preparedInput[0]); i++ {
		totalDistance += cast.ToUint(math.Abs(float64(preparedInput[0][i]) - float64(preparedInput[1][i])))
	}

	return totalDistance
}

func part2Func(preparedInput [2][]uint) uint {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")

	foundValues := map[uint]uint{}
	for _, id := range preparedInput[0] {
		if _, ok := foundValues[id]; ok {
			continue
		}
		var count uint
		for _, elem := range preparedInput[1] {
			if id == elem {
				count++
			}
		}
		foundValues[id] = count
	}

	var totalSimilarity uint
	for _, id := range preparedInput[0] {
		totalSimilarity += id * foundValues[id]
	}

	return totalSimilarity
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2024/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return string(content)
}

func prepareInput(rawInput string) [2][]uint {
	input := strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n")
	dayLogger.Info().Msgf("length of input file: %d", len(input))
	dayLogger.Debug().Msgf("plain input: %v", input)

	converted := [2][]uint{}
	converted[0] = make([]uint, len(input))
	converted[1] = make([]uint, len(input))
	for i := range input {
		splitted := strings.Split(input[i], "   ")
		converted[0][i] = cast.ToUint(splitted[0])
		converted[1][i] = cast.ToUint(splitted[1])
	}

	return converted
}

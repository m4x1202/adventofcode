package day06

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/m4x1202/adventofcode/pkg/utils"
	"github.com/m4x1202/adventofcode/resources"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DAY = "06"
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

func part1Func(preparedInput []rune) uint64 {
	partLogger = dayLogger.With().
		Int("part", 1).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	possibleMarkers := utils.SlidingWindow[[][]rune](preparedInput, 4)
	for i, possibleMarker := range possibleMarkers {
		if reflect.DeepEqual(utils.RemoveDups(possibleMarker), possibleMarker) {
			puzzleAnswer = uint64(i + 4)
			break
		}
	}

	fmt.Printf("characters processed to find start-of-packet marker: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func part2Func(preparedInput []rune) uint64 {
	partLogger = dayLogger.With().
		Int("part", 2).
		Logger()
	partLogger.Info().Msg("Start")
	var puzzleAnswer uint64

	possibleMarkers := utils.SlidingWindow[[][]rune](preparedInput, 14)
	for i, possibleMarker := range possibleMarkers {
		if reflect.DeepEqual(utils.RemoveDups(possibleMarker), possibleMarker) {
			puzzleAnswer = uint64(i + 14)
			break
		}
	}

	fmt.Printf("characters processed to find start-of-message marker: %d\n", puzzleAnswer)
	return puzzleAnswer
}

func readPuzzleInput() string {
	content, err := resources.InputFS.ReadFile(fmt.Sprintf("2022/day%s/input.txt", DAY))
	if err != nil {
		dayLogger.Fatal().Err(err).Send()
	}
	return strings.TrimSpace(string(content))
}

func prepareInput(rawInput string) []rune {
	dayLogger.Debug().Msgf("plain input: %s", rawInput)

	return []rune(rawInput)
}
